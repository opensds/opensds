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

package lvm

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/opensds/opensds/contrib/drivers/utils/config"
	pb "github.com/opensds/opensds/pkg/dock/proto"
	"github.com/opensds/opensds/pkg/model"
	"github.com/opensds/opensds/pkg/utils/config"
)

func TestSetup(t *testing.T) {
	var d = &Driver{}
	config.CONF.OsdsDock.Backends.LVM.ConfigPath = "testdata/lvm.yaml"
	var expectedDriver = &Driver{
		conf: &LVMConfig{
			Pool: map[string]PoolProperties{
				"vg001": {
					DiskType: "SSD",
					AZ:       "default",
				},
			},
			TgtBindIp: "192.168.56.105",
		},
		handler: execCmd,
	}

	if err := d.Setup(); err != nil {
		t.Errorf("Setup lvm driver failed: %+v\n", err)
	}
	f1, f2 := reflect.ValueOf(d.handler), reflect.ValueOf(expectedDriver.handler)
	if f1.Pointer() != f2.Pointer() {
		t.Errorf("The type of two methods are not the same!\n")
	}
	if !reflect.DeepEqual(d.conf, expectedDriver.conf) {
		t.Errorf("Expected %+v, got %+v", expectedDriver.conf, d.conf)
	}
}

var fd = &Driver{
	conf: &LVMConfig{
		Pool: map[string]PoolProperties{
			"vg001": {
				DiskType: "SSD",
				AZ:       "lvm",
			},
		},
	},
	handler: fakeHandler,
}

func fakeHandler(script string, cmd []string) (string, error) {
	switch script {
	case "lvcreate":
		return "", nil
	case "lvdisplay":
		return string(sampleLV), nil
	case "lvremove":
		return "", nil
	case "vgdisplay":
		return string(sampleVG), nil
	default:
		break
	}

	return "", fmt.Errorf("Script %s not supported!", script)
}

func TestCreateVolume(t *testing.T) {
	opt := &pb.CreateVolumeOpts{
		Name:        "test001",
		Description: "volume for testing",
		Size:        int64(1),
		PoolName:    "vg001",
	}
	var expected = &model.VolumeSpec{
		BaseModel:   &model.BaseModel{},
		Name:        "test001",
		Description: "volume for testing",
		Size:        int64(1),
		Status:      "available",
		Metadata: map[string]string{
			"lvPath": "/dev/vg001/test001",
		},
	}
	vol, err := fd.CreateVolume(opt)
	if err != nil {
		t.Error("Failed to create volume:", err)
	}
	vol.Id = ""
	if !reflect.DeepEqual(vol, expected) {
		t.Errorf("Expected %+v, got %+v\n", expected, vol)
	}
}

func TestPullVolume(t *testing.T) {
	volIdentifier := "/dev/vg001/test001"
	var expected = &model.VolumeSpec{
		Status: "available",
	}
	vol, err := fd.PullVolume(volIdentifier)
	if err != nil {
		t.Error("Failed to pull volume:", err)
	}
	if !reflect.DeepEqual(vol, expected) {
		t.Errorf("Expected %+v, got %+v\n", expected, vol)
	}
}

func TestDeleteVolume(t *testing.T) {
	opt := &pb.DeleteVolumeOpts{
		Metadata: map[string]string{
			"lvPath": "/dev/vg001/test001",
		},
	}
	if err := fd.DeleteVolume(opt); err != nil {
		t.Error("Failed to delete volume:", err)
	}
}

func TestCreateSnapshot(t *testing.T) {
	opt := &pb.CreateVolumeSnapshotOpts{
		Name:        "snap001",
		Description: "volume snapshot for testing",
		Size:        int64(1),
		VolumeId:    "bd5b12a8-a101-11e7-941e-d77981b584d8",
		Metadata: map[string]string{
			"lvPath": "/dev/vg001/test001",
		},
	}
	var expected = &model.VolumeSnapshotSpec{
		BaseModel:   &model.BaseModel{},
		Name:        "snap001",
		Description: "volume snapshot for testing",
		Size:        int64(1),
		Status:      "available",
		VolumeId:    "bd5b12a8-a101-11e7-941e-d77981b584d8",
		Metadata: map[string]string{
			"lvsPath": "/dev/vg001/snap001",
		},
	}
	snp, err := fd.CreateSnapshot(opt)
	if err != nil {
		t.Error("Failed to create volume snapshot:", err)
	}
	snp.Id = ""
	if !reflect.DeepEqual(snp, expected) {
		t.Errorf("Expected %+v, got %+v\n", expected, snp)
	}
}

func TestPullSnapshot(t *testing.T) {
	snpIdentifier := "/dev/vg001/snp001"
	var expected = &model.VolumeSnapshotSpec{
		Status: "available",
	}
	snp, err := fd.PullSnapshot(snpIdentifier)
	if err != nil {
		t.Error("Failed to pull volume snapshot:", err)
	}
	if !reflect.DeepEqual(snp, expected) {
		t.Errorf("Expected %+v, got %+v\n", expected, snp)
	}
}

func TestDeleteSnapshot(t *testing.T) {
	opt := &pb.DeleteVolumeSnapshotOpts{
		Metadata: map[string]string{
			"lvsPath": "/dev/vg001/snap001",
		},
	}
	if err := fd.DeleteSnapshot(opt); err != nil {
		t.Error("Failed to delete volume snapshot:", err)
	}
}

func TestListPools(t *testing.T) {
	var expected = []*model.StoragePoolSpec{
		{
			BaseModel:     &model.BaseModel{},
			Name:          "vg001",
			TotalCapacity: int64(18),
			FreeCapacity:  int64(18),
			Extras: model.ExtraSpec{
				"diskType": "SSD",
			},
			AvailabilityZone: "lvm",
		},
	}
	pols, err := fd.ListPools()
	if err != nil {
		t.Error("Failed to list pools:", err)
	}
	for i := range pols {
		pols[i].Id = ""
	}
	if !reflect.DeepEqual(pols, expected) {
		t.Errorf("Expected %+v, got %+v\n", expected, pols)
	}
}

var (
	sampleLV = `
		--- Logical volume ---
		LV Path                /dev/vg001/test001
		LV Name                test001
		VG Name                vg001
		LV UUID                mFdrHm-uiQS-TRK2-Iwua-jdQr-7sYd-ReayKW
		LV Write Access        read/write
		LV Creation host, time krej-Lenovo-IdeaPad-Y470, 2017-11-20 16:43:20 +0800
		LV Status              available
		# open                 0
		LV Size                1.00 GiB
		Current LE             256
		Segments               1
		Allocation             inherit
		Read ahead sectors     auto
		- currently set to     256
		Block device           253:0
	`
	sampleVG = `
		--- Volume group ---
		VG Name               vg001
		System ID
		Format                lvm2
		Metadata Areas        1
		Metadata Sequence No  3
		VG Access             read/write
		VG Status             resizable
		MAX LV                0
		Cur LV                0
		Open LV               0
		Max PV                0
		Cur PV                1
		Act PV                1
		VG Size               18.62 GiB
		PE Size               4.00 MiB
		Total PE              4768
		Alloc PE / Size       0 / 0
		Free  PE / Size       4768 / 18.62 GiB
		VG UUID               Yn9utl-eqjH-1sJG-0fdb-dGTX-PLJI-FjMO0v

		--- Volume group ---
		VG Name               ubuntu-vg
		System ID
		Format                lvm2
		Metadata Areas        1
		Metadata Sequence No  3
		VG Access             read/write
		VG Status             resizable
		MAX LV                0
		Cur LV                2
		Open LV               2
		Max PV                0
		Cur PV                1
		Act PV                1
		VG Size               127.52 GiB
		PE Size               4.00 MiB
		Total PE              32645
		Alloc PE / Size       32638 / 127.49 GiB
		Free  PE / Size       7 / 28.00 MiB
		VG UUID               fQbqtg-3vDQ-vk3U-gfsT-50kJ-30pq-OZVSJH
	`
	sampleLVS = `
		--- Logical volume ---
		LV Path                /dev/vg001/snap001
		LV Name                snap001
		VG Name                vg001
		LV UUID                We6GmQ-H675-mfQv-iQkO-rVUI-LuBx-YBIBwr
		LV Write Access        read only
		LV Creation host, time krej-Lenovo-IdeaPad-Y470, 2017-11-20 17:05:02 +0800
		LV snapshot status     active destination for test001
		LV Status              available
		# open                 0
		LV Size                1.00 GiB
		Current LE             256
		COW-table size         1.00 GiB
		COW-table LE           256
		Allocated to snapshot  0.00%
		Snapshot chunk size    4.00 KiB
		Segments               1
		Allocation             inherit
		Read ahead sectors     auto
		- currently set to     256
		Block device           253:3
	`
)
