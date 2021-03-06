// Code generated by go-swagger; DO NOT EDIT.

package msg

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewMsgUploadParams creates a new MsgUploadParams object
// with the default values initialized.
func NewMsgUploadParams() MsgUploadParams {
	var ()
	return MsgUploadParams{}
}

// MsgUploadParams contains all the bound params for the msg upload operation
// typically these are obtained from a http.Request
//
// swagger:parameters msg/upload/
type MsgUploadParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*add/edit/delete
	  In: formData
	*/
	Action *string
	/*
	  In: formData
	*/
	ID *int64
	/*
	  In: formData
	*/
	Status *int64
	/*content
	  In: formData
	*/
	SubTitle *string
	/*
	  In: formData
	*/
	Times *int64
	/*title
	  In: formData
	*/
	Title *string
	/*当前登录用户id
	  In: formData
	*/
	Userid *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *MsgUploadParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		if err != http.ErrNotMultipart {
			return err
		} else if err := r.ParseForm(); err != nil {
			return err
		}
	}
	fds := runtime.Values(r.Form)

	fdAction, fdhkAction, _ := fds.GetOK("action")
	if err := o.bindAction(fdAction, fdhkAction, route.Formats); err != nil {
		res = append(res, err)
	}

	fdID, fdhkID, _ := fds.GetOK("id")
	if err := o.bindID(fdID, fdhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	fdStatus, fdhkStatus, _ := fds.GetOK("status")
	if err := o.bindStatus(fdStatus, fdhkStatus, route.Formats); err != nil {
		res = append(res, err)
	}

	fdSubTitle, fdhkSubTitle, _ := fds.GetOK("sub_title")
	if err := o.bindSubTitle(fdSubTitle, fdhkSubTitle, route.Formats); err != nil {
		res = append(res, err)
	}

	fdTimes, fdhkTimes, _ := fds.GetOK("times")
	if err := o.bindTimes(fdTimes, fdhkTimes, route.Formats); err != nil {
		res = append(res, err)
	}

	fdTitle, fdhkTitle, _ := fds.GetOK("title")
	if err := o.bindTitle(fdTitle, fdhkTitle, route.Formats); err != nil {
		res = append(res, err)
	}

	fdUserid, fdhkUserid, _ := fds.GetOK("userid")
	if err := o.bindUserid(fdUserid, fdhkUserid, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *MsgUploadParams) bindAction(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Action = &raw

	return nil
}

func (o *MsgUploadParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("id", "formData", "int64", raw)
	}
	o.ID = &value

	return nil
}

func (o *MsgUploadParams) bindStatus(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("status", "formData", "int64", raw)
	}
	o.Status = &value

	return nil
}

func (o *MsgUploadParams) bindSubTitle(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.SubTitle = &raw

	return nil
}

func (o *MsgUploadParams) bindTimes(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("times", "formData", "int64", raw)
	}
	o.Times = &value

	return nil
}

func (o *MsgUploadParams) bindTitle(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Title = &raw

	return nil
}

func (o *MsgUploadParams) bindUserid(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("userid", "formData", "int64", raw)
	}
	o.Userid = &value

	return nil
}
