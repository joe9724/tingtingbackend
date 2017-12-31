// Code generated by go-swagger; DO NOT EDIT.

package feedback

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

// NewFeedbackDetailParams creates a new FeedbackDetailParams object
// with the default values initialized.
func NewFeedbackDetailParams() FeedbackDetailParams {
	var ()
	return FeedbackDetailParams{}
}

// FeedbackDetailParams contains all the bound params for the feedback detail operation
// typically these are obtained from a http.Request
//
// swagger:parameters feedback/detail
type FeedbackDetailParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*反馈id
	  In: query
	*/
	FeedbackID *string
	/*当前后台登录id
	  In: query
	*/
	Userid *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *FeedbackDetailParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qFeedbackID, qhkFeedbackID, _ := qs.GetOK("feedbackId")
	if err := o.bindFeedbackID(qFeedbackID, qhkFeedbackID, route.Formats); err != nil {
		res = append(res, err)
	}

	qUserid, qhkUserid, _ := qs.GetOK("userid")
	if err := o.bindUserid(qUserid, qhkUserid, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FeedbackDetailParams) bindFeedbackID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.FeedbackID = &raw

	return nil
}

func (o *FeedbackDetailParams) bindUserid(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("userid", "query", "int64", raw)
	}
	o.Userid = &value

	return nil
}
