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

// Banner banner
// swagger:model Banner
type Banner struct {
	ID int64 `json:"id,omitempty" gorm:"AUTO_INCREMENT"`
	// cover
	// Required: true
	Cover *string `json:"cover"`

	// id
	// Required: true
	// ID *int64 `json:"id"`

	// type
	// Required: true
	Type *int64 `json:"type"`

	Name *string `json:"name"`
}

// Validate validates this banner
func (m *Banner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCover(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Banner) validateCover(formats strfmt.Registry) error {

	if err := validate.Required("cover", "body", m.Cover); err != nil {
		return err
	}

	return nil
}

func (m *Banner) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Banner) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Banner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Banner) UnmarshalBinary(b []byte) error {
	var res Banner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
