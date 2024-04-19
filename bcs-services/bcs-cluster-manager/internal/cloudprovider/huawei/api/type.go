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

package api

type CreateClusterRequest struct {
	Body *Cluster `json:"body,omitempty"`
}

// Cluster
type Cluster struct {
	// API类型，固定值“Cluster”或“cluster”，该值不可修改。
	Kind string `json:"kind"`
	// API版本，固定值“v3”，该值不可修改。
	ApiVersion string           `json:"apiVersion"`
	Metadata   *ClusterMetadata `json:"metadata"`
	Spec       *ClusterSpec     `json:"spec"`
	Status     *ClusterStatus   `json:"status,omitempty"`
}

// ClusterMetadata 可以通过 annotations[\"cluster.install.addons/install\"] 来指定创建集群时需要安装的插件，格式形如 ``` [   {     \"addonTemplateName\": \"autoscaler\",     \"version\": \"1.15.3\",     \"values\": {       \"flavor\": {         \"description\": \"Has only one instance\",         \"name\": \"Single\",         \"replicas\": 1,         \"resources\": [           {             \"limitsCpu\": \"100m\",             \"limitsMem\": \"300Mi\",             \"name\": \"autoscaler\",             \"requestsCpu\": \"100m\",             \"requestsMem\": \"300Mi\"           }         ]       },       \"custom\": {         \"coresTotal\": 32000,         \"maxEmptyBulkDeleteFlag\": 10,         \"maxNodesTotal\": 1000,         \"memoryTotal\": 128000,         \"scaleDownDelayAfterAdd\": 10,         \"scaleDownDelayAfterDelete\": 10,         \"scaleDownDelayAfterFailure\": 3,         \"scaleDownEnabled\": false,         \"scaleDownUnneededTime\": 10,         \"scaleDownUtilizationThreshold\": 0.5,         \"scaleUpCpuUtilizationThreshold\": 1,         \"scaleUpMemUtilizationThreshold\": 1,         \"scaleUpUnscheduledPodEnabled\": true,         \"scaleUpUtilizationEnabled\": true,         \"tenant_id\": \"47eb1d64cbeb45cfa01ae20af4f4b563\",         \"unremovableNodeRecheckTimeout\": 5       }     }   } ] ```
type ClusterMetadata struct {
	// 集群名称。  命名规则：以小写字母开头，由小写字母、数字、中划线(-)组成，长度范围4-128位，且不能以中划线(-)结尾。
	Name string `json:"name"`
	// 集群ID，资源唯一标识，创建成功后自动生成，填写无效。在创建包周期集群时，响应体不返回集群ID。
	Uid *string `json:"uid,omitempty"`
	// 集群显示名，用于在 CCE 界面显示，该名称创建后可修改。  命名规则：以小写字母开头，由小写字母、数字、中划线(-)组成，长度范围4-128位，且不能以中划线(-)结尾。  显示名和其他集群的名称、显示名不可以重复。  在创建集群、更新集群请求体中，集群显示名alias未指定或取值为空，表示与集群名称name一致。在查询集群等响应体中，集群显示名alias将必然返回，未配置时将返回集群名称name。
	Alias *string `json:"alias,omitempty"`
	// 集群注解，由key/value组成：  ``` \"annotations\": {    \"key1\" : \"value1\",    \"key2\" : \"value2\" } ```  >    - Annotations不用于标识和选择对象。Annotations中的元数据可以是small或large，structured或unstructured，并且可以包括标签不允许使用的字符。 >    - 该字段不会被数据库保存，当前仅用于指定集群待安装插件。 >    - 可通过加入\"cluster.install.addons.external/install\":\"[{\"addonTemplateName\":\"icagent\"}]\"的键值对在创建集群时安装ICAgent。
	Annotations map[string]string `json:"annotations,omitempty"`
	// 集群标签，key/value对格式。  >  该字段值由系统自动生成，用于升级时前端识别集群支持的特性开关，用户指定无效。
	Labels map[string]string `json:"labels,omitempty"`
	// 集群创建时间
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`
	// 集群更新时间
	UpdateTimestamp *string `json:"updateTimestamp,omitempty"`
}

// ClusterSpec 集群参数定义。
type ClusterSpec struct {
	// 集群类别： - CCE：CCE集群   CCE集群支持虚拟机与裸金属服务器混合、GPU、NPU等异构节点的混合部署，基于高性能网络模型提供全方位、多场景、安全稳定的容器运行环境。 [- Turbo: CCE Turbo集群。   全面基于云原生基础设施构建的云原生2.0的容器引擎服务，具备软硬协同、网络无损、安全可靠、调度智能的优势，为用户提供一站式、高性价比的全新容器服务体验。](tag:hws,hws_hk,dt,hcs,g42,sbc)
	Category string `json:"category,omitempty"`
	// 集群Master节点架构：  - VirtualMachine：Master节点为x86架构服务器 [- ARM64: Master节点为鲲鹏（ARM架构）服务器](tag:hws,hws_hk,hcs)
	Type string `json:"type,omitempty"`
	// 集群规格，当集群为v1.15及以上版本时支持创建后变更，详情请参见[变更集群规格](ResizeCluster.xml)。请按实际业务需求进行选择： - cce.s1.small: 小规模单控制节点CCE集群（最大50节点） - cce.s1.medium: 中等规模单控制节点CCE集群（最大200节点） - cce.s2.small: 小规模多控制节点CCE集群（最大50节点） - cce.s2.medium: 中等规模多控制节点CCE集群（最大200节点） - cce.s2.large: 大规模多控制节点CCE集群（最大1000节点） - cce.s2.xlarge: 超大规模多控制节点CCE集群（最大2000节点）  >    关于规格参数中的字段说明如下： >    - s1：单控制节点的集群，控制节点数为1。单控制节点故障后，集群将不可用，但已运行工作负载不受影响。 >    - s2：多控制节点的集群，即高可用集群，控制节点数为3。当某个控制节点故障时，集群仍然可用。 >    [- dec：表示专属云的CCE集群规格。例如cce.dec.s1.small表示小规模单控制节点的专属云CCE集群（最大50节点）。](tag:hws,hws_hk) >    - small：表示集群支持管理的最大节点规模为50节点。 >    - medium：表示集群支持管理的最大节点规模为200节点。 >    - large：表示集群支持管理的最大节点规模为1000节点。 >    - xlarge：表示集群支持管理的最大节点规模为2000节点。
	Flavor string `json:"flavor"`
	// CCE Autopilot集群开关： - true：创建集群为CCE Autopilot集群
	EnableAutopilot *bool `json:"enableAutopilot,omitempty"`
	// 集群版本，与Kubernetes社区基线版本保持一致，建议选择最新版本。  在CCE控制台支持创建两种最新版本的集群。可登录CCE控制台创建集群，在“版本”处获取到集群版本。 其它集群版本，当前仍可通过api创建，但后续会逐渐下线，具体下线策略请关注CCE官方公告。  >    - 若不配置，默认创建最新版本的集群。 >    - 若指定集群基线版本但是不指定具体r版本，则系统默认选择对应集群版本的最新r版本。建议不指定具体r版本由系统选择最新版本。 [>    - Turbo集群支持1.19及以上版本商用。](tag:hws,hws_hk,dt) [>    - Turbo集群支持1.23及以上版本商用。](tag:hcs,g42,sbc)
	Version *string `json:"version,omitempty"`
	// CCE集群平台版本号，表示集群版本(version)下的内部版本。用于跟踪某一集群版本内的迭代，集群版本内唯一，跨集群版本重新计数。不支持用户指定，集群创建时自动选择对应集群版本的最新平台版本。  platformVersion格式为：cce.X.Y - X: 表示内部特性版本。集群版本中特性或者补丁修复，或者OS支持等变更场景。其值从1开始单调递增。 - Y: 表示内部特性版本的补丁版本。仅用于特性版本上线后的软件包更新，不涉及其他修改。其值从0开始单调递增。
	PlatformVersion *string `json:"platformVersion,omitempty"`
	// 集群描述，对于集群使用目的的描述，可根据实际情况自定义，默认为空。集群创建成功后可通过接口[更新指定的集群](cce_02_0240.xml)来做出修改，也可在CCE控制台中对应集群的“集群详情”下的“描述”处进行修改。仅支持utf-8编码。
	Description *string `json:"description,omitempty"`
	// 集群的API Server服务端证书中的自定义SAN（Subject Alternative Name）字段，遵从SSL标准X509定义的格式规范。  1. 不允许出现同名重复。 2. 格式符合IP和域名格式。  示例: ``` SAN 1: DNS Name=example.com SAN 2: DNS Name=www.example.com SAN 3: DNS Name=example.net SAN 4: IP Address=93.184.216.34 ```
	CustomSan *[]string `json:"customSan,omitempty"`
	// 集群是否使用IPv6模式，1.15版本及以上支持。
	Ipv6enable *bool `json:"ipv6enable,omitempty"`
	// CCE Turbo集群
	OffloadCluster   *bool             `json:"offloadCluster,omitempty"`
	HostNetwork      *HostNetwork      `json:"hostNetwork"`
	ContainerNetwork *ContainerNetwork `json:"containerNetwork"`
	EniNetwork       *EniNetwork       `json:"eniNetwork,omitempty"`
	ServiceNetwork   *ServiceNetwork   `json:"serviceNetwork,omitempty"`
	Authentication   *Authentication   `json:"authentication,omitempty"`
	// 集群的计费方式。 - 0: 按需计费 [- 1: 包周期](tag:hws,hws_hk)  默认为“按需计费”。
	BillingMode int32 `json:"billingMode,omitempty"`
	// 控制节点的高级配置
	Masters *[]MasterSpec `json:"masters,omitempty"`
	// 服务网段参数，kubernetes clusterIP取值范围，1.11.7版本及以上支持。创建集群时如若未传参，默认为\"10.247.0.0/16\"。该参数废弃中，推荐使用新字段serviceNetwork，包含IPv4服务网段。
	KubernetesSvcIpRange *string `json:"kubernetesSvcIpRange,omitempty"`
	// 集群资源标签
	ClusterTags []ResourceTag `json:"clusterTags,omitempty"`
	// 服务转发模式，支持以下两种实现：  - iptables：社区传统的kube-proxy模式，完全以iptables规则的方式来实现service负载均衡。该方式最主要的问题是在服务多的时候产生太多的iptables规则，非增量式更新会引入一定的时延，大规模情况下有明显的性能问题。 - ipvs：主导开发并在社区获得广泛支持的kube-proxy模式，采用增量式更新，吞吐更高，速度更快，并可以保证service更新期间连接保持不断开，适用于大规模场景。
	KubeProxyMode string `json:"kubeProxyMode,omitempty"`
	// 可用区（仅查询返回字段）。  [CCE支持的可用区请参考[地区和终端节点](https://developer.huaweicloud.com/endpoint?CCE)](tag:hws)  [CCE支持的可用区请参考[地区和终端节点](https://developer.huaweicloud.com/intl/zh-cn/endpoint?CCE)](tag:hws_hk)
	Az          *string             `json:"az,omitempty"`
	ExtendParam *ClusterExtendParam `json:"extendParam,omitempty"`
	// 支持Istio
	SupportIstio *bool `json:"supportIstio,omitempty"`
	// 集群控制节点系统盘、数据盘加密。默认使用AES_256加密算法。CCE、Turbo集群1.25及以上版本开始支持。集群创建后不支持修改。开启后存在一定的磁盘读写性能损耗。
	EnableMasterVolumeEncryption *bool `json:"enableMasterVolumeEncryption,omitempty"`
	// 覆盖集群默认组件配置  若指定了不支持的组件或组件不支持的参数，该配置项将被忽略。  当前支持的可配置组件及其参数详见 [[配置管理](https://support.huaweicloud.com/usermanual-cce/cce_10_0213.html)](tag:hws) [[配置管理](https://support.huaweicloud.com/intl/zh-cn/usermanual-cce/cce_10_0213.html)](tag:hws_hk)
	ConfigurationsOverride *[]PackageConfiguration `json:"configurationsOverride,omitempty"`
}

// ClusterStatus
type ClusterStatus struct {
	// 集群状态，取值如下 - Available：可用，表示集群处于正常状态。 - Unavailable：不可用，表示集群异常，需手动删除。 - ScalingUp：扩容中，表示集群正处于扩容过程中。 - ScalingDown：缩容中，表示集群正处于缩容过程中。 - Creating：创建中，表示集群正处于创建过程中。 - Deleting：删除中，表示集群正处于删除过程中。 - Upgrading：升级中，表示集群正处于升级过程中。 - Resizing：规格变更中，表示集群正处于变更规格中。 - RollingBack：回滚中，表示集群正处于回滚过程中。 - RollbackFailed：回滚异常，表示集群回滚异常。 - Hibernating：休眠中，表示集群正处于休眠过程中。 - Hibernation：已休眠，表示集群正处于休眠状态。 - Awaking：唤醒中，表示集群正处于从休眠状态唤醒的过程中。 - Empty：集群无任何资源（已废弃）
	Phase *string `json:"phase,omitempty"`
	// 任务ID,集群当前状态关联的任务ID。当前支持: - 创建集群时返回关联的任务ID，可通过任务ID查询创建集群的附属任务信息； - 删除集群或者删除集群失败时返回关联的任务ID，此字段非空时，可通过任务ID查询删除集群的附属任务信息。 > 任务信息具有一定时效性，仅用于短期跟踪任务进度，请勿用于集群状态判断等额外场景。
	JobID *string `json:"jobID,omitempty"`
	// 集群变为当前状态的原因，在集群在非“Available”状态下时，会返回此参数。
	Reason *string `json:"reason,omitempty"`
	// 集群变为当前状态的原因的详细信息，在集群在非“Available”状态下时，会返回此参数。
	Message *string `json:"message,omitempty"`
	// 集群中 kube-apiserver 的访问地址。
	Endpoints *[]ClusterEndpoints `json:"endpoints,omitempty"`
	// CBC资源锁定
	IsLocked *bool `json:"isLocked,omitempty"`
	// CBC资源锁定场景
	LockScene *string `json:"lockScene,omitempty"`
	// 锁定资源
	LockSource *string `json:"lockSource,omitempty"`
	// 锁定的资源ID
	LockSourceId *string `json:"lockSourceId,omitempty"`
	// 删除配置状态（仅删除请求响应包含）
	DeleteOption *interface{} `json:"deleteOption,omitempty"`
	// 删除状态信息（仅删除请求响应包含）
	DeleteStatus *interface{} `json:"deleteStatus,omitempty"`
}

// HostNetwork Node network parameters.
type HostNetwork struct {
	// 用于创建控制节点的VPC的ID。  获取方法如下： - 方法1：登录虚拟私有云服务的控制台界面，在虚拟私有云的详情页面查找VPC ID。 - 方法2：通过虚拟私有云服务的API接口查询。   [链接请参见[查询VPC列表](https://support.huaweicloud.com/api-vpc/vpc_api01_0003.html)](tag:hws)   [链接请参见[查询VPC列表](https://support.huaweicloud.com/intl/zh-cn/api-vpc/vpc_api01_0003.html)](tag:hws_hk)
	Vpc string `json:"vpc"`
	// 用于创建控制节点的subnet的网络ID。获取方法如下：  - 方法1：登录虚拟私有云服务的控制台界面，单击VPC下的子网，进入子网详情页面，查找网络ID。 - 方法2：通过虚拟私有云服务的查询子网列表接口查询。   [链接请参见[查询子网列表](https://support.huaweicloud.com/api-vpc/vpc_subnet01_0003.html)](tag:hws)   [链接请参见[查询子网列表](https://support.huaweicloud.com/intl/zh-cn/api-vpc/vpc_subnet01_0003.html)](tag:hws_hk)
	Subnet string `json:"subnet"`
	// 集群默认的Node节点安全组ID，不指定该字段系统将自动为用户创建默认Node节点安全组，指定该字段时集群将绑定指定的安全组。Node节点安全组需要放通部分端口来保证正常通信。[详细设置请参考[集群安全组规则配置](https://support.huaweicloud.com/cce_faq/cce_faq_00265.html)。](tag:hws)[详细设置请参考[集群安全组规则配置](https://support.huaweicloud.com/intl/zh-cn/cce_faq/cce_faq_00265.html)。](tag:hws_hk)
	SecurityGroup string `json:"SecurityGroup,omitempty"`
}

// ContainerNetwork Container network parameters.
type ContainerNetwork struct {
	// 容器网络类型（只可选择其一） - overlay_l2：容器隧道网络，通过OVS（OpenVSwitch）为容器构建的overlay_l2网络。 - vpc-router：VPC网络，使用ipvlan和自定义VPC路由为容器构建的Underlay的l2网络。 [- eni：云原生网络2.0，深度整合VPC原生ENI弹性网卡能力，采用VPC网段分配容器地址，支持ELB直通容器，享有高性能，创建CCE Turbo集群时指定。](tag:hws,hws_hk,dt,hcs,g42,sbc)
	Mode string `json:"mode"`
	// 容器网络网段，建议使用网段10.0.0.0/12~19，172.16.0.0/16~19，192.168.0.0/16~19，如存在网段冲突，将会报错。  此参数在集群创建后不可更改，请谨慎选择。（已废弃，如填写cidrs将忽略该cidr）
	Cidr *string `json:"cidr,omitempty"`
	// 容器网络网段列表。1.21及新版本集群使用cidrs字段，当集群网络类型为vpc-router类型时，支持多个容器网段，最多配置20个；1.21之前版本若使用cidrs字段，则取值cidrs数组中的第一个cidr元素作为容器网络网段地址。  此参数在集群创建后不可更改，请谨慎选择。
	Cidrs *[]string `json:"cidrs,omitempty"`
}

// EniNetwork ENI网络配置，创建集群指定使用云原生网络2.0网络模式时必填subnets和eniSubnetId其中一个字段(eniSubnetCIDR可选，若填写了会校验是否合法)，1.19.10及新版本集群使用subnets字段，1.19.8及老版本若使用subnets字段，则取值subnets数组中的第一个子网ID作为容器地址使用的子网ID。
type EniNetwork struct {
	// ENI所在子网的IPv4子网ID(暂不支持IPv6,废弃中)。获取方法如下：  - 方法1：登录虚拟私有云服务的控制台界面，单击VPC下的子网，进入子网详情页面，查找IPv4子网ID。 - 方法2：通过虚拟私有云服务的查询子网列表接口查询。   [链接请参见[查询子网列表](https://support.huaweicloud.com/api-vpc/vpc_subnet01_0003.html)](tag:hws)   [链接请参见[查询子网列表](https://support.huaweicloud.com/intl/zh-cn/api-vpc/vpc_subnet01_0003.html)](tag:hws_hk)
	EniSubnetId string `json:"eniSubnetId"`
	// ENI子网CIDR(废弃中)
	EniSubnetCIDR string `json:"eniSubnetCIDR"`
	// IPv4子网ID列表
	Subnets []string `json:"subnets"`
}

type ServiceNetwork struct {
	// kubernetes clusterIP IPv4 CIDR取值范围。创建集群时若未传参，默认为\"10.247.0.0/16\"。
	IPv4CIDR string `json:"IPv4CIDR,omitempty"`
}

// Authentication
type Authentication struct {
	// 集群认证模式。 - kubernetes 1.11及之前版本的集群支持“x509”、“rbac”和“authenticating_proxy”，默认取值为“x509”。 - kubernetes 1.13及以上版本的集群支持“rbac”和“authenticating_proxy”，默认取值为“rbac”。
	Mode                *string              `json:"mode,omitempty"`
	AuthenticatingProxy *AuthenticatingProxy `json:"authenticatingProxy,omitempty"`
}

// AuthenticatingProxy authenticatingProxy模式相关配置。认证模式为authenticating_proxy时必选
type AuthenticatingProxy struct {
	// authenticating_proxy模式配置的x509格式CA证书(base64编码)。当集群认证模式为authenticating_proxy时，此项必须填写。   最大长度：1M
	Ca *string `json:"ca,omitempty"`
	// authenticating_proxy模式配置的x509格式CA证书签发的客户端证书，用于kube-apiserver到扩展apiserver的认证。(base64编码)。当集群认证模式为authenticating_proxy时，此项必须填写。
	Cert *string `json:"cert,omitempty"`
	// authenticating_proxy模式配置的x509格式CA证书签发的客户端证书时对应的私钥，用于kube-apiserver到扩展apiserver的认证。Kubernetes集群使用的私钥尚不支持密码加密，请使用未加密的私钥。(base64编码)。当集群认证模式为authenticating_proxy时，此项必须填写。
	PrivateKey *string `json:"privateKey,omitempty"`
}

// MasterSpec master的配置，支持指定可用区、规格和故障域。若指定故障域，则必须所有master节点都需要指定故障字段。
type MasterSpec struct {
	// 可用区
	AvailabilityZone *string `json:"availabilityZone,omitempty"`
	// 规格
	Flavor *string `json:"flavor,omitempty"`
	// 故障域。 1. 指定该字段需要当前系统已开启故障域特性，否则校验失败。 2. 仅单az场景支持且必须显式指定az。
	FaultDomain *string `json:"faultDomain,omitempty"`
}

// ResourceTag CCE资源标签
type ResourceTag struct {
	// Key值。 - 不能为空，最多支持128个字符 - 可用UTF-8格式表示的汉字、字母、数字和空格 - 支持部分特殊字符：_.:/=+-@ - 不能以\"\\_sys\\_\"开头
	Key string `json:"key,omitempty"`
	// Value值。 - 可以为空但不能缺省，最多支持255个字符 - 可用UTF-8格式表示的汉字、字母、数字和空格 - 支持部分特殊字符：_.:/=+-@
	Value string `json:"value,omitempty"`
}

type ClusterExtendParam struct {
	// 集群控制节点可用区配置。  [CCE支持的可用区请参考[地区和终端节点](https://developer.huaweicloud.com/endpoint?CCE)](tag:hws) [CCE支持的可用区请参考[地区和终端节点](https://developer.huaweicloud.com/intl/zh-cn/endpoint?CCE)](tag:hws_hk)    - multi_az：多可用区，可选。仅使用高可用集群时才可以配置多可用区。 - 专属云计算池可用区：用于指定专属云可用区部署集群控制节点。如果需配置专属CCE集群，该字段为必选。
	ClusterAZ *string `json:"clusterAZ,omitempty"`
	// 用于指定控制节点的系统盘和数据盘使用专属分布式存储，未指定或者值为空时，默认使用EVS云硬盘。  如果配置专属CCE集群，该字段为必选，请按照如下格式设置：  ``` <rootVol.dssPoolID>.<rootVol.volType>;<dataVol.dssPoolID>.<dataVol.volType> ```  字段说明： - rootVol为系统盘；dataVol为数据盘； - dssPoolID为专属分布式存储池ID； - volType为专属分布式存储池的存储类型，如SAS、SSD。  样例：c950ee97-587c-4f24-8a74-3367e3da570f.sas;6edbc2f4-1507-44f8-ac0d-eed1d2608d38.ssd  > 非专属CCE集群不支持配置该字段。
	DssMasterVolumes *string `json:"dssMasterVolumes,omitempty"`
	// 集群所属的企业项目ID。 >   - 需要开通企业项目功能后才可配置企业项目。 >   - 集群所属的企业项目与集群下所关联的其他云服务资源所属的企业项目必须保持一致。
	EnterpriseProjectId *string `json:"enterpriseProjectId,omitempty"`
	// 服务转发模式，支持以下两种实现：  - iptables：社区传统的kube-proxy模式，完全以iptables规则的方式来实现service负载均衡。该方式最主要的问题是在服务多的时候产生太多的iptables规则，非增量式更新会引入一定的时延，大规模情况下有明显的性能问题 - ipvs：主导开发并在社区获得广泛支持的kube-proxy模式，采用增量式更新，吞吐更高，速度更快，并可以保证service更新期间连接保持不断开，适用于大规模场景。  > 此参数已废弃，若同时指定此参数和ClusterSpec下的kubeProxyMode，以ClusterSpec下的为准。
	KubeProxyMode *string `json:"kubeProxyMode,omitempty"`
	// master 弹性公网IP
	ClusterExternalIP *string `json:"clusterExternalIP,omitempty"`
	// 容器网络固定IP池掩码位数，仅vpc-router网络支持。  该参数决定节点可分配容器IP数量，与创建节点时设置的maxPods参数共同决定节点最多可以创建多少个Pod， 具体请参见[节点最多可以创建多少Pod](maxPods.xml)。   整数字符传取值范围: 24 ~ 28
	AlphaCceFixPoolMask *string `json:"alpha.cce/fixPoolMask,omitempty"`
	// 专属CCE集群指定可控制节点的规格。
	DecMasterFlavor *string `json:"decMasterFlavor,omitempty"`
	// 集群默认Docker的UmaskMode配置，可取值为secure或normal，不指定时默认为normal。
	DockerUmaskMode *string `json:"dockerUmaskMode,omitempty"`
	// 集群CPU管理策略。取值为none（或空值）或static，默认为none（或空值）。 - none(或空值)：关闭工作负载实例独占CPU核的功能，优点是CPU共享池的可分配核数较多 - static：支持给节点上的工作负载实例配置CPU独占，适用于对CPU缓存和调度延迟敏感的工作负载[，Turbo集群下仅对普通容器节点有效，安全容器节点无效](tag:hws,hws_hk,dt,g42,sbc)。
	KubernetesIoCpuManagerPolicy *string `json:"kubernetes.io/cpuManagerPolicy,omitempty"`
	// 订单ID，集群付费类型为自动付费包周期类型时，响应中会返回此字段(仅创建场景涉及)。
	OrderID *string `json:"orderID,omitempty"`
	// - month：月 - year：年 > 作为请求参数，billingMode为1（包周期）时生效，且为必选。 > 作为响应参数，仅在创建包周期集群时返回。
	PeriodType *string `json:"periodType,omitempty"`
	// 订购周期数，取值范围： - periodType=month（周期类型为月）时，取值为[1-9]。 - periodType=year（周期类型为年）时，取值为1-3。 > 作为请求参数，billingMode为1时生效，且为必选。 > 作为响应参数，仅在创建包周期集群时返回。
	PeriodNum *int32 `json:"periodNum,omitempty"`
	// 是否自动续订 - “true”：自动续订 - “false”：不自动续订 > billingMode为1时生效，不填写此参数时默认不会自动续费。
	IsAutoRenew *string `json:"isAutoRenew,omitempty"`
	// 是否自动扣款 - “true”：自动扣款 - “false”：不自动扣款 > billingMode为1时生效，不填写此参数时默认不会自动扣款。
	IsAutoPay *string `json:"isAutoPay,omitempty"`
	// 记录集群通过何种升级方式升级到当前版本。
	Upgradefrom *string `json:"upgradefrom,omitempty"`
}

type PackageConfiguration struct {
	// 组件名称
	Name *string `json:"name,omitempty"`
	// 组件配置项
	Configurations *[]ConfigurationItem `json:"configurations,omitempty"`
}

type ConfigurationItem struct {
	// 组件配置项名称
	Name *string `json:"name,omitempty"`
	// 组件配置项值
	Value *interface{} `json:"value,omitempty"`
}

type ClusterEndpoints struct {
	// 集群中 kube-apiserver 的访问地址
	Url *string `json:"url,omitempty"`
	// 集群访问地址的类型 - Internal：用户子网内访问的地址 - External：公网访问的地址
	Type *string `json:"type,omitempty"`
}
