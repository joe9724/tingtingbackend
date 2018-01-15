// Code generated by go-swagger; DO NOT EDIT.

package recharge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "backendtest/models"
)

// RechargeDetailOKCode is the HTTP code returned for type RechargeDetailOK
const RechargeDetailOKCode int = 200

/*RechargeDetailOK 获取反馈列表

swagger:response rechargeDetailOK
*/
type RechargeDetailOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse2004 `json:"body,omitempty"`
}

// NewRechargeDetailOK creates RechargeDetailOK with default headers values
func NewRechargeDetailOK() *RechargeDetailOK {
	return &RechargeDetailOK{}
}

// WithPayload adds the payload to the recharge detail o k response
func (o *RechargeDetailOK) WithPayload(payload *models.InlineResponse2004) *RechargeDetailOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the recharge detail o k response
func (o *RechargeDetailOK) SetPayload(payload *models.InlineResponse2004) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RechargeDetailOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrRechargeDetailDefault error

swagger:response rechargeDetailDefault
*/
type NrRechargeDetailDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrRechargeDetailDefault creates NrRechargeDetailDefault with default headers values
func NewNrRechargeDetailDefault(code int) *NrRechargeDetailDefault {
	if code <= 0 {
		code = 500
	}

	return &NrRechargeDetailDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the recharge detail default response
func (o *NrRechargeDetailDefault) WithStatusCode(code int) *NrRechargeDetailDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the recharge detail default response
func (o *NrRechargeDetailDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the recharge detail default response
func (o *NrRechargeDetailDefault) WithPayload(payload *models.Error) *NrRechargeDetailDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the recharge detail default response
func (o *NrRechargeDetailDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrRechargeDetailDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}