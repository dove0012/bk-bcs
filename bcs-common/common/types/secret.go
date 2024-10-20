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

package types

// SecretDataItem secret detail
type SecretDataItem struct {
	// Path    string `json:"path,omitempty"` //mesos only
	Content string `json:"content"`
}

// BcsSecretType type for secret
type BcsSecretType string

const (
	// BcsSecretTypeOpaque xx
	BcsSecretTypeOpaque BcsSecretType = "Opaque"
	// BcsSecretTypeServiceAccountToken xx
	BcsSecretTypeServiceAccountToken BcsSecretType = "kubernetes.io/service-account-token" // nolint
	// BcsSecretTypeDockercfg xx
	BcsSecretTypeDockercfg BcsSecretType = "kubernetes.io/dockercfg"
	// BcsSecretTypeDockerConfigJson xx
	BcsSecretTypeDockerConfigJson BcsSecretType = "kubernetes.io/dockerconfigjson"
	// BcsSecretTypeBasicAuth xx
	BcsSecretTypeBasicAuth BcsSecretType = "kubernetes.io/basic-auth"
	// BcsSecretTypeSSHAuth xx
	BcsSecretTypeSSHAuth BcsSecretType = "kubernetes.io/ssh-auth"
	// BcsSecretTypeTLS xx
	BcsSecretTypeTLS BcsSecretType = "kubernetes.io/tls"
)

// BcsSecret bcs secret definition
type BcsSecret struct {
	TypeMeta `json:",inline"`
	// AppMeta    `json:",inline"`
	ObjectMeta `json:"metadata"`
	Type       BcsSecretType             `json:"type,omitempty"` // k8s only
	Data       map[string]SecretDataItem `json:"datas"`
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BcsSecret) DeepCopyInto(out *BcsSecret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make(map[string]SecretDataItem, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BcsConfigMapSpec.
func (in *BcsSecret) DeepCopy() *BcsSecret {
	if in == nil {
		return nil
	}
	out := new(BcsSecret)
	in.DeepCopyInto(out)
	return out
}
