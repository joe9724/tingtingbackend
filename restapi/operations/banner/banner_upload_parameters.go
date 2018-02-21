// Code generated by go-swagger; DO NOT EDIT.

package banner

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

	strfmt "github.com/go-openapi/strfmt"
)

// NewBannerUploadParams creates a new BannerUploadParams object
// with the default values initialized.
func NewBannerUploadParams() BannerUploadParams {
	var ()
	return BannerUploadParams{}
}

// BannerUploadParams contains all the bound params for the banner upload operation
// typically these are obtained from a http.Request
//
// swagger:parameters banner/upload/
type BannerUploadParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*add/edit
	  In: formData
	*/
	Action *string
	/*
	  In: formData
	*/
	BannerID *int64
	/*The file to upload.
	  In: formData
	*/
	Cover io.ReadCloser
	/*Description of file contents.
	  In: formData
	*/
	Name *string
	/*跳转Id
	  In: formData
	*/
	TargetID *int64
	/*跳转类型
	  In: formData
	*/
	Type *int64
	/*当前登录用户id
	  In: formData
	*/
	Userid *int64
	/*
	  In: formData
	*/
	WebURL *string

	Status *int64

	CoverUrl string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *BannerUploadParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
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

	fdBannerID, fdhkBannerID, _ := fds.GetOK("bannerId")
	if err := o.bindBannerID(fdBannerID, fdhkBannerID, route.Formats); err != nil {
		res = append(res, err)
	}

	fdCoverUrl, fdhkCoverUrl, _ := fds.GetOK("coverUrl")
	if err := o.bindCoverUrl(fdCoverUrl, fdhkCoverUrl, route.Formats); err != nil {
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

	fdName, fdhkName, _ := fds.GetOK("name")
	if err := o.bindName(fdName, fdhkName, route.Formats); err != nil {
		res = append(res, err)
	}

	fdTargetID, fdhkTargetID, _ := fds.GetOK("targetId")
	if err := o.bindTargetID(fdTargetID, fdhkTargetID, route.Formats); err != nil {
		res = append(res, err)
	}

	fdType, fdhkType, _ := fds.GetOK("type")
	if err := o.bindType(fdType, fdhkType, route.Formats); err != nil {
		res = append(res, err)
	}

	fdUserid, fdhkUserid, _ := fds.GetOK("userid")
	if err := o.bindUserid(fdUserid, fdhkUserid, route.Formats); err != nil {
		res = append(res, err)
	}

	fdWebURL, fdhkWebURL, _ := fds.GetOK("webUrl")
	if err := o.bindWebURL(fdWebURL, fdhkWebURL, route.Formats); err != nil {
		res = append(res, err)
	}

	fdStatus, fdhkStatus, _ := fds.GetOK("status")
	if err := o.bindStatus(fdStatus, fdhkStatus, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *BannerUploadParams) bindAction(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *BannerUploadParams) bindBannerID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("bannerId", "formData", "int64", raw)
	}
	o.BannerID = &value

	return nil
}

func (o *BannerUploadParams) bindCover(file multipart.File, header *multipart.FileHeader) error {

	return nil
}

func (o *BannerUploadParams) bindCoverUrl(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.CoverUrl = raw

	return nil
}

func (o *BannerUploadParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Name = &raw

	return nil
}

func (o *BannerUploadParams) bindTargetID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("targetId", "formData", "int64", raw)
	}
	o.TargetID = &value

	return nil
}

func (o *BannerUploadParams) bindType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("type", "formData", "int64", raw)
	}
	o.Type = &value

	return nil
}

func (o *BannerUploadParams) bindStatus(rawData []string, hasKey bool, formats strfmt.Registry) error {
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


func (o *BannerUploadParams) bindUserid(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *BannerUploadParams) bindWebURL(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.WebURL = &raw

	return nil
}
