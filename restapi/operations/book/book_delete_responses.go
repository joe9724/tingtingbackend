// Code generated by go-swagger; DO NOT EDIT.

package book

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/tingtingbackend/models"
)

// BookDeleteOKCode is the HTTP code returned for type BookDeleteOK
const BookDeleteOKCode int = 200

/*BookDeleteOK 书本详情

swagger:response bookDeleteOK
*/
type BookDeleteOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20017 `json:"body,omitempty"`
}

// NewBookDeleteOK creates BookDeleteOK with default headers values
func NewBookDeleteOK() *BookDeleteOK {
	return &BookDeleteOK{}
}

// WithPayload adds the payload to the book delete o k response
func (o *BookDeleteOK) WithPayload(payload *models.InlineResponse20017) *BookDeleteOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the book delete o k response
func (o *BookDeleteOK) SetPayload(payload *models.InlineResponse20017) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BookDeleteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*BookDeleteDefault error

swagger:response bookDeleteDefault
*/
type BookDeleteDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewBookDeleteDefault creates BookDeleteDefault with default headers values
func NewBookDeleteDefault(code int) *BookDeleteDefault {
	if code <= 0 {
		code = 500
	}

	return &BookDeleteDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the book delete default response
func (o *BookDeleteDefault) WithStatusCode(code int) *BookDeleteDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the book delete default response
func (o *BookDeleteDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the book delete default response
func (o *BookDeleteDefault) WithPayload(payload *models.Error) *BookDeleteDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the book delete default response
func (o *BookDeleteDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BookDeleteDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
