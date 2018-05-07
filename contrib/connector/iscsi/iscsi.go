// Copyright (c) 2018 Huawei Technologies Co., Ltd. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package iscsi

import (
	"strconv"

	"github.com/opensds/opensds/contrib/connector"
)

var (
	ISCSI_DRIVER = "iscsi"
)

type Iscsi struct{}

func init() {
	connector.RegisterConnector(ISCSI_DRIVER, &Iscsi{})
}

func (isc *Iscsi) Attach(conn map[string]interface{}) (string, error) {
	iscsiCon := ParseIscsiConnectInfo(conn)

	return Connect(iscsiCon.TgtPortal, iscsiCon.TgtIQN, strconv.Itoa(iscsiCon.TgtLun))
}

func (isc *Iscsi) Detach(conn map[string]interface{}) error {
	iscsiCon := ParseIscsiConnectInfo(conn)

	return Disconnect(iscsiCon.TgtPortal, iscsiCon.TgtIQN)
}
