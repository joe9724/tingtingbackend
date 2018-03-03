// Code generated by go-swagger; DO NOT EDIT.

package web

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

// NewWebUploadParams creates a new WebUploadParams object
// with the default values initialized.
func NewWebUploadParams() WebUploadParams {
	var ()
	return WebUploadParams{}
}

// WebUploadParams contains all the bound params for the web upload operation
// typically these are obtained from a http.Request
//
// swagger:parameters web/upload/
type WebUploadParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*add/edit/delete
	  In: formData
	*/
	Action *string
	/*
	  In: formData
	*/
	Content *string
	/*
	  In: formData
	*/
	IconURL *string
	/*id
	  In: formData
	*/
	ID *int64
	/*Description of file contents.
	  In: formData
	*/
	Title *string
	/*
	  In: formData
	*/
	URL *string
	/*当前登录用户id
	  In: formData
	*/
	Userid *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *WebUploadParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
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

	fdContent, fdhkContent, _ := fds.GetOK("content")
	if err := o.bindContent(fdContent, fdhkContent, route.Formats); err != nil {
		res = append(res, err)
	}

	fdIconURL, fdhkIconURL, _ := fds.GetOK("iconUrl")
	if err := o.bindIconURL(fdIconURL, fdhkIconURL, route.Formats); err != nil {
		res = append(res, err)
	}

	fdID, fdhkID, _ := fds.GetOK("id")
	if err := o.bindID(fdID, fdhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	fdTitle, fdhkTitle, _ := fds.GetOK("title")
	if err := o.bindTitle(fdTitle, fdhkTitle, route.Formats); err != nil {
		res = append(res, err)
	}

	fdURL, fdhkURL, _ := fds.GetOK("url")
	if err := o.bindURL(fdURL, fdhkURL, route.Formats); err != nil {
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

func (o *WebUploadParams) bindAction(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *WebUploadParams) bindContent(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Content = &raw

	return nil
}

func (o *WebUploadParams) bindIconURL(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.IconURL = &raw

	return nil
}

func (o *WebUploadParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *WebUploadParams) bindTitle(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *WebUploadParams) bindURL(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.URL = &raw

	return nil
}

func (o *WebUploadParams) bindUserid(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
