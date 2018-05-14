// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import context "github.com/opensds/opensds/pkg/context"

import mock "github.com/stretchr/testify/mock"
import model "github.com/opensds/opensds/pkg/model"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// AddExtraProperty provides a mock function with given fields: ctx, prfID, ext
func (_m *Client) AddExtraProperty(ctx *context.Context, prfID string, ext model.ExtraSpec) (*model.ExtraSpec, error) {
	ret := _m.Called(ctx, prfID, ext)

	var r0 *model.ExtraSpec
	if rf, ok := ret.Get(0).(func(*context.Context, string, model.ExtraSpec) *model.ExtraSpec); ok {
		r0 = rf(ctx, prfID, ext)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ExtraSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, string, model.ExtraSpec) error); ok {
		r1 = rf(ctx, prfID, ext)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateDock provides a mock function with given fields: ctx, dck
func (_m *Client) CreateDock(ctx *context.Context, dck *model.DockSpec) (*model.DockSpec, error) {
	ret := _m.Called(ctx, dck)

	var r0 *model.DockSpec
	if rf, ok := ret.Get(0).(func(*context.Context, *model.DockSpec) *model.DockSpec); ok {
		r0 = rf(ctx, dck)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.DockSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, *model.DockSpec) error); ok {
		r1 = rf(ctx, dck)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePool provides a mock function with given fields: ctx, pol
func (_m *Client) CreatePool(ctx *context.Context, pol *model.StoragePoolSpec) (*model.StoragePoolSpec, error) {
	ret := _m.Called(ctx, pol)

	var r0 *model.StoragePoolSpec
	if rf, ok := ret.Get(0).(func(*context.Context, *model.StoragePoolSpec) *model.StoragePoolSpec); ok {
		r0 = rf(ctx, pol)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.StoragePoolSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, *model.StoragePoolSpec) error); ok {
		r1 = rf(ctx, pol)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateProfile provides a mock function with given fields: ctx, prf
func (_m *Client) CreateProfile(ctx *context.Context, prf *model.ProfileSpec) (*model.ProfileSpec, error) {
	ret := _m.Called(ctx, prf)

	var r0 *model.ProfileSpec
	if rf, ok := ret.Get(0).(func(*context.Context, *model.ProfileSpec) *model.ProfileSpec); ok {
		r0 = rf(ctx, prf)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ProfileSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, *model.ProfileSpec) error); ok {
		r1 = rf(ctx, prf)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateReplication provides a mock function with given fields: ctx, replication
func (_m *Client) CreateReplication(ctx *context.Context, replication *model.ReplicationSpec) (*model.ReplicationSpec, error) {
	ret := _m.Called(ctx, replication)

	var r0 *model.ReplicationSpec
	if rf, ok := ret.Get(0).(func(*context.Context, *model.ReplicationSpec) *model.ReplicationSpec); ok {
		r0 = rf(ctx, replication)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ReplicationSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, *model.ReplicationSpec) error); ok {
		r1 = rf(ctx, replication)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVolume provides a mock function with given fields: ctx, vol
func (_m *Client) CreateVolume(ctx *context.Context, vol *model.VolumeSpec) (*model.VolumeSpec, error) {
	ret := _m.Called(ctx, vol)

	var r0 *model.VolumeSpec
	if rf, ok := ret.Get(0).(func(*context.Context, *model.VolumeSpec) *model.VolumeSpec); ok {
		r0 = rf(ctx, vol)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.VolumeSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, *model.VolumeSpec) error); ok {
		r1 = rf(ctx, vol)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVolumeAttachment provides a mock function with given fields: ctx, attachment
func (_m *Client) CreateVolumeAttachment(ctx *context.Context, attachment *model.VolumeAttachmentSpec) (*model.VolumeAttachmentSpec, error) {
	ret := _m.Called(ctx, attachment)

	var r0 *model.VolumeAttachmentSpec
	if rf, ok := ret.Get(0).(func(*context.Context, *model.VolumeAttachmentSpec) *model.VolumeAttachmentSpec); ok {
		r0 = rf(ctx, attachment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.VolumeAttachmentSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, *model.VolumeAttachmentSpec) error); ok {
		r1 = rf(ctx, attachment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVolumeGroup provides a mock function with given fields: ctx, vg
func (_m *Client) CreateVolumeGroup(ctx *context.Context, vg *model.VolumeGroupSpec) (*model.VolumeGroupSpec, error) {
	ret := _m.Called(ctx, vg)

	var r0 *model.VolumeGroupSpec
	if rf, ok := ret.Get(0).(func(*context.Context, *model.VolumeGroupSpec) *model.VolumeGroupSpec); ok {
		r0 = rf(ctx, vg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.VolumeGroupSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, *model.VolumeGroupSpec) error); ok {
		r1 = rf(ctx, vg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVolumeSnapshot provides a mock function with given fields: ctx, vs
func (_m *Client) CreateVolumeSnapshot(ctx *context.Context, vs *model.VolumeSnapshotSpec) (*model.VolumeSnapshotSpec, error) {
	ret := _m.Called(ctx, vs)

	var r0 *model.VolumeSnapshotSpec
	if rf, ok := ret.Get(0).(func(*context.Context, *model.VolumeSnapshotSpec) *model.VolumeSnapshotSpec); ok {
		r0 = rf(ctx, vs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.VolumeSnapshotSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, *model.VolumeSnapshotSpec) error); ok {
		r1 = rf(ctx, vs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteDock provides a mock function with given fields: ctx, dckID
func (_m *Client) DeleteDock(ctx *context.Context, dckID string) error {
	ret := _m.Called(ctx, dckID)

	var r0 error
	if rf, ok := ret.Get(0).(func(*context.Context, string) error); ok {
		r0 = rf(ctx, dckID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeletePool provides a mock function with given fields: ctx, polID
func (_m *Client) DeletePool(ctx *context.Context, polID string) error {
	ret := _m.Called(ctx, polID)

	var r0 error
	if rf, ok := ret.Get(0).(func(*context.Context, string) error); ok {
		r0 = rf(ctx, polID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteProfile provides a mock function with given fields: ctx, prfID
func (_m *Client) DeleteProfile(ctx *context.Context, prfID string) error {
	ret := _m.Called(ctx, prfID)

	var r0 error
	if rf, ok := ret.Get(0).(func(*context.Context, string) error); ok {
		r0 = rf(ctx, prfID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteReplication provides a mock function with given fields: ctx, replicationId
func (_m *Client) DeleteReplication(ctx *context.Context, replicationId string) error {
	ret := _m.Called(ctx, replicationId)

	var r0 error
	if rf, ok := ret.Get(0).(func(*context.Context, string) error); ok {
		r0 = rf(ctx, replicationId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteVolume provides a mock function with given fields: ctx, volID
func (_m *Client) DeleteVolume(ctx *context.Context, volID string) error {
	ret := _m.Called(ctx, volID)

	var r0 error
	if rf, ok := ret.Get(0).(func(*context.Context, string) error); ok {
		r0 = rf(ctx, volID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteVolumeAttachment provides a mock function with given fields: ctx, attachmentId
func (_m *Client) DeleteVolumeAttachment(ctx *context.Context, attachmentId string) error {
	ret := _m.Called(ctx, attachmentId)

	var r0 error
	if rf, ok := ret.Get(0).(func(*context.Context, string) error); ok {
		r0 = rf(ctx, attachmentId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteVolumeGroup provides a mock function with given fields: ctx, vgId
func (_m *Client) DeleteVolumeGroup(ctx *context.Context, vgId string) error {
	ret := _m.Called(ctx, vgId)

	var r0 error
	if rf, ok := ret.Get(0).(func(*context.Context, string) error); ok {
		r0 = rf(ctx, vgId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteVolumeSnapshot provides a mock function with given fields: ctx, snapshotID
func (_m *Client) DeleteVolumeSnapshot(ctx *context.Context, snapshotID string) error {
	ret := _m.Called(ctx, snapshotID)

	var r0 error
	if rf, ok := ret.Get(0).(func(*context.Context, string) error); ok {
		r0 = rf(ctx, snapshotID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExtendVolume provides a mock function with given fields: ctx, vol
func (_m *Client) ExtendVolume(ctx *context.Context, vol *model.VolumeSpec) (*model.VolumeSpec, error) {
	ret := _m.Called(ctx, vol)

	var r0 *model.VolumeSpec
	if rf, ok := ret.Get(0).(func(*context.Context, *model.VolumeSpec) *model.VolumeSpec); ok {
		r0 = rf(ctx, vol)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.VolumeSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, *model.VolumeSpec) error); ok {
		r1 = rf(ctx, vol)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDefaultProfile provides a mock function with given fields: ctx
func (_m *Client) GetDefaultProfile(ctx *context.Context) (*model.ProfileSpec, error) {
	ret := _m.Called(ctx)

	var r0 *model.ProfileSpec
	if rf, ok := ret.Get(0).(func(*context.Context) *model.ProfileSpec); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ProfileSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDock provides a mock function with given fields: ctx, dckID
func (_m *Client) GetDock(ctx *context.Context, dckID string) (*model.DockSpec, error) {
	ret := _m.Called(ctx, dckID)

	var r0 *model.DockSpec
	if rf, ok := ret.Get(0).(func(*context.Context, string) *model.DockSpec); ok {
		r0 = rf(ctx, dckID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.DockSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, string) error); ok {
		r1 = rf(ctx, dckID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDockByPoolId provides a mock function with given fields: ctx, poolId
func (_m *Client) GetDockByPoolId(ctx *context.Context, poolId string) (*model.DockSpec, error) {
	ret := _m.Called(ctx, poolId)

	var r0 *model.DockSpec
	if rf, ok := ret.Get(0).(func(*context.Context, string) *model.DockSpec); ok {
		r0 = rf(ctx, poolId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.DockSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, string) error); ok {
		r1 = rf(ctx, poolId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPool provides a mock function with given fields: ctx, polID
func (_m *Client) GetPool(ctx *context.Context, polID string) (*model.StoragePoolSpec, error) {
	ret := _m.Called(ctx, polID)

	var r0 *model.StoragePoolSpec
	if rf, ok := ret.Get(0).(func(*context.Context, string) *model.StoragePoolSpec); ok {
		r0 = rf(ctx, polID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.StoragePoolSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, string) error); ok {
		r1 = rf(ctx, polID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProfile provides a mock function with given fields: ctx, prfID
func (_m *Client) GetProfile(ctx *context.Context, prfID string) (*model.ProfileSpec, error) {
	ret := _m.Called(ctx, prfID)

	var r0 *model.ProfileSpec
	if rf, ok := ret.Get(0).(func(*context.Context, string) *model.ProfileSpec); ok {
		r0 = rf(ctx, prfID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ProfileSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, string) error); ok {
		r1 = rf(ctx, prfID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReplication provides a mock function with given fields: ctx, replicationId
func (_m *Client) GetReplication(ctx *context.Context, replicationId string) (*model.ReplicationSpec, error) {
	ret := _m.Called(ctx, replicationId)

	var r0 *model.ReplicationSpec
	if rf, ok := ret.Get(0).(func(*context.Context, string) *model.ReplicationSpec); ok {
		r0 = rf(ctx, replicationId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ReplicationSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, string) error); ok {
		r1 = rf(ctx, replicationId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVolume provides a mock function with given fields: ctx, volID
func (_m *Client) GetVolume(ctx *context.Context, volID string) (*model.VolumeSpec, error) {
	ret := _m.Called(ctx, volID)

	var r0 *model.VolumeSpec
	if rf, ok := ret.Get(0).(func(*context.Context, string) *model.VolumeSpec); ok {
		r0 = rf(ctx, volID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.VolumeSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, string) error); ok {
		r1 = rf(ctx, volID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVolumeAttachment provides a mock function with given fields: ctx, attachmentId
func (_m *Client) GetVolumeAttachment(ctx *context.Context, attachmentId string) (*model.VolumeAttachmentSpec, error) {
	ret := _m.Called(ctx, attachmentId)

	var r0 *model.VolumeAttachmentSpec
	if rf, ok := ret.Get(0).(func(*context.Context, string) *model.VolumeAttachmentSpec); ok {
		r0 = rf(ctx, attachmentId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.VolumeAttachmentSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, string) error); ok {
		r1 = rf(ctx, attachmentId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVolumeGroup provides a mock function with given fields: ctx, vgId
func (_m *Client) GetVolumeGroup(ctx *context.Context, vgId string) (*model.VolumeGroupSpec, error) {
	ret := _m.Called(ctx, vgId)

	var r0 *model.VolumeGroupSpec
	if rf, ok := ret.Get(0).(func(*context.Context, string) *model.VolumeGroupSpec); ok {
		r0 = rf(ctx, vgId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.VolumeGroupSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, string) error); ok {
		r1 = rf(ctx, vgId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVolumeSnapshot provides a mock function with given fields: ctx, snapshotID
func (_m *Client) GetVolumeSnapshot(ctx *context.Context, snapshotID string) (*model.VolumeSnapshotSpec, error) {
	ret := _m.Called(ctx, snapshotID)

	var r0 *model.VolumeSnapshotSpec
	if rf, ok := ret.Get(0).(func(*context.Context, string) *model.VolumeSnapshotSpec); ok {
		r0 = rf(ctx, snapshotID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.VolumeSnapshotSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, string) error); ok {
		r1 = rf(ctx, snapshotID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDocks provides a mock function with given fields: ctx
func (_m *Client) ListDocks(ctx *context.Context) ([]*model.DockSpec, error) {
	ret := _m.Called(ctx)

	var r0 []*model.DockSpec
	if rf, ok := ret.Get(0).(func(*context.Context) []*model.DockSpec); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.DockSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDocksWithFilter provides a mock function with given fields: ctx, m
func (_m *Client) ListDocksWithFilter(ctx *context.Context, m map[string][]string) ([]*model.DockSpec, error) {
	ret := _m.Called(ctx, m)

	var r0 []*model.DockSpec
	if rf, ok := ret.Get(0).(func(*context.Context, map[string][]string) []*model.DockSpec); ok {
		r0 = rf(ctx, m)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.DockSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, map[string][]string) error); ok {
		r1 = rf(ctx, m)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListExtraProperties provides a mock function with given fields: ctx, prfID
func (_m *Client) ListExtraProperties(ctx *context.Context, prfID string) (*model.ExtraSpec, error) {
	ret := _m.Called(ctx, prfID)

	var r0 *model.ExtraSpec
	if rf, ok := ret.Get(0).(func(*context.Context, string) *model.ExtraSpec); ok {
		r0 = rf(ctx, prfID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ExtraSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, string) error); ok {
		r1 = rf(ctx, prfID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPools provides a mock function with given fields: ctx
func (_m *Client) ListPools(ctx *context.Context) ([]*model.StoragePoolSpec, error) {
	ret := _m.Called(ctx)

	var r0 []*model.StoragePoolSpec
	if rf, ok := ret.Get(0).(func(*context.Context) []*model.StoragePoolSpec); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.StoragePoolSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPoolsWithFilter provides a mock function with given fields: ctx, m
func (_m *Client) ListPoolsWithFilter(ctx *context.Context, m map[string][]string) ([]*model.StoragePoolSpec, error) {
	ret := _m.Called(ctx, m)

	var r0 []*model.StoragePoolSpec
	if rf, ok := ret.Get(0).(func(*context.Context, map[string][]string) []*model.StoragePoolSpec); ok {
		r0 = rf(ctx, m)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.StoragePoolSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, map[string][]string) error); ok {
		r1 = rf(ctx, m)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProfiles provides a mock function with given fields: ctx
func (_m *Client) ListProfiles(ctx *context.Context) ([]*model.ProfileSpec, error) {
	ret := _m.Called(ctx)

	var r0 []*model.ProfileSpec
	if rf, ok := ret.Get(0).(func(*context.Context) []*model.ProfileSpec); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ProfileSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProfilesWithFilter provides a mock function with given fields: ctx, m
func (_m *Client) ListProfilesWithFilter(ctx *context.Context, m map[string][]string) ([]*model.ProfileSpec, error) {
	ret := _m.Called(ctx, m)

	var r0 []*model.ProfileSpec
	if rf, ok := ret.Get(0).(func(*context.Context, map[string][]string) []*model.ProfileSpec); ok {
		r0 = rf(ctx, m)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ProfileSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, map[string][]string) error); ok {
		r1 = rf(ctx, m)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListReplication provides a mock function with given fields: ctx
func (_m *Client) ListReplication(ctx *context.Context) ([]*model.ReplicationSpec, error) {
	ret := _m.Called(ctx)

	var r0 []*model.ReplicationSpec
	if rf, ok := ret.Get(0).(func(*context.Context) []*model.ReplicationSpec); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ReplicationSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListReplicationWithFilter provides a mock function with given fields: ctx, m
func (_m *Client) ListReplicationWithFilter(ctx *context.Context, m map[string][]string) ([]*model.ReplicationSpec, error) {
	ret := _m.Called(ctx, m)

	var r0 []*model.ReplicationSpec
	if rf, ok := ret.Get(0).(func(*context.Context, map[string][]string) []*model.ReplicationSpec); ok {
		r0 = rf(ctx, m)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ReplicationSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, map[string][]string) error); ok {
		r1 = rf(ctx, m)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSnapshotsByVolumeId provides a mock function with given fields: ctx, volId
func (_m *Client) ListSnapshotsByVolumeId(ctx *context.Context, volId string) ([]*model.VolumeSnapshotSpec, error) {
	ret := _m.Called(ctx, volId)

	var r0 []*model.VolumeSnapshotSpec
	if rf, ok := ret.Get(0).(func(*context.Context, string) []*model.VolumeSnapshotSpec); ok {
		r0 = rf(ctx, volId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.VolumeSnapshotSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, string) error); ok {
		r1 = rf(ctx, volId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListVolumeAttachments provides a mock function with given fields: ctx, volumeId
func (_m *Client) ListVolumeAttachments(ctx *context.Context, volumeId string) ([]*model.VolumeAttachmentSpec, error) {
	ret := _m.Called(ctx, volumeId)

	var r0 []*model.VolumeAttachmentSpec
	if rf, ok := ret.Get(0).(func(*context.Context, string) []*model.VolumeAttachmentSpec); ok {
		r0 = rf(ctx, volumeId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.VolumeAttachmentSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, string) error); ok {
		r1 = rf(ctx, volumeId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListVolumeAttachmentsWithFilter provides a mock function with given fields: ctx, m
func (_m *Client) ListVolumeAttachmentsWithFilter(ctx *context.Context, m map[string][]string) ([]*model.VolumeAttachmentSpec, error) {
	ret := _m.Called(ctx, m)

	var r0 []*model.VolumeAttachmentSpec
	if rf, ok := ret.Get(0).(func(*context.Context, map[string][]string) []*model.VolumeAttachmentSpec); ok {
		r0 = rf(ctx, m)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.VolumeAttachmentSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, map[string][]string) error); ok {
		r1 = rf(ctx, m)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListVolumeGroups provides a mock function with given fields: ctx
func (_m *Client) ListVolumeGroups(ctx *context.Context) ([]*model.VolumeGroupSpec, error) {
	ret := _m.Called(ctx)

	var r0 []*model.VolumeGroupSpec
	if rf, ok := ret.Get(0).(func(*context.Context) []*model.VolumeGroupSpec); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.VolumeGroupSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListVolumeSnapshots provides a mock function with given fields: ctx
func (_m *Client) ListVolumeSnapshots(ctx *context.Context) ([]*model.VolumeSnapshotSpec, error) {
	ret := _m.Called(ctx)

	var r0 []*model.VolumeSnapshotSpec
	if rf, ok := ret.Get(0).(func(*context.Context) []*model.VolumeSnapshotSpec); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.VolumeSnapshotSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListVolumeSnapshotsWithFilter provides a mock function with given fields: ctx, m
func (_m *Client) ListVolumeSnapshotsWithFilter(ctx *context.Context, m map[string][]string) ([]*model.VolumeSnapshotSpec, error) {
	ret := _m.Called(ctx, m)

	var r0 []*model.VolumeSnapshotSpec
	if rf, ok := ret.Get(0).(func(*context.Context, map[string][]string) []*model.VolumeSnapshotSpec); ok {
		r0 = rf(ctx, m)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.VolumeSnapshotSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, map[string][]string) error); ok {
		r1 = rf(ctx, m)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListVolumes provides a mock function with given fields: ctx
func (_m *Client) ListVolumes(ctx *context.Context) ([]*model.VolumeSpec, error) {
	ret := _m.Called(ctx)

	var r0 []*model.VolumeSpec
	if rf, ok := ret.Get(0).(func(*context.Context) []*model.VolumeSpec); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.VolumeSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListVolumesByGroupId provides a mock function with given fields: ctx, vgId
func (_m *Client) ListVolumesByGroupId(ctx *context.Context, vgId string) ([]*model.VolumeSpec, error) {
	ret := _m.Called(ctx, vgId)

	var r0 []*model.VolumeSpec
	if rf, ok := ret.Get(0).(func(*context.Context, string) []*model.VolumeSpec); ok {
		r0 = rf(ctx, vgId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.VolumeSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, string) error); ok {
		r1 = rf(ctx, vgId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListVolumesWithFilter provides a mock function with given fields: ctx, m
func (_m *Client) ListVolumesWithFilter(ctx *context.Context, m map[string][]string) ([]*model.VolumeSpec, error) {
	ret := _m.Called(ctx, m)

	var r0 []*model.VolumeSpec
	if rf, ok := ret.Get(0).(func(*context.Context, map[string][]string) []*model.VolumeSpec); ok {
		r0 = rf(ctx, m)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.VolumeSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, map[string][]string) error); ok {
		r1 = rf(ctx, m)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveExtraProperty provides a mock function with given fields: ctx, prfID, extraKey
func (_m *Client) RemoveExtraProperty(ctx *context.Context, prfID string, extraKey string) error {
	ret := _m.Called(ctx, prfID, extraKey)

	var r0 error
	if rf, ok := ret.Get(0).(func(*context.Context, string, string) error); ok {
		r0 = rf(ctx, prfID, extraKey)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDock provides a mock function with given fields: ctx, dckID, name, desp
func (_m *Client) UpdateDock(ctx *context.Context, dckID string, name string, desp string) (*model.DockSpec, error) {
	ret := _m.Called(ctx, dckID, name, desp)

	var r0 *model.DockSpec
	if rf, ok := ret.Get(0).(func(*context.Context, string, string, string) *model.DockSpec); ok {
		r0 = rf(ctx, dckID, name, desp)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.DockSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, string, string, string) error); ok {
		r1 = rf(ctx, dckID, name, desp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePool provides a mock function with given fields: ctx, polID, name, desp, usedCapacity, used
func (_m *Client) UpdatePool(ctx *context.Context, polID string, name string, desp string, usedCapacity int64, used bool) (*model.StoragePoolSpec, error) {
	ret := _m.Called(ctx, polID, name, desp, usedCapacity, used)

	var r0 *model.StoragePoolSpec
	if rf, ok := ret.Get(0).(func(*context.Context, string, string, string, int64, bool) *model.StoragePoolSpec); ok {
		r0 = rf(ctx, polID, name, desp, usedCapacity, used)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.StoragePoolSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, string, string, string, int64, bool) error); ok {
		r1 = rf(ctx, polID, name, desp, usedCapacity, used)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProfile provides a mock function with given fields: ctx, prfID, input
func (_m *Client) UpdateProfile(ctx *context.Context, prfID string, input *model.ProfileSpec) (*model.ProfileSpec, error) {
	ret := _m.Called(ctx, prfID, input)

	var r0 *model.ProfileSpec
	if rf, ok := ret.Get(0).(func(*context.Context, string, *model.ProfileSpec) *model.ProfileSpec); ok {
		r0 = rf(ctx, prfID, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ProfileSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, string, *model.ProfileSpec) error); ok {
		r1 = rf(ctx, prfID, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateReplication provides a mock function with given fields: ctx, replicationId, input
func (_m *Client) UpdateReplication(ctx *context.Context, replicationId string, input *model.ReplicationSpec) (*model.ReplicationSpec, error) {
	ret := _m.Called(ctx, replicationId, input)

	var r0 *model.ReplicationSpec
	if rf, ok := ret.Get(0).(func(*context.Context, string, *model.ReplicationSpec) *model.ReplicationSpec); ok {
		r0 = rf(ctx, replicationId, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ReplicationSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, string, *model.ReplicationSpec) error); ok {
		r1 = rf(ctx, replicationId, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatus provides a mock function with given fields: ctx, object, status
func (_m *Client) UpdateStatus(ctx *context.Context, object interface{}, status string) error {
	ret := _m.Called(ctx, object, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(*context.Context, interface{}, string) error); ok {
		r0 = rf(ctx, object, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateVolume provides a mock function with given fields: ctx, vol
func (_m *Client) UpdateVolume(ctx *context.Context, vol *model.VolumeSpec) (*model.VolumeSpec, error) {
	ret := _m.Called(ctx, vol)

	var r0 *model.VolumeSpec
	if rf, ok := ret.Get(0).(func(*context.Context, *model.VolumeSpec) *model.VolumeSpec); ok {
		r0 = rf(ctx, vol)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.VolumeSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, *model.VolumeSpec) error); ok {
		r1 = rf(ctx, vol)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateVolumeAttachment provides a mock function with given fields: ctx, attachmentId, attachment
func (_m *Client) UpdateVolumeAttachment(ctx *context.Context, attachmentId string, attachment *model.VolumeAttachmentSpec) (*model.VolumeAttachmentSpec, error) {
	ret := _m.Called(ctx, attachmentId, attachment)

	var r0 *model.VolumeAttachmentSpec
	if rf, ok := ret.Get(0).(func(*context.Context, string, *model.VolumeAttachmentSpec) *model.VolumeAttachmentSpec); ok {
		r0 = rf(ctx, attachmentId, attachment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.VolumeAttachmentSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, string, *model.VolumeAttachmentSpec) error); ok {
		r1 = rf(ctx, attachmentId, attachment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateVolumeGroup provides a mock function with given fields: ctx, vg
func (_m *Client) UpdateVolumeGroup(ctx *context.Context, vg *model.VolumeGroupSpec) (*model.VolumeGroupSpec, error) {
	ret := _m.Called(ctx, vg)

	var r0 *model.VolumeGroupSpec
	if rf, ok := ret.Get(0).(func(*context.Context, *model.VolumeGroupSpec) *model.VolumeGroupSpec); ok {
		r0 = rf(ctx, vg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.VolumeGroupSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, *model.VolumeGroupSpec) error); ok {
		r1 = rf(ctx, vg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateVolumeSnapshot provides a mock function with given fields: ctx, snapshotID, vs
func (_m *Client) UpdateVolumeSnapshot(ctx *context.Context, snapshotID string, vs *model.VolumeSnapshotSpec) (*model.VolumeSnapshotSpec, error) {
	ret := _m.Called(ctx, snapshotID, vs)

	var r0 *model.VolumeSnapshotSpec
	if rf, ok := ret.Get(0).(func(*context.Context, string, *model.VolumeSnapshotSpec) *model.VolumeSnapshotSpec); ok {
		r0 = rf(ctx, snapshotID, vs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.VolumeSnapshotSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, string, *model.VolumeSnapshotSpec) error); ok {
		r1 = rf(ctx, snapshotID, vs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VolumesToUpdate provides a mock function with given fields: ctx, volumeList
func (_m *Client) VolumesToUpdate(ctx *context.Context, volumeList []*model.VolumeSpec) ([]*model.VolumeSpec, error) {
	ret := _m.Called(ctx, volumeList)

	var r0 []*model.VolumeSpec
	if rf, ok := ret.Get(0).(func(*context.Context, []*model.VolumeSpec) []*model.VolumeSpec); ok {
		r0 = rf(ctx, volumeList)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.VolumeSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, []*model.VolumeSpec) error); ok {
		r1 = rf(ctx, volumeList)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
