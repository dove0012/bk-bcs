/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package tasks

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	proto "github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/internal/cloudprovider"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/internal/cloudprovider/huawei/api"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/internal/common"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cce/v3/model"
)

// ApplyInstanceMachinesTask update desired nodes task
func ApplyInstanceMachinesTask(taskID string, stepName string) error {
	start := time.Now()

	// get task and task current step
	state, step, err := cloudprovider.GetTaskStateAndCurrentStep(taskID, stepName)
	if err != nil {
		return err
	}
	// previous step successful when retry task
	if step == nil {
		return nil
	}

	// extract parameter && check validate
	clusterID := step.Params[cloudprovider.ClusterIDKey.String()]
	nodeGroupID := step.Params[cloudprovider.NodeGroupIDKey.String()]
	cloudID := step.Params[cloudprovider.CloudIDKey.String()]
	desiredNodes := step.Params[cloudprovider.ScalingKey.String()]
	nodeNum, _ := strconv.Atoi(desiredNodes)
	operator := step.Params[cloudprovider.OperatorKey.String()]
	if len(clusterID) == 0 || len(nodeGroupID) == 0 || len(cloudID) == 0 || len(desiredNodes) == 0 || len(operator) == 0 {
		blog.Errorf("ApplyInstanceMachinesTask[%s]: check parameter validate failed", taskID)
		retErr := fmt.Errorf("ApplyInstanceMachinesTask check parameters failed")
		_ = cloudprovider.UpdateNodeGroupDesiredSize(nodeGroupID, nodeNum, true)
		_ = state.UpdateStepFailure(start, stepName, retErr)
		return retErr
	}
	dependInfo, err := cloudprovider.GetClusterDependBasicInfo(clusterID, cloudID, nodeGroupID)
	if err != nil {
		blog.Errorf("ApplyInstanceMachinesTask[%s]: GetClusterDependBasicInfo failed: %s", taskID, err.Error())
		retErr := fmt.Errorf("ApplyInstanceMachinesTask GetClusterDependBasicInfo failed")
		_ = cloudprovider.UpdateNodeGroupDesiredSize(nodeGroupID, nodeNum, true)
		_ = state.UpdateStepFailure(start, stepName, retErr)
		return retErr
	}

	// inject taskID
	ctx := cloudprovider.WithTaskIDForContext(context.Background(), taskID)
	err = applyInstanceMachines(ctx, dependInfo, int32(nodeNum))
	if err != nil {
		blog.Errorf("ApplyInstanceMachinesTask[%s]: applyInstanceMachines failed: %s", taskID, err.Error())
		retErr := fmt.Errorf("ApplyInstanceMachinesTask applyInstanceMachines failed")
		_ = cloudprovider.UpdateNodeGroupDesiredSize(nodeGroupID, nodeNum, true)
		_ = state.UpdateStepFailure(start, stepName, retErr)
		return retErr
	}

	// trans success nodes to cm DB and record common paras, not handle error
	_ = recordClusterInstanceToDB(ctx, state, dependInfo, nodeNum)

	// update step
	if err = state.UpdateStepSucc(start, stepName); err != nil {
		blog.Errorf("ApplyInstanceMachinesTask[%s] task %s %s update to storage fatal", taskID, taskID, stepName)
		return err
	}
	return nil
}

// applyInstanceMachines apply machines from MIG
func applyInstanceMachines(ctx context.Context, info *cloudprovider.CloudDependBasicInfo, nodeNum int32) error {
	taskID := cloudprovider.GetTaskIDFromContext(ctx)

	client, err := api.NewCceClient(info.CmOption)
	if err != nil {
		return err
	}

	rsp, err := client.UpdateDesiredNodes(info.Cluster.SystemID, info.NodeGroup.CloudNodeGroupID, nodeNum)
	if err != nil {
		return err
	}

	err = cloudprovider.LoopDoFunc(context.Background(), func() error {
		if *rsp.Status.CurrentNode == nodeNum && rsp.Status.Phase.Value() == "" {
			return cloudprovider.EndLoop
		} else if rsp.Status.Phase.Value() == model.GetNodePoolStatusPhaseEnum().ERROR.Value() {
			return fmt.Errorf("applyInstanceMachines[%s] GetOperation failed: %v", taskID, err)
		}

		blog.Infof("taskID[%s] operation %s still running", taskID, rsp.Status.Phase.Value())
		return nil
	}, cloudprovider.LoopInterval(3*time.Second))

	if err != nil {
		return fmt.Errorf("applyInstanceMachines[%s] GetOperation failed: %v", taskID, err)
	}

	return nil
}

// recordClusterInstanceToDB already auto build instances to cluster, thus not handle error
func recordClusterInstanceToDB(ctx context.Context, state *cloudprovider.TaskState,
	info *cloudprovider.CloudDependBasicInfo, nodeNum int) error {
	var (
		successInstanceID []string
		failedInstanceID  []string
	)
	taskID := cloudprovider.GetTaskIDFromContext(ctx)

	client, err := api.NewCceClient(info.CmOption)
	if err != nil {
		return err
	}

	instaces, err := client.ListClusterNodePoolNodes(info.Cluster.SystemID, info.NodeGroup.CloudNodeGroupID)
	if err != nil {
		return err
	}

	for _, v := range instaces {
		if v.Status.Phase.Value() == model.GetNodeStatusPhaseEnum().ACTIVE.Value() {
			successInstanceID = append(successInstanceID, *v.Status.ServerId)
		} else {
			failedInstanceID = append(failedInstanceID, *v.Status.ServerId)
		}
	}

	// rollback desired num
	if len(successInstanceID) != nodeNum {
		_ = cloudprovider.UpdateNodeGroupDesiredSize(info.NodeGroup.NodeGroupID, nodeNum-len(successInstanceID), true)
	}

	// record instanceIDs to task common
	if state.Task.CommonParams == nil {
		state.Task.CommonParams = make(map[string]string)
	}
	// remove existed instanceID
	var newInstancesID []string
	for _, n := range successInstanceID {
		if existNode, _ := cloudprovider.GetStorageModel().GetNode(ctx, n); existNode != nil && existNode.InnerIP != "" {
			continue
		}
		newInstancesID = append(newInstancesID, n)
	}
	if len(newInstancesID) > 0 {
		state.Task.CommonParams[cloudprovider.SuccessNodeIDsKey.String()] = strings.Join(newInstancesID, ",")
		state.Task.CommonParams[cloudprovider.NodeIDsKey.String()] = strings.Join(newInstancesID, ",")
	}
	if len(failedInstanceID) > 0 {
		state.Task.CommonParams[cloudprovider.FailedNodeIDsKey.String()] = strings.Join(failedInstanceID, ",")
	}

	// record successNodes to cluster manager DB
	nodeIPs, err := transInstancesToNode(ctx, newInstancesID, info)
	if err != nil {
		blog.Errorf("recordClusterInstanceToDB[%s] failed: %v", taskID, err)
	}
	if len(nodeIPs) > 0 {
		state.Task.CommonParams[cloudprovider.NodeIPsKey.String()] = strings.Join(nodeIPs, ",")
	}

	return nil
}

// transInstancesToNode record success nodes to cm DB
func transInstancesToNode(ctx context.Context, instanceID []string, info *cloudprovider.CloudDependBasicInfo) (
	[]string, error) {
	var (
		nodeIPs = make([]string, 0)
		err     error
	)
	taskID := cloudprovider.GetTaskIDFromContext(ctx)

	client, err := api.NewEcsClient(info.CmOption)
	if err != nil {
		return nil, err
	}

	servers, err := client.ListEcsDetails(instanceID)
	if err != nil {
		return nil, err
	}

	for _, v := range servers {
		node := proto.Node{}
		node.NodeID = v.Id
		node.InstanceType = v.Flavor.Name
		node.VPC = v.Metadata["vpc_id"]
		if net, ok := v.Addresses[v.Metadata["vpc_id"]]; ok {
			if len(net) > 0 {
				node.InnerIP = net[0].Addr
			}
		}
		node.Region = info.CmOption.Region

		node.ClusterID = info.NodeGroup.ClusterID
		node.NodeGroupID = info.NodeGroup.NodeGroupID
		node.Status = common.StatusInitialization
		err = cloudprovider.SaveNodeInfoToDB(&node)
		if err != nil {
			blog.Errorf("transInstancesToNode[%s] SaveNodeInfoToDB[%s] failed: %v", taskID, node.InnerIP, err)
		}

		nodeIPs = append(nodeIPs, node.InnerIP)
	}

	return nodeIPs, nil
}

// CheckClusterNodesStatusTask check update desired nodes status task. nodes already add to cluster, thus not rollback desiredNum and only record status
func CheckClusterNodesStatusTask(taskID string, stepName string) error {
	start := time.Now()

	// get task and task current step
	state, step, err := cloudprovider.GetTaskStateAndCurrentStep(taskID, stepName)
	if err != nil {
		return err
	}
	// previous step successful when retry task
	if step == nil {
		return nil
	}

	// step login started here
	// extract parameter && check validate
	clusterID := step.Params[cloudprovider.ClusterIDKey.String()]
	nodeGroupID := step.Params[cloudprovider.NodeGroupIDKey.String()]
	cloudID := step.Params[cloudprovider.CloudIDKey.String()]
	successInstanceID := strings.Split(state.Task.CommonParams[cloudprovider.SuccessNodeIDsKey.String()], ",")

	if len(clusterID) == 0 || len(nodeGroupID) == 0 || len(cloudID) == 0 || len(successInstanceID) == 0 {
		blog.Errorf("CheckClusterNodesStatusTask[%s]: check parameter validate failed", taskID)
		retErr := fmt.Errorf("CheckClusterNodesStatusTask check parameters failed")
		_ = state.UpdateStepFailure(start, stepName, retErr)
		return retErr
	}
	dependInfo, err := cloudprovider.GetClusterDependBasicInfo(clusterID, cloudID, nodeGroupID)
	if err != nil {
		blog.Errorf("CheckClusterNodesStatusTask[%s]: GetClusterDependBasicInfo failed: %s", taskID, err.Error())
		retErr := fmt.Errorf("CheckClusterNodesStatusTask GetClusterDependBasicInfo failed")
		_ = state.UpdateStepFailure(start, stepName, retErr)
		return retErr
	}

	// inject taskID
	ctx := cloudprovider.WithTaskIDForContext(context.Background(), taskID)
	successInstances, failureInstances, err := checkClusterInstanceStatus(ctx, dependInfo, successInstanceID)
	if err != nil {
		blog.Errorf("CheckClusterNodesStatusTask[%s]: checkClusterInstanceStatus failed: %s", taskID, err.Error())
		retErr := fmt.Errorf("CheckClusterNodesStatusTask checkClusterInstanceStatus failed")
		_ = state.UpdateStepFailure(start, stepName, retErr)
		return retErr
	}

	// update response information to task common params
	if state.Task.CommonParams == nil {
		state.Task.CommonParams = make(map[string]string)
	}
	if len(successInstances) > 0 {
		state.Task.CommonParams[cloudprovider.SuccessClusterNodeIDsKey.String()] = strings.Join(successInstances, ",")
		// dynamic inject paras
		state.Task.CommonParams[cloudprovider.DynamicNodeIPListKey.String()] = strings.Join(successInstances, ",")
	}
	if len(failureInstances) > 0 {
		state.Task.CommonParams[cloudprovider.FailedClusterNodeIDsKey.String()] = strings.Join(failureInstances, ",")
	}

	// update step
	if err = state.UpdateStepSucc(start, stepName); err != nil {
		blog.Errorf("CheckClusterNodesStatusTask[%s] task %s %s update to storage fatal", taskID, taskID, stepName)
		return err
	}

	return nil
}

// checkClusterInstanceStatus 检测节点加入集群的状态
func checkClusterInstanceStatus(ctx context.Context, info *cloudprovider.CloudDependBasicInfo,
	instanceIDs []string) ([]string, []string, error) {
	//由于华为云的弹性伸缩是通过华为云自身的节点池来实现,所以扩容的节点必然是在集群里的,不需要加测此状态
	return instanceIDs, []string{}, nil
}
