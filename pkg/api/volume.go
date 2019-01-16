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
This module implements a entry into the OpenSDS northbound service.

*/

package api

import (
	"encoding/json"
	"fmt"
	log "github.com/golang/glog"
	"github.com/opensds/opensds/pkg/api/policy"
	c "github.com/opensds/opensds/pkg/context"
	"github.com/opensds/opensds/pkg/controller/client"
	pb "github.com/opensds/opensds/pkg/controller/proto"
	"github.com/opensds/opensds/pkg/db"
	"github.com/opensds/opensds/pkg/model"
	. "github.com/opensds/opensds/pkg/utils/config"
	"golang.org/x/net/context"
)

func NewVolumePortal() *VolumePortal {
	return &VolumePortal{
		CtrClient: client.NewClient(),
	}
}

type VolumePortal struct {
	BasePortal

	CtrClient client.Client
}

func (v *VolumePortal) CreateVolume() {
	if !policy.Authorize(v.Ctx, "volume:create") {
		return
	}
	ctx := c.GetContext(v.Ctx)
	var volume = model.VolumeSpec{
		BaseModel: &model.BaseModel{},
	}

	// Unmarshal the request body
	if err := json.NewDecoder(v.Ctx.Request.Body).Decode(&volume); err != nil {
		reason := fmt.Sprintf("Parse volume request body failed: %s", err.Error())
		v.Ctx.Output.SetStatus(model.ErrorBadRequest)
		v.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}
	// NOTE:It will create a volume entry into the database and initialize its status
	// as "creating". It will not wait for the real volume creation to complete
	// and will return result immediately.
	result, err := CreateVolumeDBEntry(ctx, &volume)
	if err != nil {
		reason := fmt.Sprintf("Create volume failed: %s", err.Error())
		v.Ctx.Output.SetStatus(model.ErrorBadRequest)
		v.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	// Marshal the result.
	body, _ := json.Marshal(result)
	v.Ctx.Output.Body(body)
	v.Ctx.Output.SetStatus(StatusAccepted)

	// NOTE:The real volume creation process.
	// Volume creation request is sent to the Dock. Dock will update volume status to "available"
	// after volume creation is completed.
	if err := v.CtrClient.Connect(CONF.OsdsLet.ApiEndpoint); err != nil {
		log.Error("When connecting controller client:", err)
		return
	}
	defer v.CtrClient.Close()

	opt := &pb.CreateVolumeOpts{
		Message: string(body),
		Context: ctx.ToJson(),
	}
	if _, err = v.CtrClient.CreateVolume(context.Background(), opt); err != nil {
		log.Error("Create volume failed in controller service:", err)
		return
	}

	return
}

func (v *VolumePortal) ListVolumes() {
	if !policy.Authorize(v.Ctx, "volume:list") {
		return
	}
	// Call db api module to handle list volumes request.
	m, err := v.GetParameters()
	if err != nil {
		reason := fmt.Sprintf("List volumes failed: %s", err.Error())
		v.Ctx.Output.SetStatus(model.ErrorBadRequest)
		v.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	result, err := db.C.ListVolumesWithFilter(c.GetContext(v.Ctx), m)
	if err != nil {
		reason := fmt.Sprintf("List volumes failed: %s", err.Error())
		v.Ctx.Output.SetStatus(model.ErrorBadRequest)
		v.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	// Marshal the result.
	body, _ := json.Marshal(result)
	v.Ctx.Output.Body(body)
	v.Ctx.Output.SetStatus(StatusOK)

	return
}

func (v *VolumePortal) GetVolume() {
	if !policy.Authorize(v.Ctx, "volume:get") {
		return
	}
	id := v.Ctx.Input.Param(":volumeId")

	// Call db api module to handle get volume request.
	result, err := db.C.GetVolume(c.GetContext(v.Ctx), id)
	if err != nil {
		reason := fmt.Sprintf("Get volume failed: %s", err.Error())
		v.Ctx.Output.SetStatus(model.ErrorBadRequest)
		v.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	// Marshal the result.
	body, _ := json.Marshal(result)
	v.Ctx.Output.Body(body)
	v.Ctx.Output.SetStatus(StatusOK)

	return
}

func (v *VolumePortal) UpdateVolume() {
	if !policy.Authorize(v.Ctx, "volume:update") {
		return
	}
	var volume = model.VolumeSpec{
		BaseModel: &model.BaseModel{},
	}

	id := v.Ctx.Input.Param(":volumeId")
	if err := json.NewDecoder(v.Ctx.Request.Body).Decode(&volume); err != nil {
		reason := fmt.Sprintf("Parse volume request body failed: %s", err.Error())
		v.Ctx.Output.SetStatus(model.ErrorBadRequest)
		v.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	volume.Id = id
	result, err := db.C.UpdateVolume(c.GetContext(v.Ctx), &volume)
	if err != nil {
		reason := fmt.Sprintf("Update volume failed: %s", err.Error())
		v.Ctx.Output.SetStatus(model.ErrorBadRequest)
		v.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	// Marshal the result.
	body, _ := json.Marshal(result)
	v.Ctx.Output.Body(body)
	v.Ctx.Output.SetStatus(StatusOK)

	return
}

// ExtendVolume ...
func (v *VolumePortal) ExtendVolume() {
	if !policy.Authorize(v.Ctx, "volume:extend") {
		return
	}
	ctx := c.GetContext(v.Ctx)
	var extendRequestBody = model.ExtendVolumeSpec{}

	if err := json.NewDecoder(v.Ctx.Request.Body).Decode(&extendRequestBody); err != nil {
		reason := fmt.Sprintf("Parse volume request body failed: %s", err.Error())
		v.Ctx.Output.SetStatus(model.ErrorBadRequest)
		v.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	id := v.Ctx.Input.Param(":volumeId")
	// NOTE:It will update the the status of the volume waiting for expansion in
	// the database to "extending" and return the result immediately.
	result, err := ExtendVolumeDBEntry(ctx, id)
	if err != nil {
		reason := fmt.Sprintf("Extend volume failed: %s", err.Error())
		v.Ctx.Output.SetStatus(model.ErrorBadRequest)
		v.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	// Marshal the result.
	body, _ := json.Marshal(result)
	v.Ctx.Output.Body(body)
	v.Ctx.Output.SetStatus(StatusAccepted)

	// NOTE:The real volume extension process.
	// Volume extension request is sent to the Dock. Dock will update volume status to "available"
	// after volume extension is completed.
	if err := v.CtrClient.Connect(CONF.OsdsLet.ApiEndpoint); err != nil {
		log.Error("When connecting controller client:", err)
		return
	}
	defer v.CtrClient.Close()

	body, _ = json.Marshal(&extendRequestBody)
	opt := &pb.ExtendVolumeOpts{
		Id:      id,
		Message: string(body),
		Context: ctx.ToJson(),
	}
	if _, err = v.CtrClient.ExtendVolume(context.Background(), opt); err != nil {
		log.Error("Extend volume failed in controller service:", err)
		return
	}

	return
}

func (v *VolumePortal) DeleteVolume() {
	if !policy.Authorize(v.Ctx, "volume:delete") {
		return
	}
	ctx := c.GetContext(v.Ctx)

	var err error
	id := v.Ctx.Input.Param(":volumeId")
	volume, err := db.C.GetVolume(ctx, id)
	if err != nil {
		reason := fmt.Sprintf("Get volume failed: %s", err.Error())
		v.Ctx.Output.SetStatus(model.ErrorBadRequest)
		v.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	// NOTE:It will update the the status of the volume waiting for deletion in
	// the database to "deleting" and return the result immediately.
	if err = DeleteVolumeDBEntry(ctx, volume); err != nil {
		reason := fmt.Sprintf("Delete volume failed: %v", err.Error())
		v.Ctx.Output.SetStatus(model.ErrorBadRequest)
		v.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}
	v.Ctx.Output.SetStatus(StatusAccepted)

	// NOTE:The real volume deletion process.
	// Volume deletion request is sent to the Dock. Dock will delete volume from driver
	// and database or update volume status to "errorDeleting" if deletion from driver faild.
	if err := v.CtrClient.Connect(CONF.OsdsLet.ApiEndpoint); err != nil {
		log.Error("When connecting controller client:", err)
		return
	}
	defer v.CtrClient.Close()

	body, _ := json.Marshal(volume)
	opt := &pb.DeleteVolumeOpts{
		Message: string(body),
		Context: ctx.ToJson(),
	}
	if _, err = v.CtrClient.DeleteVolume(context.Background(), opt); err != nil {
		log.Error("Delete volume failed in controller service:", err)
		return
	}

	return
}

func NewVolumeAttachmentPortal() *VolumeAttachmentPortal {
	return &VolumeAttachmentPortal{
		CtrClient: client.NewClient(),
	}
}

type VolumeAttachmentPortal struct {
	BasePortal

	CtrClient client.Client
}

func (v *VolumeAttachmentPortal) CreateVolumeAttachment() {
	if !policy.Authorize(v.Ctx, "volume:create_attachment") {
		return
	}
	ctx := c.GetContext(v.Ctx)
	var attachment = model.VolumeAttachmentSpec{
		BaseModel: &model.BaseModel{},
	}

	if err := json.NewDecoder(v.Ctx.Request.Body).Decode(&attachment); err != nil {
		reason := fmt.Sprintf("Parse volume attachment request body failed: %s", err.Error())
		v.Ctx.Output.SetStatus(model.ErrorBadRequest)
		v.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	// NOTE:It will create a volume attachment entry into the database and initialize its status
	// as "creating". It will not wait for the real volume attachment creation to complete
	// and will return result immediately.
	result, err := CreateVolumeAttachmentDBEntry(ctx, &attachment)
	if err != nil {
		reason := fmt.Sprintf("Create volume attachment failed: %s", err.Error())
		v.Ctx.Output.SetStatus(model.ErrorBadRequest)
		v.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	// Marshal the result.
	body, _ := json.Marshal(result)
	v.Ctx.Output.Body(body)
	v.Ctx.Output.SetStatus(StatusAccepted)

	// NOTE:The real volume attachment creation process.
	// Volume attachment creation request is sent to the Dock. Dock will update volume attachment status to "available"
	// after volume attachment creation is completed.
	if err := v.CtrClient.Connect(CONF.OsdsLet.ApiEndpoint); err != nil {
		log.Error("When connecting controller client:", err)
		return
	}
	defer v.CtrClient.Close()

	opt := &pb.CreateVolumeAttachmentOpts{
		Message: string(body),
		Context: ctx.ToJson(),
	}
	if _, err = v.CtrClient.CreateVolumeAttachment(context.Background(), opt); err != nil {
		log.Error("Create volume attachment failed in controller service:", err)
		return
	}

	return
}

func (v *VolumeAttachmentPortal) ListVolumeAttachments() {
	if !policy.Authorize(v.Ctx, "volume:list_attachments") {
		return
	}

	m, err := v.GetParameters()
	if err != nil {
		reason := fmt.Sprintf("List volume attachments failed: %s", err.Error())
		v.Ctx.Output.SetStatus(model.ErrorBadRequest)
		v.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	result, err := db.C.ListVolumeAttachmentsWithFilter(c.GetContext(v.Ctx), m)
	if err != nil {
		reason := fmt.Sprintf("List volume attachments failed: %s", err.Error())
		v.Ctx.Output.SetStatus(model.ErrorBadRequest)
		v.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	// Marshal the result.
	body, _ := json.Marshal(result)
	v.Ctx.Output.Body(body)
	v.Ctx.Output.SetStatus(StatusOK)

	return
}

func (v *VolumeAttachmentPortal) GetVolumeAttachment() {
	if !policy.Authorize(v.Ctx, "volume:get_attachment") {
		return
	}
	id := v.Ctx.Input.Param(":attachmentId")

	result, err := db.C.GetVolumeAttachment(c.GetContext(v.Ctx), id)
	if err != nil {
		reason := fmt.Sprintf("Get volume attachment failed: %s", err.Error())
		v.Ctx.Output.SetStatus(model.ErrorBadRequest)
		v.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	// Marshal the result.
	body, _ := json.Marshal(result)
	v.Ctx.Output.Body(body)
	v.Ctx.Output.SetStatus(StatusOK)

	return
}

func (v *VolumeAttachmentPortal) UpdateVolumeAttachment() {
	if !policy.Authorize(v.Ctx, "volume:update_attachment") {
		return
	}
	var attachment = model.VolumeAttachmentSpec{
		BaseModel: &model.BaseModel{},
	}
	id := v.Ctx.Input.Param(":attachmentId")

	if err := json.NewDecoder(v.Ctx.Request.Body).Decode(&attachment); err != nil {
		reason := fmt.Sprintf("Parse volume attachment request body failed: %s", err.Error())
		v.Ctx.Output.SetStatus(model.ErrorBadRequest)
		v.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}
	attachment.Id = id

	result, err := db.C.UpdateVolumeAttachment(c.GetContext(v.Ctx), id, &attachment)
	if err != nil {
		reason := fmt.Sprintf("Update volume attachment failed: %s", err.Error())
		v.Ctx.Output.SetStatus(model.ErrorBadRequest)
		v.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	// Marshal the result.
	body, _ := json.Marshal(result)
	v.Ctx.Output.Body(body)
	v.Ctx.Output.SetStatus(StatusOK)

	return
}

func (v *VolumeAttachmentPortal) DeleteVolumeAttachment() {
	if !policy.Authorize(v.Ctx, "volume:delete_attachment") {
		return
	}

	ctx := c.GetContext(v.Ctx)

	id := v.Ctx.Input.Param(":attachmentId")
	attachment, err := db.C.GetVolumeAttachment(ctx, id)
	if err != nil {
		reason := fmt.Sprintf("Get volume attachment failed: %s", err.Error())
		v.Ctx.Output.SetStatus(model.ErrorBadRequest)
		v.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}
	// NOTE:It will not wait for the real volume attachment deletion to complete
	// and will return ok immediately.
	v.Ctx.Output.SetStatus(StatusAccepted)

	// NOTE:The real volume attachment deletion process.
	// Volume attachment deletion request is sent to the Dock. Dock will delete volume attachment from database
	// or update its status to "errorDeleting" if volume connection termination failed.
	if err := v.CtrClient.Connect(CONF.OsdsLet.ApiEndpoint); err != nil {
		log.Error("When connecting controller client:", err)
		return
	}
	defer v.CtrClient.Close()

	body, _ := json.Marshal(attachment)
	opt := &pb.DeleteVolumeAttachmentOpts{
		Message: string(body),
		Context: ctx.ToJson(),
	}
	if _, err = v.CtrClient.DeleteVolumeAttachment(context.Background(), opt); err != nil {
		log.Error("Delete volume attachment failed in controller service:", err)
		return
	}

	return
}

func NewVolumeSnapshotPortal() *VolumeSnapshotPortal {
	return &VolumeSnapshotPortal{
		CtrClient: client.NewClient(),
	}
}

type VolumeSnapshotPortal struct {
	BasePortal

	CtrClient client.Client
}

func (v *VolumeSnapshotPortal) CreateVolumeSnapshot() {
	if !policy.Authorize(v.Ctx, "snapshot:create") {
		return
	}
	ctx := c.GetContext(v.Ctx)
	var snapshot = model.VolumeSnapshotSpec{
		BaseModel: &model.BaseModel{},
	}

	if err := json.NewDecoder(v.Ctx.Request.Body).Decode(&snapshot); err != nil {
		reason := fmt.Sprintf("Parse volume snapshot request body failed: %s", err.Error())
		v.Ctx.Output.SetStatus(model.ErrorBadRequest)
		v.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	// NOTE:It will create a volume snapshot entry into the database and initialize its status
	// as "creating". It will not wait for the real volume snapshot creation to complete
	// and will return result immediately.
	result, err := CreateVolumeSnapshotDBEntry(ctx, &snapshot)
	if err != nil {
		reason := fmt.Sprintf("Create volume snapshot failed: %s", err.Error())
		v.Ctx.Output.SetStatus(model.ErrorBadRequest)
		v.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	// Marshal the result.
	body, _ := json.Marshal(result)
					v.Ctx.Output.Body(body)
					v.Ctx.Output.SetStatus(StatusAccepted)

	// NOTE:The real volume snapshot creation process.
	// Volume snapshot creation request is sent to the Dock. Dock will update volume snapshot status to "available"
	// after volume snapshot creation complete.
	if err := v.CtrClient.Connect(CONF.OsdsLet.ApiEndpoint); err != nil {
		log.Error("When connecting controller client:", err)
		return
	}
	defer v.CtrClient.Close()

	opt := &pb.CreateVolumeSnapshotOpts{
		Message: string(body),
		Context: ctx.ToJson(),
	}
	if _, err = v.CtrClient.CreateVolumeSnapshot(context.Background(), opt); err != nil {
		log.Error("Create volume snapthot failed in controller service:", err)
		return
	}

	return
}

func (v *VolumeSnapshotPortal) ListVolumeSnapshots() {
	if !policy.Authorize(v.Ctx, "snapshot:list") {
		return
	}
	m, err := v.GetParameters()
	if err != nil {
		reason := fmt.Sprintf("List volume snapshots failed: %s", err.Error())
		v.Ctx.Output.SetStatus(model.ErrorBadRequest)
		v.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	result, err := db.C.ListVolumeSnapshotsWithFilter(c.GetContext(v.Ctx), m)
	if err != nil {
		reason := fmt.Sprintf("List volume snapshots failed: %s", err.Error())
		v.Ctx.Output.SetStatus(model.ErrorBadRequest)
		v.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	// Marshal the result.
	body, _ := json.Marshal(result)
	v.Ctx.Output.Body(body)
	v.Ctx.Output.SetStatus(StatusOK)

	return
}

func (v *VolumeSnapshotPortal) GetVolumeSnapshot() {
	if !policy.Authorize(v.Ctx, "snapshot:get") {
		return
	}
	id := v.Ctx.Input.Param(":snapshotId")

	result, err := db.C.GetVolumeSnapshot(c.GetContext(v.Ctx), id)
	if err != nil {
		reason := fmt.Sprintf("Get volume snapshot failed: %s", err.Error())
		v.Ctx.Output.SetStatus(model.ErrorBadRequest)
		v.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	// Marshal the result.
	body, _ := json.Marshal(result)
	v.Ctx.Output.Body(body)
	v.Ctx.Output.SetStatus(StatusOK)

	return
}

func (v *VolumeSnapshotPortal) UpdateVolumeSnapshot() {
	if !policy.Authorize(v.Ctx, "snapshot:update") {
		return
	}
	var snapshot = model.VolumeSnapshotSpec{
		BaseModel: &model.BaseModel{},
	}

	id := v.Ctx.Input.Param(":snapshotId")

	if err := json.NewDecoder(v.Ctx.Request.Body).Decode(&snapshot); err != nil {
		reason := fmt.Sprintf("Parse volume snapshot request body failed: %s", err.Error())
		v.Ctx.Output.SetStatus(model.ErrorBadRequest)
		v.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}
	snapshot.Id = id

	result, err := db.C.UpdateVolumeSnapshot(c.GetContext(v.Ctx), id, &snapshot)
	if err != nil {
		reason := fmt.Sprintf("Update volume snapshot failed: %s", err.Error())
		v.Ctx.Output.SetStatus(model.ErrorBadRequest)
		v.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	// Marshal the result.
	body, _ := json.Marshal(result)
	v.Ctx.Output.Body(body)
	v.Ctx.Output.SetStatus(StatusOK)

	return
}

func (v *VolumeSnapshotPortal) DeleteVolumeSnapshot() {
	if !policy.Authorize(v.Ctx, "snapshot:delete") {
		return
	}
	ctx := c.GetContext(v.Ctx)
	id := v.Ctx.Input.Param(":snapshotId")

	snapshot, err := db.C.GetVolumeSnapshot(ctx, id)
	if err != nil {
		reason := fmt.Sprintf("Get volume snapshot failed: %s", err.Error())
		v.Ctx.Output.SetStatus(model.ErrorBadRequest)
		v.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	// NOTE:It will update the the status of the volume snapshot waiting for deletion in
	// the database to "deleting" and return the result immediately.
	err = DeleteVolumeSnapshotDBEntry(ctx, snapshot)
	if err != nil {
		reason := fmt.Sprintf("Delete volume snapshot failed: %v", err.Error())
		v.Ctx.Output.SetStatus(model.ErrorBadRequest)
		v.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	// NOTE:The real volume snapshot deletion process.
	// Volume snapshot deletion request is sent to the Dock. Dock will delete volume snapshot from driver and
	// database or update its status to "errorDeleting" if volume snapshot deletion from driver failed.
	if err := v.CtrClient.Connect(CONF.OsdsLet.ApiEndpoint); err != nil {
		log.Error("When connecting controller client:", err)
		return
	}
	defer v.CtrClient.Close()

	body, _ := json.Marshal(snapshot)
	opt := &pb.DeleteVolumeSnapshotOpts{
		Message: string(body),
		Context: ctx.ToJson(),
	}
	if _, err = v.CtrClient.DeleteVolumeSnapshot(context.Background(), opt); err != nil {
		log.Error("Delete volume snapthot failed in controller service:", err)
		return
	}

	v.Ctx.Output.SetStatus(StatusAccepted)
	return
}
