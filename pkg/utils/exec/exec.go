// Copyright (c) 2018 Huawei Technologies Co., Ltd. All Rights Reserved.
//
//    Licensed under the Apache License, Version 2.0 (the "License"); you may
//    not use this file except in compliance with the License. You may obtain
//    a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//    WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//    License for the specific language governing permissions and limitations
//    under the License.

package exec

import (
	"os/exec"
	"strings"

	log "github.com/golang/glog"
)

type Executer interface {
	Run(name string, arg ...string) (string, error)
}

func Run(name string, arg ...string) (string, error) {
	log.V(5).Infof("Command: %s %s", name, strings.Join(arg, " "))
	info, err := exec.Command(name, arg...).Output()
	if err != nil {
		log.Errorf("Execute command failed, error: %v", err)
		return "", err
	}
	log.V(5).Infof("Command Result:\n%s", string(info))
	return string(info), nil
}

func NewBaseExecuter() Executer {
	return &BaseExecuter{}
}

type BaseExecuter struct{}

func (r *BaseExecuter) Run(name string, arg ...string) (string, error) {
	return Run(name, arg...)
}

func NewRootExecuter() Executer {
	return &RootExeucter{}
}

type RootExeucter struct{}

func (r *RootExeucter) Run(name string, arg ...string) (string, error) {
	// TODO: Add root wrapper here
	return Run(name, arg...)
}
