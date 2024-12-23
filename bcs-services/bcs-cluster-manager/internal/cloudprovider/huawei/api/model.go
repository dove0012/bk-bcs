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

// Package api xxx
package api

import "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cce/v3/model"

// CreateKubernetesClusterCertResponse Response Object
type CreateKubernetesClusterCertResponse struct {

	// API类型，固定值“Config”，该值不可修改。
	Kind *string `json:"kind,omitempty"`

	// API版本，固定值“v1”。
	ApiVersion *string `json:"apiVersion,omitempty"`

	// 当前未使用该字段，当前默认为空。
	Preferences *interface{} `json:"preferences,omitempty"`

	// 集群列表。
	Clusters *[]model.Clusters `json:"clusters,omitempty"`

	// 存放了指定用户的一些证书信息和ClientKey信息。
	Users *[]model.Users `json:"users,omitempty"`

	// 上下文列表。
	Contexts *[]model.Contexts `json:"contexts,omitempty"`

	// 当前上下文，若存在publicIp（虚拟机弹性IP）时为 external; 若不存在publicIp为 internal。
	CurrentContext *string `json:"current-context,omitempty"`

	PortID         *string `json:"Port-ID,omitempty"`
	HttpStatusCode int     `json:"-"`
}

// Cluster 集群信息
type Cluster struct {
	Metadata *ClusterMetadata
	// Spec 集群参数定义
	Spec *ClusterSpec
	// Status 集群状态
	Status *ClusterStatus
}

// ClusterMetadata
type ClusterMetadata struct {
	// 集群名称。  命名规则：以小写字母开头，由小写字母、数字、中划线(-)组成，长度范围4-128位，且不能以中划线(-)结尾。
	Name string
	// 集群ID，资源唯一标识，创建成功后自动生成，填写无效。在创建包周期集群时，响应体不返回集群ID。
	Uid *string
}

// ClusterSpec 集群参数定义。
type ClusterSpec struct {
	// Type 集群Master节点架构：  - VirtualMachine：Master节点为x86架构服务器 [- ARM64: Master节点为鲲鹏（ARM架构）服务器](tag:hws,hws_hk,hcs)
	Type string
	// Version 集群版本，与Kubernetes社区基线版本保持一致，建议选择最新版本。
	Version *string
	// Description 集群描述
	Description *string
	// HostNetwork 集群网络配置
	HostNetwork *HostNetwork
	// ContainerNetwork 集群容器网络配置
	ContainerNetwork *ContainerNetwork
	// ServiceNetwork 集群服务网络配置
	ServiceNetwork *ServiceNetwork
	// KubeProxyMode kube-proxy代理模式
	KubeProxyMode string
	// 是否为Autopilot集群。
	EnableAutopilot *bool
	// ExtendParam 集群扩展参数
	ExtendParam *ClusterExtendParam
}

type ClusterStatus struct {
	// 集群状态，取值如下 - Available：可用，表示集群处于正常状态。 - Unavailable：不可用，表示集群异常，需手动删除。 - ScalingUp：扩容中，表示集群正处于扩容过程中。 - ScalingDown：缩容中，表示集群正处于缩容过程中。 - Creating：创建中，表示集群正处于创建过程中。 - Deleting：删除中，表示集群正处于删除过程中。 - Upgrading：升级中，表示集群正处于升级过程中。 - Resizing：规格变更中，表示集群正处于变更规格中。 - RollingBack：回滚中，表示集群正处于回滚过程中。 - RollbackFailed：回滚异常，表示集群回滚异常。 - Hibernating：休眠中，表示集群正处于休眠过程中。 - Hibernation：已休眠，表示集群正处于休眠状态。 - Awaking：唤醒中，表示集群正处于从休眠状态唤醒的过程中。 - Empty：集群无任何资源（已废弃）
	Phase *string
	// 任务ID,集群当前状态关联的任务ID。当前支持: - 创建集群时返回关联的任务ID，可通过任务ID查询创建集群的附属任务信息； - 删除集群或者删除集群失败时返回关联的任务ID，此字段非空时，可通过任务ID查询删除集群的附属任务信息。 > 任务信息具有一定时效性，仅用于短期跟踪任务进度，请勿用于集群状态判断等额外场景。
	JobID *string
	// 集群变为当前状态的原因，在集群在非“Available”状态下时，会返回此参数。
	Reason *string
	// 集群变为当前状态的原因的详细信息，在集群在非“Available”状态下时，会返回此参数。
	Message *string
}

// HostNetwork Node network parameters.
type HostNetwork struct {
	// Vpc 用于创建控制节点的VPC的ID。  获取方法如下： - 方法1：登录虚拟私有云服务的控制台界面，在虚拟私有云的详情页面查找VPC ID。 - 方法2：通过虚拟私有云服务的API接口查询。   [链接请参见[查询VPC列表](https://support.huaweicloud.com/api-vpc/vpc_api01_0003.html)](tag:hws)   [链接请参见[查询VPC列表](https://support.huaweicloud.com/intl/zh-cn/api-vpc/vpc_api01_0003.html)](tag:hws_hk)
	Vpc string
	// Subnet 用于创建控制节点的subnet的网络ID。获取方法如下：  - 方法1：登录虚拟私有云服务的控制台界面，单击VPC下的子网，进入子网详情页面，查找网络ID。 - 方法2：通过虚拟私有云服务的查询子网列表接口查询。   [链接请参见[查询子网列表](https://support.huaweicloud.com/api-vpc/vpc_subnet01_0003.html)](tag:hws)   [链接请参见[查询子网列表](https://support.huaweicloud.com/intl/zh-cn/api-vpc/vpc_subnet01_0003.html)](tag:hws_hk)
	Subnet string
	// SecurityGroup 集群默认的Node节点安全组ID，不指定该字段系统将自动为用户创建默认Node节点安全组，指定该字段时集群将绑定指定的安全组。Node节点安全组需要放通部分端口来保证正常通信。[详细设置请参考[集群安全组规则配置](https://support.huaweicloud.com/cce_faq/cce_faq_00265.html)。](tag:hws)[详细设置请参考[集群安全组规则配置](https://support.huaweicloud.com/intl/zh-cn/cce_faq/cce_faq_00265.html)。](tag:hws_hk)
	SecurityGroup *string
}

// ContainerNetwork Container network parameters.
type ContainerNetwork struct {
	// Mode 容器网络类型
	Mode string
	// Cidr 容器网络网段
	Cidr *string `json:"cidr,omitempty"`
	// Cidrs 容器网络网段列表
	Cidrs *[]string `json:"cidrs,omitempty"`
}

type ServiceNetwork struct {
	// IPv4CIDR kubernetes clusterIP IPv4 CIDR取值范围
	IPv4CIDR *string `json:"IPv4CIDR,omitempty"`
}

type ClusterExtendParam struct {
	// AlphaCceFixPoolMask 容器网络固定IP池掩码位数
	AlphaCceFixPoolMask *string
}

// CceClusterRsp2Cluster 转换cce集群响应数据到集群结构体
func CceClusterRsp2Cluster(rsp *model.ShowClusterResponse) *Cluster {
	return &Cluster{
		Metadata: &ClusterMetadata{
			Name: rsp.Metadata.Name,
			Uid:  rsp.Metadata.Uid,
		},
		Spec: &ClusterSpec{
			Type:        rsp.Spec.Type.Value(),
			Version:     rsp.Spec.Version,
			Description: rsp.Spec.Description,
			HostNetwork: &HostNetwork{
				Vpc:           rsp.Spec.HostNetwork.Vpc,
				Subnet:        rsp.Spec.HostNetwork.Subnet,
				SecurityGroup: rsp.Spec.HostNetwork.SecurityGroup,
			},
			ContainerNetwork: &ContainerNetwork{
				Mode: rsp.Spec.ContainerNetwork.Mode.Value(),
				Cidr: rsp.Spec.ContainerNetwork.Cidr,
				Cidrs: func() *[]string {
					cidrs := make([]string, 0)
					if rsp.Spec.ContainerNetwork.Cidrs != nil {
						for _, v := range *rsp.Spec.ContainerNetwork.Cidrs {
							cidrs = append(cidrs, v.Cidr)
						}
					}
					return &cidrs
				}(),
			},
			ServiceNetwork: &ServiceNetwork{
				IPv4CIDR: rsp.Spec.ServiceNetwork.IPv4CIDR,
			},
			KubeProxyMode: rsp.Spec.KubeProxyMode.Value(),
			ExtendParam: &ClusterExtendParam{
				AlphaCceFixPoolMask: rsp.Spec.ExtendParam.AlphaCceFixPoolMask,
			},
		},
		Status: &ClusterStatus{
			Phase:   rsp.Status.Phase,
			JobID:   rsp.Status.JobID,
			Reason:  rsp.Status.Reason,
			Message: rsp.Status.Message,
		},
	}
}

// CceClusterCreateRsp2Cluster 转换cce集群响应数据到集群结构体
func CceClusterCreateRsp2Cluster(rsp *model.CreateClusterResponse) *Cluster {
	return &Cluster{
		Metadata: &ClusterMetadata{
			Name: rsp.Metadata.Name,
			Uid:  rsp.Metadata.Uid,
		},
		Spec: &ClusterSpec{
			Type:        rsp.Spec.Type.Value(),
			Version:     rsp.Spec.Version,
			Description: rsp.Spec.Description,
			HostNetwork: &HostNetwork{
				Vpc:           rsp.Spec.HostNetwork.Vpc,
				Subnet:        rsp.Spec.HostNetwork.Subnet,
				SecurityGroup: rsp.Spec.HostNetwork.SecurityGroup,
			},
			ContainerNetwork: &ContainerNetwork{
				Mode: rsp.Spec.ContainerNetwork.Mode.Value(),
				Cidr: rsp.Spec.ContainerNetwork.Cidr,
				Cidrs: func() *[]string {
					cidrs := make([]string, 0)
					if rsp.Spec.ContainerNetwork.Cidrs != nil {
						for _, v := range *rsp.Spec.ContainerNetwork.Cidrs {
							cidrs = append(cidrs, v.Cidr)
						}
					}
					return &cidrs
				}(),
			},
			ServiceNetwork: &ServiceNetwork{
				IPv4CIDR: rsp.Spec.ServiceNetwork.IPv4CIDR,
			},
			KubeProxyMode: rsp.Spec.KubeProxyMode.Value(),
			ExtendParam: &ClusterExtendParam{
				AlphaCceFixPoolMask: rsp.Spec.ExtendParam.AlphaCceFixPoolMask,
			},
		},
		Status: &ClusterStatus{
			Phase:   rsp.Status.Phase,
			JobID:   rsp.Status.JobID,
			Reason:  rsp.Status.Reason,
			Message: rsp.Status.Message,
		},
	}
}

// AutopilotClusterRsp2Cluster 转换autopilot集群响应数据到集群结构体
func AutopilotClusterRsp2Cluster(rsp *model.ShowAutopilotClusterResponse) *Cluster {
	return &Cluster{
		Metadata: &ClusterMetadata{
			Name: rsp.Metadata.Name,
			Uid:  rsp.Metadata.Uid,
		},
		Spec: &ClusterSpec{
			Type:        rsp.Spec.Type.Value(),
			Version:     rsp.Spec.Version,
			Description: rsp.Spec.Description,
			HostNetwork: &HostNetwork{
				Vpc:    rsp.Spec.HostNetwork.Vpc,
				Subnet: rsp.Spec.HostNetwork.Subnet,
			},
			ContainerNetwork: &ContainerNetwork{
				Mode: rsp.Spec.ContainerNetwork.Mode.Value(),
			},
			ServiceNetwork: &ServiceNetwork{
				IPv4CIDR: rsp.Spec.ServiceNetwork.IPv4CIDR,
			},
			KubeProxyMode:   rsp.Spec.KubeProxyMode.Value(),
			EnableAutopilot: rsp.Spec.EnableAutopilot,
		},
		Status: &ClusterStatus{
			Phase:   rsp.Status.Phase,
			JobID:   rsp.Status.JobID,
			Reason:  rsp.Status.Reason,
			Message: rsp.Status.Message,
		},
	}
}

// AutopilotClusterCreateRsp2Cluster 转换autopilot集群响应数据到集群结构体
func AutopilotClusterCreateRsp2Cluster(rsp *model.CreateAutopilotClusterResponse) *Cluster {
	return &Cluster{
		Metadata: &ClusterMetadata{
			Name: rsp.Metadata.Name,
			Uid:  rsp.Metadata.Uid,
		},
		Spec: &ClusterSpec{
			Type:        rsp.Spec.Type.Value(),
			Version:     rsp.Spec.Version,
			Description: rsp.Spec.Description,
			HostNetwork: &HostNetwork{
				Vpc:    rsp.Spec.HostNetwork.Vpc,
				Subnet: rsp.Spec.HostNetwork.Subnet,
			},
			ContainerNetwork: &ContainerNetwork{
				Mode: rsp.Spec.ContainerNetwork.Mode.Value(),
			},
			ServiceNetwork: &ServiceNetwork{
				IPv4CIDR: rsp.Spec.ServiceNetwork.IPv4CIDR,
			},
			KubeProxyMode:   rsp.Spec.KubeProxyMode.Value(),
			EnableAutopilot: rsp.Spec.EnableAutopilot,
		},
		Status: &ClusterStatus{
			Phase:   rsp.Status.Phase,
			JobID:   rsp.Status.JobID,
			Reason:  rsp.Status.Reason,
			Message: rsp.Status.Message,
		},
	}
}
