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
This module implements a standard SouthBound interface of resources to
storage plugins.

*/

package dock

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	log "github.com/golang/glog"
	"github.com/opensds/opensds/contrib/connector"
	"github.com/opensds/opensds/contrib/drivers"
	c "github.com/opensds/opensds/pkg/context"
	"github.com/opensds/opensds/pkg/db"
	"github.com/opensds/opensds/pkg/dock/discovery"
	pb "github.com/opensds/opensds/pkg/dock/proto"
	"github.com/opensds/opensds/pkg/model"
	"github.com/opensds/opensds/pkg/utils/constants"

	_ "github.com/opensds/opensds/contrib/connector/iscsi"
	_ "github.com/opensds/opensds/contrib/connector/rbd"
)

// Brain is a global variable that controls the dock module.
var Brain *DockHub

// DockHub is a reference structure with fields that represent some required
// parameters for initializing and controlling the volume driver.
type DockHub struct {
	// Discoverer represents the mechanism of DockHub discovering the storage
	// capabilities from different backends.
	Discoverer discovery.DockDiscoverer
	// Driver represents the specified backend resource. This field is used
	// for initializing the specified volume driver.
	Driver drivers.VolumeDriver
}

// NewDockHub method creates a new DockHub and returns its pointer.
func NewDockHub(dockType string) *DockHub {
	return &DockHub{
		Discoverer: discovery.NewDockDiscoverer(dockType),
	}
}

// TriggerDiscovery
func (d *DockHub) TriggerDiscovery() error {
	var err error

	if err = d.Discoverer.Init(); err != nil {
		return err
	}

	ctx := &discovery.Context{
		StopChan: make(chan bool),
		ErrChan:  make(chan error),
		MetaChan: make(chan string),
	}
	go discovery.DiscoveryAndReport(d.Discoverer, ctx)
	go func(ctx *discovery.Context) {
		if err = <-ctx.ErrChan; err != nil {
			log.Error("When calling capabilty report method:", err)
			ctx.StopChan <- true
		}
	}(ctx)

	return err
}

// CreateVolume
func (d *DockHub) CreateVolume(opt *pb.CreateVolumeOpts) (*model.VolumeSpec, error) {
	//Get the storage drivers and do some initializations.
	d.Driver = drivers.Init(opt.GetDriverName())
	defer drivers.Clean(d.Driver)

	log.Info("Calling volume driver to create volume...")

	//Call function of StorageDrivers configured by storage drivers.
	vol, err := d.Driver.CreateVolume(opt)
	if err != nil {
		log.Error("When calling volume driver to create volume:", err)
		return nil, err
	}
	return vol, nil
}

// DeleteVolume
func (d *DockHub) DeleteVolume(opt *pb.DeleteVolumeOpts) error {
	var err error

	//Get the storage drivers and do some initializations.
	d.Driver = drivers.Init(opt.GetDriverName())
	defer drivers.Clean(d.Driver)

	log.Info("Calling volume driver to delete volume...")

	//Call function of StorageDrivers configured by storage drivers.
	if err = d.Driver.DeleteVolume(opt); err != nil {
		log.Error("When calling volume driver to delete volume:", err)
		return err
	}
	return nil
}

// ExtendVolume ...
func (d *DockHub) ExtendVolume(opt *pb.ExtendVolumeOpts) (*model.VolumeSpec, error) {
	//Get the storage drivers and do some initializations.
	d.Driver = drivers.Init(opt.GetDriverName())
	defer drivers.Clean(d.Driver)

	log.Info("Calling volume driver to extend volume...")

	//Call function of StorageDrivers configured by storage drivers.
	vol, err := d.Driver.ExtendVolume(opt)
	if err != nil {
		log.Error("When calling volume driver to extend volume:", err)
		return nil, err
	}
	return vol, nil
}

// CreateVolumeAttachment
func (d *DockHub) CreateVolumeAttachment(opt *pb.CreateAttachmentOpts) (*model.VolumeAttachmentSpec, error) {
	//Get the storage drivers and do some initializations.
	d.Driver = drivers.Init(opt.GetDriverName())
	defer drivers.Clean(d.Driver)

	log.Info("Calling volume driver to initialize volume connection...")

	//Call function of StorageDrivers configured by storage drivers.
	connInfo, err := d.Driver.InitializeConnection(opt)
	if err != nil {
		log.Error("Call driver to initialize volume connection failed:", err)
		return nil, err
	}

	var atc = &model.VolumeAttachmentSpec{
		BaseModel: &model.BaseModel{
			Id: opt.GetId(),
		},
		VolumeId: opt.GetVolumeId(),
		HostInfo: model.HostInfo{
			Platform:  opt.HostInfo.GetPlatform(),
			OsType:    opt.HostInfo.GetOsType(),
			Ip:        opt.HostInfo.GetIp(),
			Host:      opt.HostInfo.GetHost(),
			Initiator: opt.HostInfo.GetInitiator(),
		},
		ConnectionInfo: *connInfo,
		Metadata:       opt.GetMetadata(),
	}

	return atc, nil
}

// DeleteVolumeAttachment
func (d *DockHub) DeleteVolumeAttachment(opt *pb.DeleteAttachmentOpts) error {
	//Get the storage drivers and do some initializations.
	d.Driver = drivers.Init(opt.GetDriverName())
	defer drivers.Clean(d.Driver)

	log.Info("Calling volume driver to terminate volume connection...")

	//Call function of StorageDrivers configured by storage drivers.
	if err := d.Driver.TerminateConnection(opt); err != nil {
		log.Error("Call driver to terminate volume connection failed:", err)
		return err
	}
	return nil
}

// CreateSnapshot
func (d *DockHub) CreateSnapshot(opt *pb.CreateVolumeSnapshotOpts) (*model.VolumeSnapshotSpec, error) {
	//Get the storage drivers and do some initializations.
	d.Driver = drivers.Init(opt.GetDriverName())
	defer drivers.Clean(d.Driver)

	log.Info("Calling volume driver to create snapshot...")

	//Call function of StorageDrivers configured by storage drivers.
	snp, err := d.Driver.CreateSnapshot(opt)
	if err != nil {
		log.Error("Call driver to create volume snashot failed:", err)
		return nil, err
	}
	return snp, nil
}

// DeleteSnapshot
func (d *DockHub) DeleteSnapshot(opt *pb.DeleteVolumeSnapshotOpts) error {
	var err error

	//Get the storage drivers and do some initializations.
	d.Driver = drivers.Init(opt.GetDriverName())
	defer drivers.Clean(d.Driver)

	log.Info("Calling volume driver to delete snapshot...")

	//Call function of StorageDrivers configured by storage drivers.
	if err = d.Driver.DeleteSnapshot(opt); err != nil {
		log.Error("When calling volume driver to delete volume:", err)
		return err
	}
	return nil
}

// AttachVolume
func (d *DockHub) AttachVolume(opt *pb.AttachVolumeOpts) (string, error) {
	var connData = make(map[string]interface{})
	if err := json.Unmarshal([]byte(opt.GetConnectionData()), &connData); err != nil {
		return "", fmt.Errorf("Error occurred in dock module when unmarshalling connection data!")
	}

	con := connector.NewConnector(opt.GetAccessProtocol())
	if con == nil {
		return "", fmt.Errorf("Can not find connector (%s)!", opt.GetAccessProtocol())
	}

	return con.Attach(connData)
}

// DetachVolume
func (d *DockHub) DetachVolume(opt *pb.DetachVolumeOpts) error {
	var connData = make(map[string]interface{})
	if err := json.Unmarshal([]byte(opt.GetConnectionData()), &connData); err != nil {
		return fmt.Errorf("Error occurred in dock module when unmarshalling connection data!")
	}

	con := connector.NewConnector(opt.GetAccessProtocol())
	if con == nil {
		return fmt.Errorf("Can not find connector (%s)!", opt.GetAccessProtocol())
	}

	return con.Detach(connData)
}

func (d *DockHub) CreateVolumeGroup(opt *pb.CreateVolumeGroupOpts) (*model.VolumeGroupSpec, error) {
	// Get the storage drivers and do some initializations.
	d.Driver = drivers.Init(opt.GetDriverName())
	defer drivers.Clean(d.Driver)

	log.Info("Creating group...", opt.GetId())

	// NOTE Opt parameter requires complete volumegroup information, because driver may use it.
	group, err := db.C.GetVolumeGroup(c.NewContextFormJson(opt.GetContext()), opt.GetId())
	if err != nil {
		return nil, err
	}

	groupUpdate, notImplementErr, err := d.Driver.CreateVolumeGroup(opt, group)
	if err != nil {
		db.C.UpdateStatus(c.NewContextFormJson(opt.GetContext()), group, model.VOLUMEGROUP_ERROR)
		log.Error("When calling volume driver to create volume group:", err)
		return nil, err
	}

	if notImplementErr == model.ErrorNotImplemented {
		groupUpdate = &model.VolumeGroupSpec{
			BaseModel: &model.BaseModel{
				Id: opt.GetId(),
			},
			Status: model.VOLUMEGROUP_AVAILABLE,
		}
	}

	if groupUpdate != nil && groupUpdate.Status == model.VOLUMEGROUP_ERROR {
		msg := "Error occurred when creating volume group" + opt.GetId()
		log.Error(msg)
		db.C.UpdateStatus(c.NewContextFormJson(opt.GetContext()), group, model.VOLUMEGROUP_ERROR)
		return nil, errors.New(msg)
	}

	group.Status = model.VOLUMEGROUP_AVAILABLE
	group.CreatedAt = time.Now().Format(constants.TimeFormat)
	group.PoolId = opt.GetPoolId()
	db.C.UpdateStatus(c.NewContextFormJson(opt.GetContext()), group, group.Status)
	log.Info("Create group successfully.")

	return group, nil
}

func (d *DockHub) UpdateVolumeGroup(opt *pb.UpdateVolumeGroupOpts) error {
	add := true
	addVolumesRef, err := d.getVolumesForGroup(opt, opt.AddVolumes, add)
	if err != nil {
		return err
	}
	add = false
	removeVolumesRef, err := d.getVolumesForGroup(opt, opt.RemoveVolumes, add)
	if err != nil {
		return err
	}

	// Get the storage drivers and do some initializations.
	d.Driver = drivers.Init(opt.GetDriverName())
	defer drivers.Clean(d.Driver)

	log.Info("Calling volume driver to update volume group...")

	group, err := db.C.GetVolumeGroup(c.NewContextFormJson(opt.GetContext()), opt.GetId())
	if err != nil {
		return err
	}

	groupUpdate, addVolumesUpdate, removeVolumesUpdate, notImplementErr, err := d.Driver.UpdateVolumeGroup(opt, group, addVolumesRef, removeVolumesRef)
	// Group update faild...
	if err != nil {
		err = db.C.UpdateStatus(c.NewContextFormJson(opt.GetContext()), group, model.VOLUMEGROUP_ERROR)
		if err != nil {
			return err
		}

		for _, addVol := range addVolumesRef {
			if err = db.C.UpdateStatus(c.NewContextFormJson(opt.GetContext()), addVol, model.VOLUME_ERROR); err != nil {
				return err
			}
		}
		for _, remVol := range removeVolumesRef {
			if err = db.C.UpdateStatus(c.NewContextFormJson(opt.GetContext()), remVol, model.VOLUME_ERROR); err != nil {
				return err
			}
		}
		return errors.New("Error occured when updating group" + opt.GetId() + "," + err.Error())
	}

	if notImplementErr == model.ErrorNotImplemented {
		groupUpdate, addVolumesUpdate, removeVolumesUpdate = nil, nil, nil
	}

	// Group update successfully...
	// Update volumes return from driver, because volumes somewhere may be modified by driver.
	var volumesToUpdate []*model.VolumeSpec
	if addVolumesUpdate != nil {
		for _, v := range addVolumesUpdate {
			volumesToUpdate = append(volumesToUpdate, v)
		}
	}
	if removeVolumesUpdate != nil {
		for _, v := range removeVolumesUpdate {
			volumesToUpdate = append(volumesToUpdate, v)
		}
	}
	if len(volumesToUpdate) > 0 {
		db.C.VolumesToUpdate(c.NewContextFormJson(opt.GetContext()), volumesToUpdate)
	}

	if groupUpdate != nil {
		if groupUpdate.Status == model.VOLUMEGROUP_ERROR {
			msg := "Error occurred when updating volume group " + opt.GetId()
			log.Error(msg)
			return errors.New(msg)
		}
	}

	for _, addVol := range addVolumesRef {
		addVol.GroupId = opt.GetId()
		if _, err = db.C.UpdateVolume(c.NewContextFormJson(opt.GetContext()), addVol); err != nil {
			return err
		}
	}
	for _, remVol := range removeVolumesRef {
		remVol.GroupId = ""
		if _, err = db.C.UpdateVolume(c.NewContextFormJson(opt.GetContext()), remVol); err != nil {
			return err
		}
	}
	if err = db.C.UpdateStatus(c.NewContextFormJson(opt.GetContext()), group, model.VOLUMEGROUP_AVAILABLE); err != nil {
		return err
	}

	log.Info("Update group " + opt.GetId() + "successfully.")
	return nil
}

func (d *DockHub) getVolumesForGroup(opt *pb.UpdateVolumeGroupOpts, volumes []string, add bool) ([]*model.VolumeSpec, error) {
	var volumesRef []*model.VolumeSpec
	for _, v := range volumes {
		vol, err := db.C.GetVolume(c.NewContextFormJson(opt.GetContext()), v)
		if err != nil {
			log.Error("Update group failed", err)
			return nil, err
		}
		if add == true && vol.Status != model.VOLUME_AVAILABLE && vol.Status != model.VOLUME_IN_USE {
			msg := "Update group failed, wrong status for volume " + vol.Id + vol.Status
			log.Error(msg)
			return nil, errors.New(msg)
		}
		if add == false && vol.Status != model.VOLUME_AVAILABLE && vol.Status != model.VOLUME_IN_USE && vol.Status != model.VOLUME_ERROR && vol.Status != model.VOLUEM_ERROR_DELETING {
			msg := "Update group failed, wrong status for volume " + vol.Id + vol.Status
			log.Error(msg)
			return nil, errors.New(msg)
		}
		volumesRef = append(volumesRef, vol)
	}
	return volumesRef, nil
}

func (d *DockHub) DeleteVolumeGroup(opt *pb.DeleteVolumeGroupOpts) error {
	volumes, err := db.C.ListVolumesByGroupId(c.NewContextFormJson(opt.GetContext()), opt.GetId())
	if err != nil {
		return err
	}

	for _, vol := range volumes {
		if vol.AttachStatus == model.VOLUME_ATTACHED {
			msg := "Volume " + vol.Id + " is still attached, need to detach first."
			return errors.New(msg)
		}
	}

	group, err := db.C.GetVolumeGroup(c.NewContextFormJson(opt.GetContext()), opt.GetId())
	if err != nil {
		return err
	}

	// Get the storage drivers and do some initializations.
	d.Driver = drivers.Init(opt.GetDriverName())
	defer drivers.Clean(d.Driver)
	log.Info("Calling volume driver to delete volume group...")

	groupUpdate, volumesUpdate, notImplementErr, err := d.Driver.DeleteVolumeGroup(opt, group, volumes)
	if err != nil {
		db.C.UpdateStatus(c.NewContextFormJson(opt.GetContext()), group, model.VOLUMEGROUP_ERROR)
		// If driver returns none for volumesUpdate, set volume status to error.
		if volumesUpdate == nil {
			for _, v := range volumes {
				v.Status = model.VOLUME_ERROR
			}
			db.C.UpdateStatus(c.NewContextFormJson(opt.GetContext()), volumes, "")
		}
		return err
	}

	if notImplementErr == model.ErrorNotImplemented {
		groupUpdate, volumesUpdate = d.deleteGroupGeneric(d.Driver, group, volumes)
	}

	if volumesUpdate != nil {
		for _, v := range volumesUpdate {
			if (v.Status == model.VOLUME_ERROR || v.Status == model.VOLUEM_ERROR_DELETING) && (groupUpdate.Status != model.VOLUMEGROUP_ERROR_DELETING || groupUpdate.Status != model.VOLUMEGROUP_ERROR) {
				groupUpdate.Status = v.Status
				break
			}
		}

		db.C.UpdateStatus(c.NewContextFormJson(opt.GetContext()), volumesUpdate, "")

	}

	if groupUpdate != nil {
		if groupUpdate.Status == model.VOLUMEGROUP_ERROR || groupUpdate.Status == model.VOLUMEGROUP_ERROR_DELETING {
			msg := "Delete group failed"
			log.Error(msg)
			return errors.New(msg)
		}
		db.C.UpdateStatus(c.NewContextFormJson(opt.GetContext()), groupUpdate, groupUpdate.Status)
	}

	if err = db.C.DeleteVolumeGroup(c.NewContextFormJson(opt.GetContext()), group.Id); err != nil {
		msg := "Delete volume group failed" + err.Error()
		log.Error(msg)
		return errors.New(msg)
	}

	log.Info("Delete group successfully.")
	return nil
}

func (d *DockHub) deleteGroupGeneric(driver drivers.VolumeDriver, group *model.VolumeGroupSpec, volumes []*model.VolumeSpec) (*model.VolumeGroupSpec, []*model.VolumeSpec) {
	//Delete a group and volumes in the group
	var volumesUpdate []*model.VolumeSpec
	groupUpdate := &model.VolumeGroupSpec{
		BaseModel: &model.BaseModel{
			Id: group.Id,
		},
		Status: group.Status,
	}

	for _, volumeRef := range volumes {
		v := &model.VolumeSpec{
			BaseModel: &model.BaseModel{
				Id: volumeRef.Id,
			},
		}
		if err := driver.DeleteVolume(&pb.DeleteVolumeOpts{Metadata: volumeRef.Metadata}); err != nil {
			v.Status = model.VOLUME_ERROR
			groupUpdate.Status = model.VOLUMEGROUP_ERROR
			log.Error("Error occurred when delete volume " + volumeRef.Id + " from group.")
		} else {
			v.Status = model.VOLUME_DELETED
		}
		volumesUpdate = append(volumesUpdate, v)
	}

	return groupUpdate, volumesUpdate
}
