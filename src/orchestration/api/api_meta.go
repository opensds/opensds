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
	//"log"

	"metaData"
)

func CreateDatabase(name string, size int) (string, error) {
	result, err := metaData.CreateDatabase(name, size)

	if err != nil {
		return "Error", err
	} else {
		return result, nil
	}
}

func GetDatabase(id int, name string) (string, error) {
	result, err := metaData.GetDatabase(id, name)

	if err != nil {
		return result, err
	} else {
		return result, nil
	}
}

func GetAllDatabases() (string, error) {
	result, err := metaData.GetAllDatabases()

	if err != nil {
		return result, err
	} else {
		return result, nil
	}
}

func UpdateDatabase(id int, size int, name string) (string, error) {
	result, err := metaData.UpdateDatabase(id, size, name)

	if err != nil {
		return "Error", err
	} else {
		return result, nil
	}
}

func DeleteDatabase(id int, name string, cascade bool) (string, error) {
	result, err := metaData.DeleteDatabase(id, name, cascade)

	if err != nil {
		return "Error", err
	} else {
		return result, nil
	}
}

func CreateFileSystem(name string, size int) (string, error) {
	result, err := metaData.CreateFileSystem(name, size)

	if err != nil {
		return "Error", err
	} else {
		return result, nil
	}
}

func GetFileSystem(id int, name string) (string, error) {
	result, err := metaData.GetFileSystem(id, name)

	if err != nil {
		return result, err
	} else {
		return result, nil
	}
}

func GetAllFileSystems() (string, error) {
	result, err := metaData.GetAllFileSystems()

	if err != nil {
		return result, err
	} else {
		return result, nil
	}
}

func UpdateFileSystem(id int, size int, name string) (string, error) {
	result, err := metaData.UpdateFileSystem(id, size, name)

	if err != nil {
		return "Error", err
	} else {
		return result, nil
	}
}

func DeleteFileSystem(id int, name string, cascade bool) (string, error) {
	result, err := metaData.DeleteFileSystem(id, name, cascade)

	if err != nil {
		return "Error", err
	} else {
		return result, nil
	}
}
