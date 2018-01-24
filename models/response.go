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

// Response response
// swagger:model Response
type Response struct {

	// code
	// Required: true
	Code *int64 `json:"code"`

	// msg
	// Required: true
	Msg *string `json:"msg"`

	TotalCount int64 `json:"totalCount"`
}

// Validate validates this response
func (m *Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMsg(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Response) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *Response) validateMsg(formats strfmt.Registry) error {

	if err := validate.Required("msg", "body", m.Msg); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Response) UnmarshalBinary(b []byte) error {
	var res Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
