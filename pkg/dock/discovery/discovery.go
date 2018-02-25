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
This module implements the entry into operations of storageDock module.

*/

package discovery

import (
	"fmt"
	"os"

	log "github.com/golang/glog"
	"github.com/opensds/opensds/contrib/drivers"
	"github.com/opensds/opensds/pkg/db"
	"github.com/opensds/opensds/pkg/model"
	. "github.com/opensds/opensds/pkg/utils/config"
	"github.com/satori/go.uuid"
)

// NewDiscoverer method creates a new DockDiscoverer and return its pointer.
func NewDiscoverer() *DockDiscoverer {
	return &DockDiscoverer{
		c: db.C,
	}
}

// DockDiscoverer is a struct for exposing some operations of service discovery.
type DockDiscoverer struct {
	dcks []*model.DockSpec
	pols []*model.StoragePoolSpec

	c db.Client
}

// Init
func (dd *DockDiscoverer) Init() error {
	// Load resource from specified file
	bm := GetBackendsMap()
	host, err := os.Hostname()
	if err != nil {
		log.Error("When get os hostname:", err)
		return err
	}

	for _, v := range CONF.EnabledBackends {
		b := bm[v]
		if b.Name == "" {
			continue
		}

		dck := &model.DockSpec{
			BaseModel: &model.BaseModel{
				Id: uuid.NewV5(uuid.NamespaceOID, host+":"+b.DriverName).String(),
			},
			Name:        b.Name,
			Description: b.Description,
			DriverName:  b.DriverName,
			Endpoint:    CONF.OsdsDock.ApiEndpoint,
		}
		dd.dcks = append(dd.dcks, dck)
	}

	return nil
}

// Discover
func (dd *DockDiscoverer) Discover(d drivers.VolumeDriver) error {

	for _, dck := range dd.dcks {
		//Call function of StorageDrivers configured by storage drivers.
		d = drivers.Init(dck.DriverName)
		defer drivers.Clean(d)
		pols, err := d.ListPools()
		if err != nil {
			log.Error("Call driver to list pools failed:", err)
			continue
		}

		if len(pols) == 0 {
			log.Warningf("The pool of dock %s is empty!\n", dck.Id)
		}

		for _, pol := range pols {
			log.Infof("Backend %s discovered pool %s", dck.DriverName, pol.Name)
			pol.DockId = dck.Id
		}
		dd.pols = append(dd.pols, pols...)
	}
	if len(dd.pols) == 0 {
		return fmt.Errorf("There is no pool can be found.")
	}
	return nil
}

// Store
func (dd *DockDiscoverer) Store() error {
	var err error

	// Store dock resources in database.
	for _, dck := range dd.dcks {
		// Call db module to create dock resource.
		if _, err = dd.c.CreateDock(dck); err != nil {
			log.Errorf("When create dock %s in db: %v\n", dck.Name, err)
			return err
		}
	}

	// Store pool resources in database.
	for _, pol := range dd.pols {
		// Call db module to create pool resource.
		if _, err = dd.c.CreatePool(pol); err != nil {
			log.Errorf("When create pool %s in db: %v\n", pol.Name, err)
			return err
		}
	}

	return err
}
