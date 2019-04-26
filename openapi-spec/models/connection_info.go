// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ConnectionInfo ConnectionInfo is a structure for all properties of connection when creating a volume attachment.
// swagger:model ConnectionInfo
type ConnectionInfo struct {

	// data
	Data map[string]interface{} `json:"data,omitempty"`

	// driver volume type
	DriverVolumeType string `json:"driverVolumeType,omitempty"`

	// extra properties
	ExtraProperties map[string]interface{} `json:"extraProperties,omitempty"`
}

// Validate validates this connection info
func (m *ConnectionInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConnectionInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConnectionInfo) UnmarshalBinary(b []byte) error {
	var res ConnectionInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
