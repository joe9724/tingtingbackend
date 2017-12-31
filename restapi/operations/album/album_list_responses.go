// Code generated by go-swagger; DO NOT EDIT.

package album

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "tingtingbackend/models"
)

// AlbumListOKCode is the HTTP code returned for type AlbumListOK
const AlbumListOKCode int = 200

/*AlbumListOK 获取专辑列表

swagger:response albumListOK
*/
type AlbumListOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20020 `json:"body,omitempty"`
}

// NewAlbumListOK creates AlbumListOK with default headers values
func NewAlbumListOK() *AlbumListOK {
	return &AlbumListOK{}
}

// WithPayload adds the payload to the album list o k response
func (o *AlbumListOK) WithPayload(payload *models.InlineResponse20020) *AlbumListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the album list o k response
func (o *AlbumListOK) SetPayload(payload *models.InlineResponse20020) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AlbumListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AlbumListDefault error

swagger:response albumListDefault
*/
type AlbumListDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAlbumListDefault creates AlbumListDefault with default headers values
func NewAlbumListDefault(code int) *AlbumListDefault {
	if code <= 0 {
		code = 500
	}

	return &AlbumListDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the album list default response
func (o *AlbumListDefault) WithStatusCode(code int) *AlbumListDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the album list default response
func (o *AlbumListDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the album list default response
func (o *AlbumListDefault) WithPayload(payload *models.Error) *AlbumListDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the album list default response
func (o *AlbumListDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AlbumListDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
