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

// Package business xxx
package business

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/internal/cloudprovider"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/internal/cloudprovider/google/api"
)

// GetRuntimeInfo get runtime info
func GetUpgradeVersions(clusterName string, opt *cloudprovider.CommonOption) (string, []string, string, []string, error) {
	client, err := api.NewContainerServiceClient(opt)
	if err != nil {
		return "", nil, "", nil, fmt.Errorf("create google client failed, err %s", err.Error())
	}

	cluster, err := client.GetCluster(context.Background(), clusterName)
	if err != nil {
		return "", nil, "", nil, fmt.Errorf("get cluster failed, err %s", err.Error())
	}

	conf, err := client.GetUpgradeVersions(context.Background(), clusterName)
	if err != nil {
		return "", nil, "", nil, fmt.Errorf("get upgrade versions failed, err %s", err.Error())
	}

	// 获取当前集群的master和node版本
	curMasterVer := cluster.CurrentMasterVersion
	curNodeVer := cluster.CurrentNodeVersion

	blog.Infof("get upgrade versions for cluster %s, current master version %s, valid master versions %v, valid node versions %v",
		clusterName, curMasterVer, conf.ValidMasterVersions, conf.ValidNodeVersions)

	var (
		validMasterVersions = make([]string, 0)
		validNodeVersions   = make([]string, 0)
	)

	masterPreVers, err := getPreVer(curMasterVer)
	if err != nil {
		return "", nil, "", nil, fmt.Errorf("get pre version for master version %s failed, err %s", curMasterVer, err)
	}

	nodePreVer, err := getPreVer(curNodeVer)
	if err != nil {
		return "", nil, "", nil, fmt.Errorf("get pre version for node version %s failed, err %s", curNodeVer, err)
	}

	for _, ver := range conf.ValidMasterVersions {
		perVers, err := getPreVer(ver)
		if err != nil {
			return "", nil, "", nil, fmt.Errorf("get pre version %s failed, err %s", ver, err)
		}

		if compareVer(masterPreVers, perVers) > 0 {
			break
		}

		validMasterVersions = append(validMasterVersions, ver)
		if ver == curMasterVer {
			break
		}
	}

	for _, ver := range conf.ValidNodeVersions {
		perVers, err := getPreVer(ver)
		if err != nil {
			return "", nil, "", nil, fmt.Errorf("get pre version %s failed, err %s", ver, err)
		}

		if compareVer(nodePreVer, perVers) > 0 {
			break
		}

		validNodeVersions = append(validNodeVersions, ver)
		if ver == curNodeVer {
			break
		}
	}

	return curMasterVer, validMasterVersions, curNodeVer, validNodeVersions, nil
}

func getPreVer(ver string) ([]int, error) {
	preVers := strings.Split(ver, "-gke.")
	if len(preVers) != 2 {
		return nil, fmt.Errorf("invalid version format %s", ver)
	}

	preVer := strings.Split(preVers[0], ".")
	if len(preVer) != 3 {
		return nil, fmt.Errorf("invalid version format %s", ver)
	}

	vers := make([]int, 0)
	for _, ver := range preVer {
		i, err := strconv.Atoi(ver)
		if err != nil {
			return nil, fmt.Errorf("invalid version format %s", ver)
		}

		vers = append(vers, i)
	}

	return vers, nil
}

func compareVer(ver1, ver2 []int) int {
	for i := range ver1 {
		if ver1[i] > ver2[i] {
			return 1
		} else if ver1[i] < ver2[i] {
			return -1
		}
	}

	return 0
}
