// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AppVersion app version
// swagger:model AppVersion
type AppVersion struct {

	// client
	Client string `json:"client,omitempty"`

	// download Url
	DownloadURL string `json:"downloadUrl,omitempty"`

	// force
	Force float64 `json:"force,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// msg
	Msg string `json:"msg,omitempty"`

	// number
	Number float64 `json:"number,omitempty"`

	// status
	Status int64 `json:"status,omitempty"`
}

// Validate validates this app version
func (m *AppVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AppVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppVersion) UnmarshalBinary(b []byte) error {
	var res AppVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}