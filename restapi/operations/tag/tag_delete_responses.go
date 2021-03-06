// Code generated by go-swagger; DO NOT EDIT.

package tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "tingtingbackend/models"
)

// TagDeleteOKCode is the HTTP code returned for type TagDeleteOK
const TagDeleteOKCode int = 200

/*TagDeleteOK 专辑详情

swagger:response tagDeleteOK
*/
type TagDeleteOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20014 `json:"body,omitempty"`
}

// NewTagDeleteOK creates TagDeleteOK with default headers values
func NewTagDeleteOK() *TagDeleteOK {
	return &TagDeleteOK{}
}

// WithPayload adds the payload to the tag delete o k response
func (o *TagDeleteOK) WithPayload(payload *models.InlineResponse20014) *TagDeleteOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the tag delete o k response
func (o *TagDeleteOK) SetPayload(payload *models.InlineResponse20014) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TagDeleteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*TagDeleteDefault error

swagger:response tagDeleteDefault
*/
type TagDeleteDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewTagDeleteDefault creates TagDeleteDefault with default headers values
func NewTagDeleteDefault(code int) *TagDeleteDefault {
	if code <= 0 {
		code = 500
	}

	return &TagDeleteDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the tag delete default response
func (o *TagDeleteDefault) WithStatusCode(code int) *TagDeleteDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the tag delete default response
func (o *TagDeleteDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the tag delete default response
func (o *TagDeleteDefault) WithPayload(payload *models.Error) *TagDeleteDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the tag delete default response
func (o *TagDeleteDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TagDeleteDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
