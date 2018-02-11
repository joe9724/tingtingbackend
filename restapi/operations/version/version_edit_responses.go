// Code generated by go-swagger; DO NOT EDIT.

package version

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "tingtingbackendtest/models"
)

// VersionEditOKCode is the HTTP code returned for type VersionEditOK
const VersionEditOKCode int = 200

/*VersionEditOK 返回成功信息

swagger:response versionEditOK
*/
type VersionEditOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20019 `json:"body,omitempty"`
}

// NewVersionEditOK creates VersionEditOK with default headers values
func NewVersionEditOK() *VersionEditOK {
	return &VersionEditOK{}
}

// WithPayload adds the payload to the version edit o k response
func (o *VersionEditOK) WithPayload(payload *models.InlineResponse20019) *VersionEditOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the version edit o k response
func (o *VersionEditOK) SetPayload(payload *models.InlineResponse20019) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VersionEditOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*VersionEditDefault error

swagger:response versionEditDefault
*/
type VersionEditDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewVersionEditDefault creates VersionEditDefault with default headers values
func NewVersionEditDefault(code int) *VersionEditDefault {
	if code <= 0 {
		code = 500
	}

	return &VersionEditDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the version edit default response
func (o *VersionEditDefault) WithStatusCode(code int) *VersionEditDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the version edit default response
func (o *VersionEditDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the version edit default response
func (o *VersionEditDefault) WithPayload(payload *models.Error) *VersionEditDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the version edit default response
func (o *VersionEditDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VersionEditDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
