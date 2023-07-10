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

package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	proto "github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
	"github.com/amoghe/go-crypt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cce/v3/model"
	"strconv"
)

// GetClusterKubeConfig get cce cluster kebeconfig
func GetClusterKubeConfig(client *CceClient, clusterId string) (string, error) {
	req := model.CreateKubernetesClusterCertRequest{
		ClusterId: clusterId, // 集群ID，可在CCE管理控制台中查看
		Body: &model.CertDuration{
			Duration: int32(-1), // 集群证书有效时间，单位为天，最小值为1，最大值为10950(30*365，1年固定计365天，忽略闰年影响)；若填-1则为最大值30年。
		},
	}

	rsp, err := client.CreateKubernetesClusterCert(&req)
	if err != nil {
		return "", err
	}

	bt, err := json.Marshal(rsp)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(bt), nil
}

// GenerateCreateNodePoolRequest get cce nodepool request
func GenerateCreateNodePoolRequest(group *proto.NodeGroup, cluster *proto.Cluster) (*model.CreateNodePoolRequest, error) {
	var (
		initialNodeCount int32 = 0
		minNodeCount     int32 = 0
		maxNodeCount     int32 = 110
		enable                 = false
	)

	if group.NodeTemplate.MaxPodsPerNode != 0 {
		maxNodeCount = int32(group.NodeTemplate.MaxPodsPerNode)
	}

	nodeTemplate, err := GenerateNodeSpec(group)
	if err != nil {
		return nil, err
	}

	return &model.CreateNodePoolRequest{
		ClusterId: cluster.SystemID,
		Body: &model.NodePool{
			Kind:       "NodePool",
			ApiVersion: "v3",
			Metadata: &model.NodePoolMetadata{
				Name: group.NodeGroupID,
			},
			Spec: &model.NodePoolSpec{
				InitialNodeCount: &initialNodeCount,
				Autoscaling: &model.NodePoolNodeAutoscaling{
					Enable:       &enable,
					MinNodeCount: &minNodeCount,
					MaxNodeCount: &maxNodeCount,
				},
				NodeTemplate: nodeTemplate,
			},
		},
	}, nil
}

func GenerateNodeSpec(nodeGroup *proto.NodeGroup) (*model.NodeSpec, error) {
	if nodeGroup.LaunchTemplate == nil {
		return nil, fmt.Errorf("node group launch template is nil")
	}

	diskSize, err := strconv.Atoi(nodeGroup.LaunchTemplate.SystemDisk.DiskSize)
	if err != nil {
		return nil, err
	}

	dataVolumes := make([]model.Volume, 0)
	for _, v := range nodeGroup.LaunchTemplate.DataDisks {
		var size int
		size, err = strconv.Atoi(v.DiskSize)
		if err != nil {
			return nil, err
		}

		dataVolumes = append(dataVolumes, model.Volume{
			Volumetype: v.DiskType,
			Size:       int32(size),
		})
	}

	var (
		nodeBillingMode int32 = 0
	)
	password, err := Crypt(nodeGroup.LaunchTemplate.InitLoginPassword)
	if err != nil {
		return nil, err
	}

	return &model.NodeSpec{
		Flavor: nodeGroup.LaunchTemplate.InstanceType,
		Az:     nodeGroup.Region,
		Os:     &nodeGroup.NodeOS,
		Login: &model.Login{
			UserPassword: &model.UserPassword{
				//username不填默认为root，password必须加盐并base64加密
				Password: password,
			},
		},
		RootVolume: &model.Volume{
			Volumetype: nodeGroup.LaunchTemplate.SystemDisk.DiskType,
			Size:       int32(diskSize),
		},
		DataVolumes: dataVolumes,
		BillingMode: &nodeBillingMode,
	}, nil
}

// Crypt encryption node password
func Crypt(password string) (string, error) {
	salt := "tM3|cY3+tI4)"
	str, err := crypt.Crypt(password, salt)
	if err != nil {
		return "", err
	}

	return base64.RawStdEncoding.EncodeToString([]byte(str)), nil
}

// GenerateModifyClusterNodePoolInput get cce update nodepool input
func GenerateModifyClusterNodePoolInput(group *proto.NodeGroup,
	clusterID string) *model.UpdateNodePoolRequest {
	enable := false
	req := &model.UpdateNodePoolRequest{
		NodepoolId: group.CloudNodeGroupID,
		ClusterId:  clusterID,
		Body: &model.NodePoolUpdate{
			Metadata: &model.NodePoolMetadataUpdate{
				Name: group.NodeGroupID,
			},
			Spec: &model.NodePoolSpecUpdate{
				NodeTemplate:     &model.NodeSpecUpdate{},
				InitialNodeCount: int32(group.AutoScaling.DesiredSize),
				Autoscaling: &model.NodePoolNodeAutoscaling{
					Enable: &enable,
				},
			},
		},
	}

	if group.NodeTemplate != nil {
		for _, v := range group.NodeTemplate.Taints {
			effect := model.GetTaintEffectEnum().NO_SCHEDULE
			if v.Effect == "PreferNoSchedule" {
				effect = model.GetTaintEffectEnum().PREFER_NO_SCHEDULE
			} else if v.Effect == "NoExecute" {
				effect = model.GetTaintEffectEnum().NO_EXECUTE
			}
			value := v.Value
			req.Body.Spec.NodeTemplate.Taints = append(req.Body.Spec.NodeTemplate.Taints, model.Taint{
				Key:    v.Key,
				Value:  &value,
				Effect: effect,
			})
		}

		req.Body.Spec.NodeTemplate.K8sTags = group.Tags

		for k, v := range group.Tags {
			key := k
			value := v
			req.Body.Spec.NodeTemplate.UserTags = append(req.Body.Spec.NodeTemplate.UserTags, model.UserTag{
				Key:   &key,
				Value: &value,
			})
		}
	}

	return req
}
