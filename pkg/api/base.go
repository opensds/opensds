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

package api

import (
	"fmt"
	"net/url"

	"github.com/astaxie/beego"
	log "github.com/golang/glog"
	"github.com/opensds/opensds/pkg/model"
)

type BasePortal struct {
	beego.Controller
}

func (this *BasePortal) GetParameters() (map[string][]string, error) {

	u, err := url.Parse(this.Ctx.Request.URL.String())
	if err != nil {
		return nil, err
	}
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (this *BasePortal) ErrorHandle(errMsg string, errType int, err error) {
	reason := fmt.Sprintf(errMsg+": %s", err.Error())
	this.Ctx.Output.SetStatus(model.ErrorBadRequest)
	this.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
	log.Error(reason)
}

func (this *BasePortal) SuccessHandle(status int, body []byte) {
	this.Ctx.Output.SetStatus(status)
	if body != nil {
		this.Ctx.Output.Body(body)
	}
}
