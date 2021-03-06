// Code generated by go-swagger; DO NOT EDIT.

package album

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "tingtingbackend/models"
)

// AlbumDetailOKCode is the HTTP code returned for type AlbumDetailOK
const AlbumDetailOKCode int = 200

/*AlbumDetailOK 专辑详情

swagger:response albumDetailOK
*/
type AlbumDetailOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20019 `json:"body,omitempty"`
}

// NewAlbumDetailOK creates AlbumDetailOK with default headers values
func NewAlbumDetailOK() *AlbumDetailOK {
	return &AlbumDetailOK{}
}

// WithPayload adds the payload to the album detail o k response
func (o *AlbumDetailOK) WithPayload(payload *models.InlineResponse20019) *AlbumDetailOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the album detail o k response
func (o *AlbumDetailOK) SetPayload(payload *models.InlineResponse20019) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AlbumDetailOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AlbumDetailDefault error

swagger:response albumDetailDefault
*/
type AlbumDetailDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAlbumDetailDefault creates AlbumDetailDefault with default headers values
func NewAlbumDetailDefault(code int) *AlbumDetailDefault {
	if code <= 0 {
		code = 500
	}

	return &AlbumDetailDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the album detail default response
func (o *AlbumDetailDefault) WithStatusCode(code int) *AlbumDetailDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the album detail default response
func (o *AlbumDetailDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the album detail default response
func (o *AlbumDetailDefault) WithPayload(payload *models.Error) *AlbumDetailDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the album detail default response
func (o *AlbumDetailDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AlbumDetailDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
