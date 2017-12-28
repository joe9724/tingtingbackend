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

// Chapter chapter
// swagger:model Chapter
type Chapter struct {

	// big cover
	// Required: true
	BigCover *string `json:"bigCover"`

	// duration
	// Required: true
	Duration *int64 `json:"duration"`

	// icon
	// Required: true
	Icon *string `json:"icon"`

	// name
	// Required: true
	Name *string `json:"name"`

	// order
	// Required: true
	Order *int64 `json:"order"`

	// play count
	// Required: true
	PlayCount *int64 `json:"playCount"`

	// show icon
	// Required: true
	ShowIcon *bool `json:"showIcon"`

	// sub title
	// Required: true
	SubTitle *string `json:"subTitle"`

	// update tips
	// Required: true
	UpdateTips *string `json:"updateTips"`

	// url
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this chapter
func (m *Chapter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBigCover(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIcon(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOrder(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePlayCount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateShowIcon(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSubTitle(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUpdateTips(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Chapter) validateBigCover(formats strfmt.Registry) error {

	if err := validate.Required("bigCover", "body", m.BigCover); err != nil {
		return err
	}

	return nil
}

func (m *Chapter) validateDuration(formats strfmt.Registry) error {

	if err := validate.Required("duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *Chapter) validateIcon(formats strfmt.Registry) error {

	if err := validate.Required("icon", "body", m.Icon); err != nil {
		return err
	}

	return nil
}

func (m *Chapter) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Chapter) validateOrder(formats strfmt.Registry) error {

	if err := validate.Required("order", "body", m.Order); err != nil {
		return err
	}

	return nil
}

func (m *Chapter) validatePlayCount(formats strfmt.Registry) error {

	if err := validate.Required("playCount", "body", m.PlayCount); err != nil {
		return err
	}

	return nil
}

func (m *Chapter) validateShowIcon(formats strfmt.Registry) error {

	if err := validate.Required("showIcon", "body", m.ShowIcon); err != nil {
		return err
	}

	return nil
}

func (m *Chapter) validateSubTitle(formats strfmt.Registry) error {

	if err := validate.Required("subTitle", "body", m.SubTitle); err != nil {
		return err
	}

	return nil
}

func (m *Chapter) validateUpdateTips(formats strfmt.Registry) error {

	if err := validate.Required("updateTips", "body", m.UpdateTips); err != nil {
		return err
	}

	return nil
}

func (m *Chapter) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Chapter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Chapter) UnmarshalBinary(b []byte) error {
	var res Chapter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
