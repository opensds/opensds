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

// ExtendVolumeSpec Extends the size of a volume to a requested size, in gibibytes (GiB).
// swagger:model ExtendVolumeSpec
type ExtendVolumeSpec struct {

	// new size
	// Required: true
	NewSize *int64 `json:"newSize"`
}

// Validate validates this extend volume spec
func (m *ExtendVolumeSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNewSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExtendVolumeSpec) validateNewSize(formats strfmt.Registry) error {

	if err := validate.Required("newSize", "body", m.NewSize); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExtendVolumeSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExtendVolumeSpec) UnmarshalBinary(b []byte) error {
	var res ExtendVolumeSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
