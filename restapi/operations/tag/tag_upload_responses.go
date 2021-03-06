// Code generated by go-swagger; DO NOT EDIT.

package tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "tingtingbackend/models"
)

// TagUploadOKCode is the HTTP code returned for type TagUploadOK
const TagUploadOKCode int = 200

/*TagUploadOK 上传成功，返回成功信息

swagger:response tagUploadOK
*/
type TagUploadOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20019 `json:"body,omitempty"`
}

// NewTagUploadOK creates TagUploadOK with default headers values
func NewTagUploadOK() *TagUploadOK {
	return &TagUploadOK{}
}

// WithPayload adds the payload to the tag upload o k response
func (o *TagUploadOK) WithPayload(payload *models.InlineResponse20019) *TagUploadOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the tag upload o k response
func (o *TagUploadOK) SetPayload(payload *models.InlineResponse20019) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TagUploadOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*TagUploadDefault error

swagger:response tagUploadDefault
*/
type TagUploadDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewTagUploadDefault creates TagUploadDefault with default headers values
func NewTagUploadDefault(code int) *TagUploadDefault {
	if code <= 0 {
		code = 500
	}

	return &TagUploadDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the tag upload default response
func (o *TagUploadDefault) WithStatusCode(code int) *TagUploadDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the tag upload default response
func (o *TagUploadDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the tag upload default response
func (o *TagUploadDefault) WithPayload(payload *models.Error) *TagUploadDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the tag upload default response
func (o *TagUploadDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TagUploadDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
