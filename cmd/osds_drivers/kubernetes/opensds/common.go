// Copyright (c) 2016 Huawei Technologies Co., Ltd. All Rights Reserved.
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

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/opensds/opensds/cmd/utils"
)

type Result struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Device  string `json:"device,omitempty"`
}

type DefaultOptions struct {
	Action_type string `json:"action_type"`
	MountPath   string `json:"mountPath`
	FsType      string `json:"kubernetes.io/fsType"`
}

// VolumeRequest is a structure for all properties of
// a volume request
type VolumeRequest struct {
	DockId       string `json:"dockId,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	Id           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	VolumeType   string `json:"volumeType,omitempty"`
	Size         int    `json:"size"`
	AllowDetails bool   `json:"allowDetails"`

	Device   string `json:"device"`
	MountDir string `json:"mountDir"`
	FsType   string `json:"fsType"`
}

// VolumeResponse is a structure for all properties of
// an volume for a non detailed query
type VolumeResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type FlexVolumePlugin interface {
	NewOptions() interface{}
	Init() Result
	Attach(opt interface{}) Result
	Detach(device string) Result
	Mount(mountDir string, device string, opt interface{}) Result
	Unmount(mountDir string) Result
}

func Succeed(a ...interface{}) Result {
	return Result{
		Status:  "Success",
		Message: fmt.Sprint(a...),
	}
}

func Fail(a ...interface{}) Result {
	return Result{
		Status:  "Failure",
		Message: fmt.Sprint(a...),
	}
}

func finish(result Result) {
	code := 1
	if result.Status == "Success" {
		code = 0
	}
	res, err := json.Marshal(result)
	if err != nil {
		fmt.Println("{\"status\":\"Failure\",\"message\":\"JSON error\"}")
	} else {
		fmt.Println(string(res))
	}
	os.Exit(code)
}

func RunPlugin(plugin FlexVolumePlugin) {
	if len(os.Args) < 2 {
		finish(Fail("Expected at least one argument"))
	}

	switch os.Args[1] {
	case "init":
		finish(plugin.Init())

	case "attach":
		if len(os.Args) != 3 {
			finish(Fail("attach expected exactly 3 arguments; got ", os.Args))
		}

		opt := plugin.NewOptions()
		if err := json.Unmarshal([]byte(os.Args[2]), opt); err != nil {
			finish(Fail("Could not parse options for attach:", err))
		}

		finish(plugin.Attach(opt))

	case "detach":
		if len(os.Args) != 3 {
			finish(Fail("detach expected exactly 3 arguments; got ", os.Args))
		}

		device := os.Args[2]
		finish(plugin.Detach(device))

	case "mount":
		if len(os.Args) != 5 {
			finish(Fail("mount expected exactly 5 argument; got ", os.Args))
		}

		mountDir := os.Args[2]
		device := os.Args[3]

		opt := plugin.NewOptions()
		if err := json.Unmarshal([]byte(os.Args[4]), opt); err != nil {
			finish(Fail("Could not parse options for mount; got ", os.Args[4]))
		}

		finish(plugin.Mount(mountDir, device, opt))

	case "unmount":
		if len(os.Args) != 3 {
			finish(Fail("mount expected exactly 5 argument; got ", os.Args))
		}

		mountDir := os.Args[2]

		finish(plugin.Unmount(mountDir))

	default:
		finish(Fail("Not sure what to do. Called with: ", os.Args))
	}
}

// CheckHTTPResponseStatusCode compares http response header StatusCode against expected
// statuses. Primary function is to ensure StatusCode is in the 20x (return nil).
// Ok: 200. Created: 201. Accepted: 202. No Content: 204. Partial Content: 206.
// Otherwise return error message.
func CheckHTTPResponseStatusCode(resp *http.Response) error {
	switch resp.StatusCode {
	case 200, 201, 202, 204, 206:
		return nil
	case 400:
		return errors.New("Error: response == 400 bad request")
	case 401:
		return errors.New("Error: response == 401 unauthorised")
	case 403:
		return errors.New("Error: response == 403 forbidden")
	case 404:
		return errors.New("Error: response == 404 not found")
	case 405:
		return errors.New("Error: response == 405 method not allowed")
	case 409:
		return errors.New("Error: response == 409 conflict")
	case 413:
		return errors.New("Error: response == 413 over limit")
	case 415:
		return errors.New("Error: response == 415 bad media type")
	case 422:
		return errors.New("Error: response == 422 unprocessable")
	case 429:
		return errors.New("Error: response == 429 too many request")
	case 500:
		return errors.New("Error: response == 500 instance fault / server err")
	case 501:
		return errors.New("Error: response == 501 not implemented")
	case 503:
		return errors.New("Error: response == 503 service unavailable")
	}
	return errors.New("Error: unexpected response status code")
}

type DockNode struct {
	Id      string   `json:"id"`
	Drivers []string `json:"drivers"`
	Address string   `json:"address"`
}

func getDockNode() (*DockNode, error) {
	var nodePtr = &DockNode{}

	userJSON, err := ioutil.ReadFile("/etc/opensds/dock_node.json")
	if err != nil {
		log.Println("ReadFile json failed:", err)
		return nodePtr, err
	}

	if err = json.Unmarshal(userJSON, nodePtr); err != nil {
		log.Println("Unmarshal json failed:", err)
		return nodePtr, err
	}
	return nodePtr, nil
}

func GetDockId() (string, error) {
	dock, err := getDockNode()
	if err != nil {
		log.Println("Could not get dock routes:", err)
		return "", err
	}

	host, err := utils.GetHostIP()
	if err != nil {
		log.Println("Could not get dock host ip:", err)
		return "", err
	}

	if dock.Address == host {
		return dock.Id, nil
	} else {
		err = errors.New("Could not find dock service!")
		return "", err
	}
}
