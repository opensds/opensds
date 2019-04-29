// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VolumeSnapshotSpec Snapshot is a description of volume snapshot resource.
// swagger:model VolumeSnapshotSpec
type VolumeSnapshotSpec struct {
	BaseModel

	// description
	Description string `json:"description,omitempty"`

	// metadata
	Metadata interface{} `json:"metadata,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// profile Id
	// Required: true
	ProfileID *string `json:"profileId"`

	// project Id
	// Required: true
	// Read Only: true
	ProjectID string `json:"projectId"`

	// size
	// Read Only: true
	Size int64 `json:"size,omitempty"`

	// status
	// Required: true
	// Read Only: true
	Status string `json:"status"`

	// user Id
	// Read Only: true
	UserID string `json:"userId,omitempty"`

	// volume Id
	// Required: true
	VolumeID *string `json:"volumeId"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VolumeSnapshotSpec) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseModel
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseModel = aO0

	// AO1
	var dataAO1 struct {
		Description string `json:"description,omitempty"`

		Metadata interface{} `json:"metadata,omitempty"`

		Name *string `json:"name"`

		ProfileID *string `json:"profileId"`

		ProjectID string `json:"projectId"`

		Size int64 `json:"size,omitempty"`

		Status string `json:"status"`

		UserID string `json:"userId,omitempty"`

		VolumeID *string `json:"volumeId"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Description = dataAO1.Description

	m.Metadata = dataAO1.Metadata

	m.Name = dataAO1.Name

	m.ProfileID = dataAO1.ProfileID

	m.ProjectID = dataAO1.ProjectID

	m.Size = dataAO1.Size

	m.Status = dataAO1.Status

	m.UserID = dataAO1.UserID

	m.VolumeID = dataAO1.VolumeID

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VolumeSnapshotSpec) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BaseModel)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Description string `json:"description,omitempty"`

		Metadata interface{} `json:"metadata,omitempty"`

		Name *string `json:"name"`

		ProfileID *string `json:"profileId"`

		ProjectID string `json:"projectId"`

		Size int64 `json:"size,omitempty"`

		Status string `json:"status"`

		UserID string `json:"userId,omitempty"`

		VolumeID *string `json:"volumeId"`
	}

	dataAO1.Description = m.Description

	dataAO1.Metadata = m.Metadata

	dataAO1.Name = m.Name

	dataAO1.ProfileID = m.ProfileID

	dataAO1.ProjectID = m.ProjectID

	dataAO1.Size = m.Size

	dataAO1.Status = m.Status

	dataAO1.UserID = m.UserID

	dataAO1.VolumeID = m.VolumeID

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this volume snapshot spec
func (m *VolumeSnapshotSpec) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseModel
	if err := m.BaseModel.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfileID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeSnapshotSpec) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VolumeSnapshotSpec) validateProfileID(formats strfmt.Registry) error {

	if err := validate.Required("profileId", "body", m.ProfileID); err != nil {
		return err
	}

	return nil
}

func (m *VolumeSnapshotSpec) validateProjectID(formats strfmt.Registry) error {

	if err := validate.RequiredString("projectId", "body", string(m.ProjectID)); err != nil {
		return err
	}

	return nil
}

func (m *VolumeSnapshotSpec) validateStatus(formats strfmt.Registry) error {

	if err := validate.RequiredString("status", "body", string(m.Status)); err != nil {
		return err
	}

	return nil
}

func (m *VolumeSnapshotSpec) validateVolumeID(formats strfmt.Registry) error {

	if err := validate.Required("volumeId", "body", m.VolumeID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VolumeSnapshotSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumeSnapshotSpec) UnmarshalBinary(b []byte) error {
	var res VolumeSnapshotSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
