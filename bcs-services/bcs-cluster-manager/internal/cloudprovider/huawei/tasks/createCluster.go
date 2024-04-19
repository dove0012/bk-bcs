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
 */

package tasks

import (
	"context"
	"fmt"
	"time"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/internal/cloudprovider"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/internal/cloudprovider/huawei/api"
)

// CreateClusterTask call huawei interface to create cluster
func CreateClusterTask(taskID string, stepName string) error {
	start := time.Now()

	// get task and task current step
	state, step, err := cloudprovider.GetTaskStateAndCurrentStep(taskID, stepName)
	if err != nil {
		return err
	}
	// previous step successful when retry task
	if step == nil {
		blog.Infof("CreateClusterTask[%s]: current step[%s] successful and skip", taskID, stepName)
		return nil
	}
	blog.Infof("CreateClusterTask[%s]: task %s run step %s, system: %s, old state: %s, params %v",
		taskID, taskID, stepName, step.System, step.Status, step.Params)

	clusterID := step.Params[cloudprovider.ClusterIDKey.String()]
	cloudID := step.Params[cloudprovider.CloudIDKey.String()]
	operator := state.Task.CommonParams[cloudprovider.OperatorKey.String()]

	// get dependent basic info
	dependInfo, err := cloudprovider.GetClusterDependBasicInfo(cloudprovider.GetBasicInfoReq{
		ClusterID: clusterID,
		CloudID:   cloudID,
	})
	if err != nil {
		blog.Errorf("CreateClusterTask[%s]: GetClusterDependBasicInfo for cluster %s in task %s "+
			"step %s failed, %s", taskID, clusterID, taskID, stepName, err.Error()) // nolint
		retErr := fmt.Errorf("get cloud/project information failed, %s", err.Error())
		_ = state.UpdateStepFailure(start, stepName, retErr)
		return retErr
	}

	// inject taskID
	ctx := cloudprovider.WithTaskIDForContext(context.Background(), taskID)

	clsId, err := createCCECluster(ctx, dependInfo, operator)
	if err != nil {
		blog.Errorf("CreateClusterTask[%s] createCluster for cluster[%s] failed, %s",
			taskID, clusterID, err.Error())
		retErr := fmt.Errorf("createCluster err, %s", err.Error())
		_ = state.UpdateStepFailure(start, stepName, retErr)

		_ = cloudprovider.UpdateClusterErrMessage(clusterID, fmt.Sprintf("submit createCluster[%s] failed: %v",
			dependInfo.Cluster.GetClusterID(), err))
		return retErr
	}

	// update response information to task common params
	if state.Task.CommonParams == nil {
		state.Task.CommonParams = make(map[string]string)
	}
	state.Task.CommonParams[cloudprovider.CloudSystemID.String()] = clsId

	// update step
	if err = state.UpdateStepSucc(start, stepName); err != nil {
		blog.Errorf("CreateClusterTask[%s] task %s %s update to storage fatal", taskID, taskID, stepName)
		return err
	}

	return nil
}

// createCCECluster create huawei cce cluster
func createCCECluster(ctx context.Context, info *cloudprovider.CloudDependBasicInfo, operator string) (string, error) {
	taskID := cloudprovider.GetTaskIDFromContext(ctx)

	req, err := generateCreateClusterRequest(ctx, info, operator)
	if err != nil {
		blog.Errorf("CreateClusterTask[%s] generateCreateClusterRequest failed: %v", taskID, err)
		return "", err
	}

	systemId, err := createCluster(ctx, info, req, info.Cluster.SystemID)
	if err != nil {
		blog.Errorf("CreateClusterTask[%s] call createCluster[%s] failed, %s",
			taskID, info.Cluster.ClusterID, err.Error())
		retErr := fmt.Errorf("call CreateClusterTask[%s] api err, %s", info.Cluster.ClusterID, err.Error())
		return "", retErr
	}

	blog.Infof("CreateClusterTask[%s] ClusterID[%s] successful", taskID, info.Cluster.ClusterID)

	// update cluster systemID
	info.Cluster.SystemID = systemId

	err = cloudprovider.GetStorageModel().UpdateCluster(ctx, info.Cluster)
	if err != nil {
		blog.Errorf("CreateClusterTask[%s] updateClusterSystemID[%s] failed %s",
			taskID, info.Cluster.ClusterID, err.Error())
		retErr := fmt.Errorf("call CreateClusterTask updateClusterSystemID[%s] api err: %s",
			info.Cluster.ClusterID, err.Error())
		return "", retErr
	}
	blog.Infof("CreateClusterTask[%s] call CreateClusterTask updateClusterSystemID successful", taskID)

	return systemId, nil
}

func generateCreateClusterRequest(ctx context.Context, info *cloudprovider.CloudDependBasicInfo,
	operator string) (*api.CreateClusterRequest, error) {

	ClusterTags := make([]api.ResourceTag, 0)
	for k, v := range info.Cluster.Labels {
		ClusterTags = append(ClusterTags, api.ResourceTag{Key: k, Value: v})
	}
	req := &api.CreateClusterRequest{
		Body: &api.Cluster{
			Kind:       "cluster",
			ApiVersion: "v3",
			Metadata: &api.ClusterMetadata{
				Name: info.Cluster.ClusterName,
			},
			Spec: &api.ClusterSpec{
				Category:    "CCE",
				Type:        "VirtualMachine",
				Flavor:      info.Cluster.ClusterBasicSettings.ClusterLevel, // 需要转换成cce.s1.medium这样的形式
				Version:     &info.Cluster.ClusterBasicSettings.Version,
				Description: &info.Cluster.Description,
				HostNetwork: &api.HostNetwork{
					Vpc:           info.Cluster.VpcID,
					Subnet:        "", // 可能需要遍历instances字段里判断nodeRole为MASTER_ETCD的节点中获取
					SecurityGroup: "",
				},
				ServiceNetwork: &api.ServiceNetwork{
					IPv4CIDR: info.Cluster.NetworkSettings.ServiceIPv4CIDR, // 需要判断前端传的网络插件类型，如果是Global Router则需要根据ip数量计算
				},
				BillingMode: 0,
				ClusterTags: ClusterTags,
				//KubeProxyMode: "iptables",
			},
		},
	}
	return req, nil
}

func createCluster(ctx context.Context, info *cloudprovider.CloudDependBasicInfo,
	req *api.CreateClusterRequest, systemID string) (string, error) {
	return "", nil
}
