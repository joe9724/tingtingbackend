// Code generated by go-swagger; DO NOT EDIT.

package recharge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "backendtest/models"
)

// RechargeListOKCode is the HTTP code returned for type RechargeListOK
const RechargeListOKCode int = 200

/*RechargeListOK 获取反馈列表

swagger:response rechargeListOK
*/
type RechargeListOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse2003 `json:"body,omitempty"`
}

// NewRechargeListOK creates RechargeListOK with default headers values
func NewRechargeListOK() *RechargeListOK {
	return &RechargeListOK{}
}

// WithPayload adds the payload to the recharge list o k response
func (o *RechargeListOK) WithPayload(payload *models.InlineResponse2003) *RechargeListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the recharge list o k response
func (o *RechargeListOK) SetPayload(payload *models.InlineResponse2003) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RechargeListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*RechargeListDefault error

swagger:response rechargeListDefault
*/
type RechargeListDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRechargeListDefault creates RechargeListDefault with default headers values
func NewRechargeListDefault(code int) *RechargeListDefault {
	if code <= 0 {
		code = 500
	}

	return &RechargeListDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the recharge list default response
func (o *RechargeListDefault) WithStatusCode(code int) *RechargeListDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the recharge list default response
func (o *RechargeListDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the recharge list default response
func (o *RechargeListDefault) WithPayload(payload *models.Error) *RechargeListDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the recharge list default response
func (o *RechargeListDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RechargeListDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}