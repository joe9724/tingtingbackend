// Code generated by go-swagger; DO NOT EDIT.

package album

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewAlbumUploadParams creates a new AlbumUploadParams object
// with the default values initialized.
func NewAlbumUploadParams() AlbumUploadParams {
	var ()
	return AlbumUploadParams{}
}

// AlbumUploadParams contains all the bound params for the album upload operation
// typically these are obtained from a http.Request
//
// swagger:parameters album/upload/
type AlbumUploadParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: formData
	*/
	Content *string
	/*The file to upload.
	  Required: true
	  In: formData
	*/
	Cover io.ReadCloser
	/*The file to upload.
	  Required: true
	  In: formData
	*/
	Icon io.ReadCloser
	/*专辑价格
	  In: formData
	*/
	PriceValue *float64
	/*
	  Required: true
	  In: formData
	*/
	SubTitle string
	/*
	  In: formData
	*/
	Summary *string
	/*Description of file contents.
	  Required: true
	  In: formData
	*/
	Title string
	/*当前登录用户id
	  In: formData
	*/
	Userid *int64

	IconUrl string

	Status *int64

	AlbumId *int64

	Grade *int64

}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *AlbumUploadParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
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

	fdContent, fdhkContent, _ := fds.GetOK("content")
	if err := o.bindContent(fdContent, fdhkContent, route.Formats); err != nil {
		res = append(res, err)
	}

	fdIconUrl, fdhkIconUrl, _ := fds.GetOK("iconUrl")
	if err := o.bindIconUrl(fdIconUrl, fdhkIconUrl, route.Formats); err != nil {
		res = append(res, err)
	}

	cover, coverHeader, err := r.FormFile("cover")
	if err != nil && err != http.ErrMissingFile {
		res = append(res, errors.New(400, "reading file %q failed: %v", "cover", err))
	} else if err == http.ErrMissingFile {
		// no-op for missing but optional file parameter
	} else if err := o.bindCover(cover, coverHeader); err != nil {
		res = append(res, err)
	} else {
		o.Cover = &runtime.File{Data: cover, Header: coverHeader}
	}

	icon, iconHeader, err := r.FormFile("icon")
	if err != nil && err != http.ErrMissingFile {
		res = append(res, errors.New(400, "reading file %q failed: %v", "icon", err))
	} else if err == http.ErrMissingFile {
		// no-op for missing but optional file parameter
	} else if err := o.bindIcon(icon, iconHeader); err != nil {
		res = append(res, err)
	} else {
		o.Icon = &runtime.File{Data: icon, Header: iconHeader}
	}


	fdPriceValue, fdhkPriceValue, _ := fds.GetOK("priceValue")
	if err := o.bindPriceValue(fdPriceValue, fdhkPriceValue, route.Formats); err != nil {
		res = append(res, err)
	}

	fdStatus, fdhkStatus, _ := fds.GetOK("status")
	if err := o.bindStatus(fdStatus, fdhkStatus, route.Formats); err != nil {
		res = append(res, err)
	}

	fdSubTitle, fdhkSubTitle, _ := fds.GetOK("subTitle")
	if err := o.bindSubTitle(fdSubTitle, fdhkSubTitle, route.Formats); err != nil {
		res = append(res, err)
	}

	fdAlbumId, fdhkAlbumId, _ := fds.GetOK("albumId")
	if err := o.bindAlbumId(fdAlbumId, fdhkAlbumId, route.Formats); err != nil {
		res = append(res, err)
	}

	fdGrade, fdhkGrade, _ := fds.GetOK("grade")
	if err := o.bindGrade(fdGrade, fdhkGrade, route.Formats); err != nil {
		res = append(res, err)
	}

	fdSummary, fdhkSummary, _ := fds.GetOK("summary")
	if err := o.bindSummary(fdSummary, fdhkSummary, route.Formats); err != nil {
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

func (o *AlbumUploadParams) bindIconUrl(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.IconUrl = raw

	return nil
}

func (o *AlbumUploadParams) bindContent(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *AlbumUploadParams) bindCover(file multipart.File, header *multipart.FileHeader) error {

	return nil
}

func (o *AlbumUploadParams) bindIcon(file multipart.File, header *multipart.FileHeader) error {

	return nil
}

func (o *AlbumUploadParams) bindPriceValue(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertFloat64(raw)
	if err != nil {
		return errors.InvalidType("priceValue", "formData", "float64", raw)
	}
	o.PriceValue = &value

	return nil
}

func (o *AlbumUploadParams) bindAlbumId(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("albumId", "formData", "int64", raw)
	}
	o.AlbumId = &value

	return nil
}

func (o *AlbumUploadParams) bindGrade(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("grade", "formData", "int64", raw)
	}
	o.Grade = &value

	return nil
}

func (o *AlbumUploadParams) bindSubTitle(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("subTitle", "formData")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("subTitle", "formData", raw); err != nil {
		return err
	}

	o.SubTitle = raw

	return nil
}

func (o *AlbumUploadParams) bindSummary(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Summary = &raw

	return nil
}

func (o *AlbumUploadParams) bindTitle(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("title", "formData")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("title", "formData", raw); err != nil {
		return err
	}

	o.Title = raw

	return nil
}

func (o *AlbumUploadParams) bindUserid(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *AlbumUploadParams) bindStatus(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
