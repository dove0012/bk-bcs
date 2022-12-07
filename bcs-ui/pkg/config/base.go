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

package config

import (
	"time"
	_ "time/tzdata" // tzdata TODO
)

const (
	// DevEnv TODO
	DevEnv = "dev"
	// StagEnv TODO
	StagEnv = "stag"
	// ProdEnv TODO
	ProdEnv = "prod"
)

// BaseConf :
type BaseConf struct {
	TimeZone     string         `yaml:"time_zone"`
	LanguageCode string         `yaml:"language_code"`
	RunEnv       string         `yaml:"run_env"`
	Location     *time.Location `yaml:"-"`
}

// Init :
func (c *BaseConf) Init() error {
	var err error
	c.TimeZone = "Asia/Shanghai"
	c.LanguageCode = ""
	c.RunEnv = DevEnv
	c.Location, err = time.LoadLocation(c.TimeZone)
	if err != nil {
		return err
	}
	return nil
}