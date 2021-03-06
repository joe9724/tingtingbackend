// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse2001 inline response 200 1
// swagger:model inline_response_200_1
type InlineResponse2001513 struct {

	// status
	Status *Response `json:"status"`

	// web
	Web *Web `json:"web"`

	Dashboard DashBoard `json:"dashboard"`
}

type DashBoard struct{
	NumberNewUser int64 `json:"number_newuser"`
	NumberTodayBuyAlbums int64 `json:"number_today_buy_albums"`
	MoneyToday int64 `json:"money_today"`
	NumberTodayRecord int64 `json:"number_today_record"`
	NumberMonthBuyAlbums int64 `json:"number_month_buy_albums"`
    NumberTodayAddCategory int64 `json:"number_today_add_category"`
    NumberTodayAddAlbums int64 `json:"number_today_add_albums"`
    NumberTodayAddBook int64 `json:"number_today_add_book"`
    NumberTodayAddChapter int64 `json:"number_today_add_chapter"`
    HotAlbums []HotAlbum `json:"hotalbums"`
    MonthBuyedAlbum []MonthAlbumModel `json:"month_buyed_album"`

}

type HotAlbum struct{
	Name string `json:"name"`
	Percent int64 `json:"percent"`
	PlayCount int64 `json:"play_count"`
}

type MonthAlbumModel struct{
	AlbumCount int64 `json:"album_count"`
	Money int64 `json:"money"`
}

// Validate validates this inline response 200 1
func (m *InlineResponse2001513) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateWeb(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse2001513) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {

		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2001513) validateWeb(formats strfmt.Registry) error {

	if swag.IsZero(m.Web) { // not required
		return nil
	}

	if m.Web != nil {

		if err := m.Web.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("web")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse2001513) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse2001513) UnmarshalBinary(b []byte) error {
	var res InlineResponse2001513
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
