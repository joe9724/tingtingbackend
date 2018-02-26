// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Album album
// swagger:model Album
type Album struct {
	Id int64 `json:"id" gorm:"AUTO_INCREMENT"`
	// author avatar
	AuthorAvatar string `json:"authorAvatar"`

	// author name
	AuthorName string `json:"authorName"`

	// books number
	BooksNumber int64 `json:"booksNumber"`

	// icon
	Icon string `json:"icon"`

	Cover string `json:"cover"`

	// name
	Name string `json:"name"`

	// play count
	PlayCount int64 `json:"playCount"`

	// show icon
	ShowIcon bool `json:"showIcon"`

	// status
	Status int64 `json:"status"`

	// sub category Id
	SubCategoryID int64 `json:"subCategoryId"`

	// sub title
	SubTitle string `json:"subTitle"`

	// summary
	Summary string `json:"summary"`

	// time
	Time int64 `json:"time"`

	// value
	Value float64 `json:"value"`

	HasPushed int64 `json:"hasPushed"`

	Grade int64 `json:"grade"`

	GradeRange string `json:"gradeRange"`

	Times int64 `json:"times"`
}

// Validate validates this album
func (m *Album) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Album) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Album) UnmarshalBinary(b []byte) error {
	var res Album
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
