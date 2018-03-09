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

This package includes a collection of fake stuffs for testing work.
*/

package collection

import (
	"github.com/opensds/opensds/pkg/model"
)

var (
	SampleProfiles = []model.ProfileSpec{
		{
			BaseModel: &model.BaseModel{
				Id: "1106b972-66ef-11e7-b172-db03f3689c9c",
			},
			Name:        "default",
			Description: "default policy",
			Extras:      model.ExtraSpec{},
		},
		{
			BaseModel: &model.BaseModel{
				Id: "2f9c0a04-66ef-11e7-ade2-43158893e017",
			},
			Name:        "silver",
			Description: "silver policy",
			Extras: model.ExtraSpec{
				"diskType": "SAS",
				"thin":     true,
			},
		},
	}

	SampleDocks = []model.DockSpec{
		{
			BaseModel: &model.BaseModel{
				Id: "b7602e18-771e-11e7-8f38-dbd6d291f4e4",
			},
			Name:        "sample1",
			Description: "sample backend service",
			Endpoint:    "localhost:50050",
			DriverName:  "docktest",
		},
		{
			BaseModel: &model.BaseModel{
				Id: "b7602e18-771e-11e7-8f38-dbd6d291f4e3",
			},
			Name:        "sample1",
			Description: "sample backend service",
			Endpoint:    "localhost:50050",
			DriverName:  "docktest",
		},
	}
	SampleDocks_discover = []model.DockSpec{
		{
			BaseModel: &model.BaseModel{
				Id: "b7602e18-771e-11e7-8f38-dbd6d291f4e0",
			},
			Name:        "sample",
			Description: "sample backend service",
			Endpoint:    "localhost:50050",
			DriverName:  "sample",
		},
	}
	SamplePools = []model.StoragePoolSpec{
		{
			BaseModel: &model.BaseModel{
				Id: "084bf71e-a102-11e7-88a8-e31fe6d52248",
			},
			Name:             "sample-pool-01",
			Description:      "This is the first sample storage pool for testing",
			TotalCapacity:    int64(100),
			FreeCapacity:     int64(90),
			DockId:           "b7602e18-771e-11e7-8f38-dbd6d291f4e1",
			AvailabilityZone: "default",
			Extras: model.ExtraSpec{
				"diskType": "SSD",
				"thin":     true,
			},
		},
		{
			BaseModel: &model.BaseModel{
				Id: "084bf71e-a102-11e7-88a8-e31fe6d52248",
			},
			Name:             "sample-pool-01",
			Description:      "This is the first sample storage pool for testing",
			TotalCapacity:    int64(100),
			FreeCapacity:     int64(90),
			DockId:           "b7602e18-771e-11e7-8f38-dbd6d291f4e0",
			AvailabilityZone: "default",
			Extras: model.ExtraSpec{
				"diskType": "SSD",
				"thin":     true,
			},
		},
	}

	SamplePools_discovery = []model.StoragePoolSpec{
		{

			BaseModel: &model.BaseModel{
				Id: "084bf71e-a102-11e7-88a8-e31fe6d52248",
			},
			Name:             "sample-pool-01",
			Description:      "This is the first sample storage pool for testing",
			TotalCapacity:    int64(100),
			FreeCapacity:     int64(90),
			DockId:           "b7602e18-771e-11e7-8f38-dbd6d291f4e0",
			AvailabilityZone: "default",
			Extras: model.ExtraSpec{
				"diskType": "SSD",
				"thin":     true,
			},
		},
		{
			BaseModel: &model.BaseModel{
				Id: "a594b8ac-a103-11e7-985f-d723bcf01b5f",
			},
			Name:             "sample-pool-02",
			Description:      "This is the second sample storage pool for testing",
			TotalCapacity:    int64(200),
			FreeCapacity:     int64(170),
			AvailabilityZone: "default",
			DockId:           "b7602e18-771e-11e7-8f38-dbd6d291f4e0",
			Extras: model.ExtraSpec{
				"diskType": "SAS",
				"thin":     true,
			},
		},
	}

	SampleVolumes = []model.VolumeSpec{
		{
			BaseModel: &model.BaseModel{
				Id: "bd5b12a8-a101-11e7-941e-d77981b584d8",
			},
			Name:        "sample-volume",
			Description: "This is a sample volume for testing",
			Size:        int64(1),
			Status:      "available",
			PoolId:      "084bf71e-a102-11e7-88a8-e31fe6d52248",
			ProfileId:   "1106b972-66ef-11e7-b172-db03f3689c9c",
		},
	}

	SampleConnection = model.ConnectionInfo{
		DriverVolumeType: "iscsi",
		ConnectionData: map[string]interface{}{
			"targetDiscovered": true,
			"targetIqn":        "iqn.2017-10.io.opensds:volume:00000001",
			"targetPortal":     "127.0.0.0.1:3260",
			"discard":          false,
		},
	}

	SampleAttachments = []model.VolumeAttachmentSpec{
		{
			BaseModel: &model.BaseModel{
				Id: "f2dda3d2-bf79-11e7-8665-f750b088f63e",
			},
			Status:   "available",
			VolumeId: "bd5b12a8-a101-11e7-941e-d77981b584d8",
			HostInfo: model.HostInfo{},
			ConnectionInfo: model.ConnectionInfo{
				DriverVolumeType: "iscsi",
				ConnectionData: map[string]interface{}{
					"targetDiscovered": true,
					"targetIqn":        "iqn.2017-10.io.opensds:volume:00000001",
					"targetPortal":     "127.0.0.0.1:3260",
					"discard":          false,
				},
			},
		},
	}

	SampleSnapshots = []model.VolumeSnapshotSpec{
		{
			BaseModel: &model.BaseModel{
				Id: "3769855c-a102-11e7-b772-17b880d2f537",
			},
			Name:        "sample-snapshot-01",
			Description: "This is the first sample snapshot for testing",
			Size:        int64(1),
			Status:      "created",
			VolumeId:    "bd5b12a8-a101-11e7-941e-d77981b584d8",
		},
		{
			BaseModel: &model.BaseModel{
				Id: "3bfaf2cc-a102-11e7-8ecb-63aea739d755",
			},
			Name:        "sample-snapshot-02",
			Description: "This is the second sample snapshot for testing",
			Size:        int64(1),
			Status:      "created",
			VolumeId:    "bd5b12a8-a101-11e7-941e-d77981b584d8",
		},
	}
)

// The Byte*** variable here is designed for unit test in client package.
// For how to ultilize these pre-assigned variables, please refer to
// (github.com/opensds/opensds/client/dock_test.go).
var (
	ByteProfile = `{
		"id": "1106b972-66ef-11e7-b172-db03f3689c9c",
		"name": "default",
		"description": "default policy"
	}`

	ByteProfiles = `[
		{
			"id": "1106b972-66ef-11e7-b172-db03f3689c9c",
			"name": "default",
			"description": "default policy"
		},
		{
			"id": "2f9c0a04-66ef-11e7-ade2-43158893e017",
			"name": "silver",
			"description": "silver policy",
			"extras": {
				"diskType":"SAS"
			}
		}
	]`

	ByteExtras = `{
		"diskType":"SAS"
	}`

	ByteDock = `{
		"id": "b7602e18-771e-11e7-8f38-dbd6d291f4e0",
		"name":        "sample",
		"description": "sample backend service",
		"endpoint":    "localhost:50050",
		"driverName":  "sample"
	}`

	ByteDocks = `[
		{
			"id": "b7602e18-771e-11e7-8f38-dbd6d291f4e0",
			"name":        "sample",
			"description": "sample backend service",
			"endpoint":    "localhost:50050",
			"driverName":  "sample"
		}
	]`

	BytePool = `{
		"id": "084bf71e-a102-11e7-88a8-e31fe6d52248",
		"name": "sample-pool-01",
		"description": "This is the first sample storage pool for testing",
		"totalCapacity": 100,
		"freeCapacity": 90,
		"dockId": "b7602e18-771e-11e7-8f38-dbd6d291f4e0",
		"extras": {
			"diskType": "SSD"
		}
	}`

	BytePools = `[
		{
			"id": "084bf71e-a102-11e7-88a8-e31fe6d52248",
			"name": "sample-pool-01",
			"description": "This is the first sample storage pool for testing",
			"totalCapacity": 100,
			"freeCapacity": 90,
			"dockId": "b7602e18-771e-11e7-8f38-dbd6d291f4e0",
			"extras": {
				"diskType": "SSD"
			}
		},
		{
			"id": "a594b8ac-a103-11e7-985f-d723bcf01b5f",
			"name": "sample-pool-02",
			"description": "This is the second sample storage pool for testing",
			"totalCapacity": 200,
			"freeCapacity": 170,
			"dockId": "b7602e18-771e-11e7-8f38-dbd6d291f4e0",
			"extras": {
				"diskType": "SAS"
			}
		}
	]`

	ByteVolume = `{
		"id": "bd5b12a8-a101-11e7-941e-d77981b584d8",
		"name": "sample-volume",
		"description": "This is a sample volume for testing",
		"size": 1,
		"status": "available",
		"poolId": "084bf71e-a102-11e7-88a8-e31fe6d52248",
		"profileId": "1106b972-66ef-11e7-b172-db03f3689c9c"
	}`

	ByteVolumes = `[
		{
			"id": "bd5b12a8-a101-11e7-941e-d77981b584d8",
			"name": "sample-volume",
			"description": "This is a sample volume for testing",
			"size": 1,
			"status": "available",
			"poolId": "084bf71e-a102-11e7-88a8-e31fe6d52248",
			"profileId": "1106b972-66ef-11e7-b172-db03f3689c9c"
		}
	]`

	ByteAttachment = `{
		"id": "f2dda3d2-bf79-11e7-8665-f750b088f63e",
		"name": "sample-volume-attachment",
		"description": "This is a sample volume attachment for testing",
		"status": "available",
		"volumeId": "bd5b12a8-a101-11e7-941e-d77981b584d8",
		"hostInfo": {},
		"connectionInfo": {
			"driverVolumeType": "iscsi",
			"data": {
				"targetDiscovered": true,
				"targetIqn": "iqn.2017-10.io.opensds:volume:00000001",
				"targetPortal": "127.0.0.0.1:3260",
				"discard": false
			}
		}
	}`

	ByteAttachments = `[
		{
			"id": "f2dda3d2-bf79-11e7-8665-f750b088f63e",
			"name": "sample-volume-attachment",
			"description": "This is a sample volume attachment for testing",
			"status": "available",
			"volumeId": "bd5b12a8-a101-11e7-941e-d77981b584d8",
			"hostInfo": {},
			"connectionInfo": {
				"driverVolumeType": "iscsi",
				"data": {
					"targetDiscovered": true,
					"targetIqn": "iqn.2017-10.io.opensds:volume:00000001",
					"targetPortal": "127.0.0.0.1:3260",
					"discard": false
				}
			}
		}
	]`

	ByteSnapshot = `{
		"id": "3769855c-a102-11e7-b772-17b880d2f537",
		"name": "sample-snapshot-01",
		"description": "This is the first sample snapshot for testing",
		"size": 1,
		"status": "created",
		"volumeId": "bd5b12a8-a101-11e7-941e-d77981b584d8"		
	}`

	ByteSnapshots = `[
		{
			"id": "3769855c-a102-11e7-b772-17b880d2f537",
			"name": "sample-snapshot-01",
			"description": "This is the first sample snapshot for testing",
			"size": 1,
			"status": "created",
			"volumeId": "bd5b12a8-a101-11e7-941e-d77981b584d8"	
		},
		{
			"id": "3bfaf2cc-a102-11e7-8ecb-63aea739d755",
			"name": "sample-snapshot-02",
			"description": "This is the second sample snapshot for testing",
			"size": 1,
			"status": "created",
			"volumeId": "bd5b12a8-a101-11e7-941e-d77981b584d8"	
		}
	]`

	ByteVersion = `{
		"name": "v1beta",
		"status": "SUPPORTED",
		"updatedAt": "2017-04-10T14:36:58.014Z"
	}`

	ByteVersions = `[
		{
			"name": "v1beta",
			"status": "CURRENT",
			"updatedAt": "2017-07-10T14:36:58.014Z"
		}
	]`
)

// The StringSlice*** variable here is designed for unit test in etcd package.
// For how to ultilize these pre-assigned variables, please refer to
// (github.com/opensds/opensds/pkg/db/drivers/etcd/etcd_test.go).
var (
	StringSliceProfiles = []string{
		`{
			"id": "1106b972-66ef-11e7-b172-db03f3689c9c",
			"name":        "default",
			"description": "default policy",
			"extras": {}
		}`,
		`{
			"id": "2f9c0a04-66ef-11e7-ade2-43158893e017",
			"name":        "silver",
			"description": "silver policy",
			"extras": {
				"diskType": "SAS",
				"thin":     true
			}
		}`,
	}

	StringSliceDocks = []string{
		`{
			"id": "b7602e18-771e-11e7-8f38-dbd6d291f4e4",
			"name":        "sample1",
			"description": "sample backend service",
			"endpoint":    "localhost:50050",
			"driverName":  "docktest"
		}`,
		`{
			"id": "b7602e18-771e-11e7-8f38-dbd6d291f4e0",
			"name":        "sample1",
			"description": "sample backend service",
			"endpoint":    "localhost:50050",
			"driverName":  "sample"
		}`,
		`{
			"id": "b7602e18-771e-11e7-8f38-dbd6d291f4e0",
			"name":        "sample2",
			"description": "sample backend service",
			"endpoint":    "localhost:50050",
			"driverName":  "sample"
		}`,
		`{
			"id": "b7602e18-771e-11e7-8f38-dbd6d291f4e0",
			"name":        "sample3",
			"description": "sample backend service",
			"endpoint":    "localhost:50050",
			"driverName":  "sample"
		}`,
		`{
			"id": "b7602e18-771e-11e7-8f38-dbd6d291f4e2",
			"name":        "sample1",
			"description": "sample backend service",
			"endpoint":    "localhost:50050",
			"driverName":  "sample1"
		}`,
		`{
			"id": "b7602e18-771e-11e7-8f38-dbd6d291f4e3",
			"name":        "sample1",
			"description": "sample backend service",
			"endpoint":    "localhost:50050",
			"driverName":  "docktest"
		}`,

		`{
			"id": "b7602e18-771e-11e7-8f38-dbd6d291f4e5",
			"name":        "sample1",
			"description": "sample backend service",
			"endpoint":    "localhost:50050",
			"driverName":  "docktest"
		}`,
		`{
			"id": "b7602e18-771e-11e7-8f38-dbd6d291f4e6",
			"name":        "sample1",
			"description": "sample backend service",
			"endpoint":    "localhost:50050",
			"driverName":  "docktest"
		}`,
	}

	StringSlicePools = []string{
		`{
			"id": "084bf71e-a102-11e7-88a8-e31fe6d52248",
			"name":             "sample-pool-01",
			"description":      "This is the first sample storage pool for testing",
			"totalCapacity":    100,
			"freeCapacity":     90,
			"dockId":           "b7602e18-771e-11e7-8f38-dbd6d291f4e1",
			"availabilityZone": "default",
			"extras": {
				"diskType": "SSD",
				"thin":     true
			}
		}`,

		`{
			"id": "084bf71e-a102-11e7-88a8-e31fe6d52248",
			"name":             "sample-pool-01",
			"description":      "This is the first sample storage pool for testing",
			"totalCapacity":    100,
			"freeCapacity":     90,
			"dockId":           "b7602e18-771e-11e7-8f38-dbd6d291f4e0",
			"availabilityZone": "default",
			"extras": {
				"diskType": "SSD",
				"thin":     true
			}
		}`,
		`{
			"id": "a594b8ac-a103-11e7-985f-d723bcf01b5f",
			"name":             "sample-pool-02",
			"description":      "This is the second sample storage pool for testing",
			"totalCapacity":    200,
			"freeCapacity":     170,
			"availabilityZone": "default",
			"dockId":           "b7602e18-771e-11e7-8f38-dbd6d291f4e0",
			"extras": {
				"diskType": "SAS",
				"thin":     true
			}
		}`,
		`{
			"id": "a594b8ac-a103-11e7-985f-d723bcf01b5f",
			"name":             "sample-pool-03",
			"description":      "This is the second sample storage pool for testing",
			"totalCapacity":    200,
			"freeCapacity":     170,
			"availabilityZone": "default",
			"dockId":           "b7602e18-771e-11e7-8f38-dbd6d291f4e0",
			"extras": {
				"diskType": "SAS",
				"thin":     true
			}
		}`,
	}

	StringSliceVolumes = []string{
		`{
			"id": "bd5b12a8-a101-11e7-941e-d77981b584d8",
			"name":        "sample-volume",
			"description": "This is a sample volume for testing",
			"size":        1,
			"status":      "available",
			"poolId":      "084bf71e-a102-11e7-88a8-e31fe6d52248",
			"profileId":   "1106b972-66ef-11e7-b172-db03f3689c9c"
		}`,
	}

	StringSliceAttachments = []string{
		`{
			"id": "f2dda3d2-bf79-11e7-8665-f750b088f63e",
			"status":   "available",
			"volumeId": "bd5b12a8-a101-11e7-941e-d77981b584d8",
			"hostInfo": {},
			"connectionInfo": {
				"driverVolumeType": "iscsi",
				"data": {
					"targetDiscovered": true,
					"targetIqn":        "iqn.2017-10.io.opensds:volume:00000001",
					"targetPortal":     "127.0.0.0.1:3260",
					"discard":          false
				}
			}
		}`,
	}

	StringSliceSnapshots = []string{
		`{
			"id": "3769855c-a102-11e7-b772-17b880d2f537",
			"name":        "sample-snapshot-01",
			"description": "This is the first sample snapshot for testing",
			"size":        1,
			"status":      "created",
			"volumeId":    "bd5b12a8-a101-11e7-941e-d77981b584d8"
		}`,
		`{
			"id": "3bfaf2cc-a102-11e7-8ecb-63aea739d755",
			"name":        "sample-snapshot-02",
			"description": "This is the second sample snapshot for testing",
			"size":        1,
			"status":      "created",
			"volumeId":    "bd5b12a8-a101-11e7-941e-d77981b584d8"
		}`,
	}
)
