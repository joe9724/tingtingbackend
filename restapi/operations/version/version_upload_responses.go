// Code generated by go-swagger; DO NOT EDIT.

package version

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "tingtingbackend/models"
)

// VersionUploadOKCode is the HTTP code returned for type VersionUploadOK
const VersionUploadOKCode int = 200

/*VersionUploadOK 成功信息

swagger:response versionUploadOK
*/
type VersionUploadOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20019 `json:"body,omitempty"`
}

// NewVersionUploadOK creates VersionUploadOK with default headers values
func NewVersionUploadOK() *VersionUploadOK {
	return &VersionUploadOK{}
}

// WithPayload adds the payload to the version upload o k response
func (o *VersionUploadOK) WithPayload(payload *models.InlineResponse20019) *VersionUploadOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the version upload o k response
func (o *VersionUploadOK) SetPayload(payload *models.InlineResponse20019) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VersionUploadOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*VersionUploadDefault error

swagger:response versionUploadDefault
*/
type VersionUploadDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewVersionUploadDefault creates VersionUploadDefault with default headers values
func NewVersionUploadDefault(code int) *VersionUploadDefault {
	if code <= 0 {
		code = 500
	}

	return &VersionUploadDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the version upload default response
func (o *VersionUploadDefault) WithStatusCode(code int) *VersionUploadDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the version upload default response
func (o *VersionUploadDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the version upload default response
func (o *VersionUploadDefault) WithPayload(payload *models.Error) *VersionUploadDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the version upload default response
func (o *VersionUploadDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VersionUploadDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}