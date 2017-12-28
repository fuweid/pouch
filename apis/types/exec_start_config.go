// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ExecStartConfig exec start config
// swagger:model ExecStartConfig

type ExecStartConfig struct {

	// ExecStart will first check if it's detached
	Detach bool `json:"Detach,omitempty"`

	// Check if there's a tty
	Tty bool `json:"Tty,omitempty"`
}

/* polymorph ExecStartConfig Detach false */

/* polymorph ExecStartConfig Tty false */

// Validate validates this exec start config
func (m *ExecStartConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ExecStartConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExecStartConfig) UnmarshalBinary(b []byte) error {
	var res ExecStartConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
