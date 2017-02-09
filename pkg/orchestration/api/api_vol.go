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

/*
This module implements the enry into the operations of orchestration module.

Request about volume operation will be passed to the grpc client and requests
about other resources (database, fileSystem, etc) will be passed to metaData
service module.

*/

package api

import (
	"log"

	"github.com/opensds/opensds/pkg/orchestration/grpcapi"
)

func CreateVolume(resourceType string, name string, size int) (string, error) {
	result, err := grpcapi.CreateVolume(resourceType, name, size)

	if err != nil {
		log.Println("Error occured in orchestration module when create volume!")
		return "", err
	} else {
		return result, nil
	}
}

func GetVolume(resourceType string, volID string) (string, error) {
	result, err := grpcapi.GetVolume(resourceType, volID)

	if err != nil {
		log.Println("Error occured in orchestration module when get volume!")
		return "", err
	} else {
		return result, nil
	}
}

func GetAllVolumes(resourceType string, allowDetails bool) (string, error) {
	result, err := grpcapi.GetAllVolumes(resourceType, allowDetails)

	if err != nil {
		log.Println("Error occured in orchestration module when get all volumes!")
		return "", err
	} else {
		return result, nil
	}
}

func UpdateVolume(resourceType string, volID string, name string) (string, error) {
	result, err := grpcapi.UpdateVolume(resourceType, volID, name)

	if err != nil {
		log.Println("Error occured in orchestration module when update volume!")
		return "", err
	} else {
		return result, nil
	}
}

func DeleteVolume(resourceType string, volID string) (string, error) {
	result, err := grpcapi.DeleteVolume(resourceType, volID)

	if err != nil {
		log.Println("Error occured in orchestration module when delete volume!")
		return "", err
	} else {
		return result, nil
	}
}

func MountVolume(resourceType, volID, host, mountpoint string) (string, error) {
	result, err := grpcapi.MountVolume(resourceType, volID, host, mountpoint)

	if err != nil {
		log.Println("Error occured in orchestration module when mount volume!")
		return "", err
	} else {
		return result, nil
	}
}

func UnmountVolume(resourceType, volID, attachment string) (string, error) {
	result, err := grpcapi.UnmountVolume(resourceType, volID, attachment)

	if err != nil {
		log.Println("Error occured in orchestration module when unmount volume!")
		return "", err
	} else {
		return result, nil
	}
}
