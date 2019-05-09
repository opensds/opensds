// Copyright (c) 2017 Huawei Technologies Co., Ltd. All Rights Reserved.
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

package targets

const (
	iscsiTgtPrefix  = "iqn.2017-10.io.opensds:"
	nvmeofTgtPrefix = "nqn.2019-01.com.opensds:nvme:"
	iscsiAccess     = "iscsi"
	nvmeofAccess    = "nvmeof"
)

// Target is an interface for exposing some operations of different targets,
// currently support iscsiTarget.
type Target interface {
	CreateExport(volId, path, hostIp, initiator string, chapAuth []string) (map[string]interface{}, error)

	RemoveExport(volId string) error
}

// NewTarget method creates a new target based on its type.
func NewTarget(bip string, tgtConfDir string, access string) Target {
	switch access {
	case iscsiAccess:
		return &iscsiTarget{
			ISCSITarget: NewISCSITarget(bip, tgtConfDir),
		}
	case nvmeofAccess:
		return &nvmeofTarget{
			NvmeofTarget: NewNvmeofTarget(bip, tgtConfDir),
		}
	default:
		return nil
	}
}

type iscsiTarget struct {
	ISCSITarget
}

func (t *iscsiTarget) CreateExport(volId, path, hostIp, initiator string, chapAuth []string) (map[string]interface{}, error) {
	tgtIqn := iscsiTgtPrefix + volId
	if err := t.CreateISCSITarget(volId, tgtIqn, path, hostIp, initiator, chapAuth); err != nil {
		return nil, err
	}
	lunId := t.GetLun(path)
	conn := map[string]interface{}{
		"targetDiscovered": true,
		"targetIQN":        []string{tgtIqn},
		"targetPortal":     []string{t.ISCSITarget.(*tgtTarget).BindIp + ":3260"},
		"discard":          false,
		"targetLun":        lunId,
	}
	if len(chapAuth) == 2 {
		conn["authMethod"] = "chap"
		conn["authUserName"] = chapAuth[0]
		conn["authPassword"] = chapAuth[1]
	}
	return conn, nil
}

func (t *iscsiTarget) RemoveExport(volId string) error {
	tgtIqn := iscsiTgtPrefix + volId
	return t.RemoveISCSITarget(volId, tgtIqn)
}

type nvmeofTarget struct {
	NvmeofTarget
}

func (t *nvmeofTarget) CreateExport(volId, path, hostIp, initiator string, chapAuth []string) (map[string]interface{}, error) {
	tgtNqn := nvmeofTgtPrefix + volId
	if err := t.CreateNvmeofTarget(volId, tgtNqn, path, hostIp, initiator, chapAuth); err != nil {
		return nil, err
	}
	conn := map[string]interface{}{
		"targetDiscovered": true,
		"targetNQN":        tgtNqn,
		"targetIP":         t.NvmeofTarget.(*NvmeoftgtTarget).BindIp,
		"targetPort":       "4420",
		"discard":          false,
	}
	if len(chapAuth) == 2 {
		conn["authMethod"] = "chap"
		conn["authUserName"] = chapAuth[0]
		conn["authPassword"] = chapAuth[1]
	}
	return conn, nil
}

func (t *nvmeofTarget) RemoveExport(volId string) error {
	tgtNqn := nvmeofTgtPrefix + volId
	return t.RemoveNvmeofTarget(volId, tgtNqn)
}
