// Code generated by go-swagger; DO NOT EDIT.

package msg

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "tingtingbackend/models"
)

// MsgUploadOKCode is the HTTP code returned for type MsgUploadOK
const MsgUploadOKCode int = 200

/*MsgUploadOK 上传成功，返回成功信息

swagger:response msgUploadOK
*/
type MsgUploadOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse200 `json:"body,omitempty"`
}

// NewMsgUploadOK creates MsgUploadOK with default headers values
func NewMsgUploadOK() *MsgUploadOK {
	return &MsgUploadOK{}
}

// WithPayload adds the payload to the msg upload o k response
func (o *MsgUploadOK) WithPayload(payload *models.InlineResponse200) *MsgUploadOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the msg upload o k response
func (o *MsgUploadOK) SetPayload(payload *models.InlineResponse200) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MsgUploadOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*MsgUploadDefault error

swagger:response msgUploadDefault
*/
type MsgUploadDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewMsgUploadDefault creates MsgUploadDefault with default headers values
func NewMsgUploadDefault(code int) *MsgUploadDefault {
	if code <= 0 {
		code = 500
	}

	return &MsgUploadDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the msg upload default response
func (o *MsgUploadDefault) WithStatusCode(code int) *MsgUploadDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the msg upload default response
func (o *MsgUploadDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the msg upload default response
func (o *MsgUploadDefault) WithPayload(payload *models.Error) *MsgUploadDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the msg upload default response
func (o *MsgUploadDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MsgUploadDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
