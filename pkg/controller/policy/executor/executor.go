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
This module implements the policy-based scheduling by parsing storage
profiles configured by admin.

*/

package executor

import (
	"errors"
	"log"

	pb "github.com/opensds/opensds/pkg/grpc/opensds"
)

type AsynchronizedExecutor interface {
	Init(in string) error
	Asynchronized() error
}

type AsynchronizedWorkflow map[string]AsynchronizedExecutor

func RegisterAsynchronizedWorkflow(vr *pb.DockRequest, tags map[string]string, in string) (AsynchronizedWorkflow, error) {
	var asynWorkflow = AsynchronizedWorkflow{}

	for key := range tags {
		switch key {
		case "intervalSnapshot":
			ise := &IntervalSnapshotExecutor{
				Request:  vr,
				Interval: tags[key],
			}

			if err := ise.Init(in); err != nil {
				log.Printf("[Error] When register async policy %s: %v\n", key, err)
				return asynWorkflow, err
			} else {
				asynWorkflow[key] = ise
			}
		case "deleteSnapshotPolicy":
			ise := &DeleteSnapshotExecutor{
				Request: vr,
			}

			if err := ise.Init(in); err != nil {
				log.Printf("[Error] When register async policy %s: %v\n", key, err)
				return asynWorkflow, err
			} else {
				asynWorkflow[key] = ise
			}
		}
	}

	log.Println("[Info] Register asynchronized work flow success, awf =", asynWorkflow)
	return asynWorkflow, nil
}

func ExecuteAsynchronizedWorkflow(asynWorkflow AsynchronizedWorkflow) error {
	for key := range asynWorkflow {
		if asynWorkflow[key] == nil {
			return errors.New("Could not execute the policy " + key)
		}
		return asynWorkflow[key].Asynchronized()
	}
	return nil
}

type SynchronizedExecutor interface {
	Init() error
	Synchronized() error
}

type SynchronizedWorkflow map[string]SynchronizedExecutor

func RegisterSynchronizedWorkflow(vr *pb.DockRequest, tags map[string]string) (SynchronizedWorkflow, error) {
	return SynchronizedWorkflow{}, nil
}

func ExecuteSynchronizedWorkflow(synWorkflow SynchronizedWorkflow) error {
	for key := range synWorkflow {
		if synWorkflow[key] == nil {
			return errors.New("Could not execute the policy " + key)
		}
		if err := synWorkflow[key].Synchronized(); err != nil {
			return err
		}
	}
	return nil
}
