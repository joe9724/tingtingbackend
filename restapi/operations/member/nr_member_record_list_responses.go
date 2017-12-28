// Code generated by go-swagger; DO NOT EDIT.

package member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/tingtingbackend/models"
)

// MemberRecordListOKCode is the HTTP code returned for type MemberRecordListOK
const MemberRecordListOKCode int = 200

/*MemberRecordListOK 登录成功，录音列表

swagger:response memberRecordListOK
*/
type MemberRecordListOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20013 `json:"body,omitempty"`
}

// NewMemberRecordListOK creates MemberRecordListOK with default headers values
func NewMemberRecordListOK() *MemberRecordListOK {
	return &MemberRecordListOK{}
}

// WithPayload adds the payload to the member record list o k response
func (o *MemberRecordListOK) WithPayload(payload *models.InlineResponse20013) *MemberRecordListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the member record list o k response
func (o *MemberRecordListOK) SetPayload(payload *models.InlineResponse20013) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MemberRecordListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrMemberRecordListDefault error

swagger:response memberRecordListDefault
*/
type NrMemberRecordListDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrMemberRecordListDefault creates NrMemberRecordListDefault with default headers values
func NewNrMemberRecordListDefault(code int) *NrMemberRecordListDefault {
	if code <= 0 {
		code = 500
	}

	return &NrMemberRecordListDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the member record list default response
func (o *NrMemberRecordListDefault) WithStatusCode(code int) *NrMemberRecordListDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the member record list default response
func (o *NrMemberRecordListDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the member record list default response
func (o *NrMemberRecordListDefault) WithPayload(payload *models.Error) *NrMemberRecordListDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the member record list default response
func (o *NrMemberRecordListDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrMemberRecordListDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
