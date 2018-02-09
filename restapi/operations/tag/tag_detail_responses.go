// Code generated by go-swagger; DO NOT EDIT.

package tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "tingtingbackend/models"
)

// TagDetailOKCode is the HTTP code returned for type TagDetailOK
const TagDetailOKCode int = 200

/*TagDetailOK 专辑详情

swagger:response tagDetailOK
*/
type TagDetailOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20033 `json:"body,omitempty"`
}

// NewTagDetailOK creates TagDetailOK with default headers values
func NewTagDetailOK() *TagDetailOK {
	return &TagDetailOK{}
}

// WithPayload adds the payload to the tag detail o k response
func (o *TagDetailOK) WithPayload(payload *models.InlineResponse20033) *TagDetailOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the tag detail o k response
func (o *TagDetailOK) SetPayload(payload *models.InlineResponse20033) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TagDetailOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*TagDetailDefault error

swagger:response tagDetailDefault
*/
type TagDetailDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewTagDetailDefault creates TagDetailDefault with default headers values
func NewTagDetailDefault(code int) *TagDetailDefault {
	if code <= 0 {
		code = 500
	}

	return &TagDetailDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the tag detail default response
func (o *TagDetailDefault) WithStatusCode(code int) *TagDetailDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the tag detail default response
func (o *TagDetailDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the tag detail default response
func (o *TagDetailDefault) WithPayload(payload *models.Error) *TagDetailDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the tag detail default response
func (o *TagDetailDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TagDetailDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}