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

package bcs_system

// AvailableNodeMetrics 蓝鲸监控节点的metrics
var AvailableNodeMetrics = []string{
	"bcs:cluster:cpu:usage",
	"bcs:cluster:cpu:total",
	"bcs:cluster:cpu:used",
	"bcs:cluster:memory:total",
	"bcs:cluster:memory:used",
	"bcs:cluster:memory:usage",
	"bcs:cluster:disk:total",
	"bcs:cluster:disk:used",
	"bcs:cluster:disk:usage",
	"bcs:node:info",
	"bcs:node:cpu:usage",
	"bcs:node:disk:usage",
	"bcs:node:diskio:usage",
	"bcs:node:memory:usage",
	"bcs:node:container_count",
	"bcs:node:pod_count",
	"bcs:node:network_transmit",
	"bcs:node:network_receive",
	"bcs:pod:cpu_usage",
	"bcs:pod:memory_usage",
	"bcs:pod:network_transmit",
	"bcs:pod:network_receive",
	"bcs:container:cpu_usage",
	"bcs:container:memory_usage",
	"bcs:container:disk_read",
	"bcs:container:disk_write",
}