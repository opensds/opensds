// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/opensds/opensds/pkg/model"
import proto "github.com/opensds/opensds/pkg/model/proto"

// ReplicationDriver is an autogenerated mock type for the ReplicationDriver type
type ReplicationDriver struct {
	mock.Mock
}

// CreateReplication provides a mock function with given fields: opt
func (_m *ReplicationDriver) CreateReplication(opt *proto.CreateReplicationOpts) (*model.ReplicationSpec, error) {
	ret := _m.Called(opt)

	var r0 *model.ReplicationSpec
	if rf, ok := ret.Get(0).(func(*proto.CreateReplicationOpts) *model.ReplicationSpec); ok {
		r0 = rf(opt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ReplicationSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*proto.CreateReplicationOpts) error); ok {
		r1 = rf(opt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteReplication provides a mock function with given fields: opt
func (_m *ReplicationDriver) DeleteReplication(opt *proto.DeleteReplicationOpts) error {
	ret := _m.Called(opt)

	var r0 error
	if rf, ok := ret.Get(0).(func(*proto.DeleteReplicationOpts) error); ok {
		r0 = rf(opt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DisableReplication provides a mock function with given fields: opt
func (_m *ReplicationDriver) DisableReplication(opt *proto.DisableReplicationOpts) error {
	ret := _m.Called(opt)

	var r0 error
	if rf, ok := ret.Get(0).(func(*proto.DisableReplicationOpts) error); ok {
		r0 = rf(opt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnableReplication provides a mock function with given fields: opt
func (_m *ReplicationDriver) EnableReplication(opt *proto.EnableReplicationOpts) error {
	ret := _m.Called(opt)

	var r0 error
	if rf, ok := ret.Get(0).(func(*proto.EnableReplicationOpts) error); ok {
		r0 = rf(opt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FailoverReplication provides a mock function with given fields: opt
func (_m *ReplicationDriver) FailoverReplication(opt *proto.FailoverReplicationOpts) error {
	ret := _m.Called(opt)

	var r0 error
	if rf, ok := ret.Get(0).(func(*proto.FailoverReplicationOpts) error); ok {
		r0 = rf(opt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Setup provides a mock function with given fields:
func (_m *ReplicationDriver) Setup() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Unset provides a mock function with given fields:
func (_m *ReplicationDriver) Unset() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
