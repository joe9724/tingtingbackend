// Code generated by go-swagger; DO NOT EDIT.

package relation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	models "tingtingbackend/models"
)

// NewNrRelationBookTaglistEditParams creates a new NrRelationBookTaglistEditParams object
// with the default values initialized.
func NewNrRelationBookTaglistEditParams() NrRelationBookTaglistEditParams {
	var ()
	return NrRelationBookTaglistEditParams{}
}

// NrRelationBookTaglistEditParams contains all the bound params for the relation book taglist edit operation
// typically these are obtained from a http.Request
//
// swagger:parameters /relation/book/taglist/edit
type NrRelationBookTaglistEditParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*/relation/book/taglist/edit
	  In: body
	*/
	Body *models.Body2
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *NrRelationBookTaglistEditParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Body2
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}