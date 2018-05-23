// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SuggestedBook suggested book
// swagger:model SuggestedBook
type SuggestedBook struct {

	// author avatar
	AuthorAvatar string `json:"authorAvatar,omitempty"`

	// author name
	AuthorName string `json:"authorName,omitempty"`

	// clips number
	ClipsNumber int64 `json:"clipsNumber,omitempty"`

	// end time
	EndTime string `json:"endTime"`

	// grade
	Grade int64 `json:"grade,omitempty"`

	// icon
	Icon string `json:"icon,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// play count
	PlayCount int64 `json:"playCount,omitempty"`

	// show icon
	ShowIcon bool `json:"showIcon,omitempty"`

	// start time
	StartTime string `json:"start_time"`

	// status
	Status int64 `json:"status"`

	// sub category Id
	SubCategoryID int64 `json:"subCategoryId,omitempty"`

	// sub title
	SubTitle string `json:"subTitle,omitempty"`

	// summary
	Summary string `json:"summary,omitempty"`

	// time
	Time int64 `json:"time,omitempty"`

	BookId int64 `json:"bookId"`

	Id int64 `json:"id"`
}

// Validate validates this suggested book
func (m *SuggestedBook) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SuggestedBook) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SuggestedBook) UnmarshalBinary(b []byte) error {
	var res SuggestedBook
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
