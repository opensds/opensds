// Copyright (c) 2017 Huawei Technologies Co., Ltd. All Rights Reserved.
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

/*
This module defines an standard table of storage driver. The default storage
driver is sample driver used for testing. If you want to use other storage
plugin, just modify Init() and Clean() method.
*/

package filesharedrivers

import (
	_ "github.com/opensds/opensds/contrib/backup/multicloud"
	_ "github.com/opensds/opensds/contrib/drivers/filesharedrivers/nfsnative"
	"github.com/opensds/opensds/pkg/model"
	pb "github.com/opensds/opensds/pkg/model/fileshareproto"
)

type FileShareDriver interface {
	//Any initialization the fileshare driver does while starting.
	Setup() error
	//Any operation the fileshare driver does while stopping.
	Unset() error

	CreateFileShare(opt *pb.CreateFileShareOpts) (*model.FileShareSpec, error)
	//ListPools() (interface{}, interface{})
	ListPools() ([]*model.StoragePoolSpec, error)
	DeleteFileShare(opts *pb.DeleteFileShareOpts) (*model.FileShareSpec, error)
}


// Init
func Init(resourceType string) FileShareDriver {
	var f FileShareDriver
	switch resourceType {
	//case config.NFSNativeDriverType:
	//	f = &nfsnative.Driver{}
	//	break
	default:
	//	f = &sample.Driver{}
		break
	}
	return f
}



// Clean
func Clean(f FileShareDriver) FileShareDriver {
	// Execute different clean operations according to the VolumeDriver type.
	switch f.(type) {
	//case *nfsnative.Driver:
	//	break
	//case *sample.Driver:
	//	break
	default:
		break
	}
	_ = f.Unset()
	f = nil

	return f
}
