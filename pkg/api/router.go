// Copyright 2017 The OpenSDS Authors.
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
This module implements a entry into the OpenSDS northbound REST service.

*/

package api

import (
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

const (
	StatusOK       = http.StatusOK
	StatusAccepted = http.StatusAccepted
)

func Run(host string) {

	// add router for v1beta api
	ns :=
		beego.NewNamespace("/v1beta",
			beego.NSCond(func(ctx *context.Context) bool {
				// To judge whether the scheme is legal or not.
				if ctx.Input.Scheme() != "http" && ctx.Input.Scheme() != "https" {
					return false
				}
				return true
			}),

			// List all dock services, including a list of dock object
			beego.NSRouter("/docks", &DockPortal{}, "get:ListDocks"),
			// Show one dock service, including endpoint, driverName and so on
			beego.NSRouter("/docks/:dockId", &DockPortal{}, "get:GetDock"),

			// Profile is a set of policies configured by admin and provided for users
			// CreateProfile, UpdateProfile and DeleteProfile are used for admin only
			// ListProfiles and GetProfile are used for both admin and users
			beego.NSRouter("/profiles", &ProfilePortal{}, "post:CreateProfile;get:ListProfiles"),
			beego.NSRouter("/profiles/:profileId", &ProfilePortal{}, "get:GetProfile;put:UpdateProfile;delete:DeleteProfile"),

			// All operations of extras are used for Admin only
			beego.NSRouter("/profiles/:profileId/extras", &ProfilePortal{}, "post:AddExtraProperty;get:ListExtraProperties"),
			beego.NSRouter("/profiles/:profileId/extras/:extraKey", &ProfilePortal{}, "delete:RemoveExtraProperty"),

			// Pool is the virtual description of backend storage, usually divided into block, file and object,
			// and every pool is atomic, which means every pool contains a specific set of features.
			// ListPools and GetPool are used for checking the status of backend pool, admin only
			beego.NSRouter("/pools", &PoolPortal{}, "get:ListPools"),
			beego.NSRouter("/pools/:poolId", &PoolPortal{}, "get:GetPool"),

			beego.NSNamespace("/block",

				// Volume is the logical description of a piece of storage, which can be directly used by users.
				// All operations of volume can be used for both admin and users.
				beego.NSRouter("/volumes", &VolumePortal{}, "post:CreateVolume;get:ListVolumes"),
				beego.NSRouter("/volumes/:volumeId", &VolumePortal{}, "get:GetVolume;put:UpdateVolume;delete:DeleteVolume"),

				// Creates, shows, lists, unpdates and deletes attachment.
				beego.NSRouter("/attachments", &VolumeAttachmentPortal{}, "post:CreateVolumeAttachment;get:ListVolumeAttachments"),
				beego.NSRouter("/attachments/:attachmentId", &VolumeAttachmentPortal{}, "get:GetVolumeAttachment;put:UpdateVolumeAttachment;delete:DeleteVolumeAttachment"),

				// Snapshot is a point-in-time copy of the data that a volume contains.
				// Creates, shows, lists, unpdates and deletes snapshot.
				beego.NSRouter("/snapshots", &VolumeSnapshotPortal{}, "post:CreateVolumeSnapshot;get:ListVolumeSnapshots"),
				beego.NSRouter("/snapshots/:snapshotId", &VolumeSnapshotPortal{}, "get:GetVolumeSnapshot;put:UpdateVolumeSnapshot;delete:DeleteVolumeSnapshot"),
			),
			// Extend Volume
			beego.NSRouter("/volumes/:volumeId/action", &VolumePortal{}, "post:ExtendVolume"),
		)

	beego.AddNamespace(ns)

	// add router for api version
	beego.Router("/", &VersionPortal{}, "get:ListVersions")
	beego.Router("/:apiVersion", &VersionPortal{}, "get:GetVersion")

	// start service
	beego.Run(host)
}
