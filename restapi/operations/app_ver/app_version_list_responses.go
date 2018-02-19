// Code generated by go-swagger; DO NOT EDIT.

package app_ver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "tingtingbackend/models"
)

// AppVersionListOKCode is the HTTP code returned for type AppVersionListOK
const AppVersionListOKCode int = 200

/*AppVersionListOK 获取专辑列表

swagger:response appVersionListOK
*/
type AppVersionListOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20038 `json:"body,omitempty"`
}

// NewAppVersionListOK creates AppVersionListOK with default headers values
func NewAppVersionListOK() *AppVersionListOK {
	return &AppVersionListOK{}
}

// WithPayload adds the payload to the app version list o k response
func (o *AppVersionListOK) WithPayload(payload *models.InlineResponse20038) *AppVersionListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the app version list o k response
func (o *AppVersionListOK) SetPayload(payload *models.InlineResponse20038) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppVersionListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AppVersionListDefault error

swagger:response appVersionListDefault
*/
type AppVersionListDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAppVersionListDefault creates AppVersionListDefault with default headers values
func NewAppVersionListDefault(code int) *AppVersionListDefault {
	if code <= 0 {
		code = 500
	}

	return &AppVersionListDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the app version list default response
func (o *AppVersionListDefault) WithStatusCode(code int) *AppVersionListDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the app version list default response
func (o *AppVersionListDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the app version list default response
func (o *AppVersionListDefault) WithPayload(payload *models.Error) *AppVersionListDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the app version list default response
func (o *AppVersionListDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppVersionListDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
