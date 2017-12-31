// Code generated by go-swagger; DO NOT EDIT.

package relation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "tingtingbackend/models"
)

// AlbumBookListRelationOKCode is the HTTP code returned for type AlbumBookListRelationOK
const AlbumBookListRelationOKCode int = 200

/*AlbumBookListRelationOK 登录成功，返回登录信息

swagger:response albumBookListRelationOK
*/
type AlbumBookListRelationOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20026 `json:"body,omitempty"`
}

// NewAlbumBookListRelationOK creates AlbumBookListRelationOK with default headers values
func NewAlbumBookListRelationOK() *AlbumBookListRelationOK {
	return &AlbumBookListRelationOK{}
}

// WithPayload adds the payload to the album book list relation o k response
func (o *AlbumBookListRelationOK) WithPayload(payload *models.InlineResponse20026) *AlbumBookListRelationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the album book list relation o k response
func (o *AlbumBookListRelationOK) SetPayload(payload *models.InlineResponse20026) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AlbumBookListRelationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrAlbumBookListRelationDefault error

swagger:response albumBookListRelationDefault
*/
type NrAlbumBookListRelationDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrAlbumBookListRelationDefault creates NrAlbumBookListRelationDefault with default headers values
func NewNrAlbumBookListRelationDefault(code int) *NrAlbumBookListRelationDefault {
	if code <= 0 {
		code = 500
	}

	return &NrAlbumBookListRelationDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the album book list relation default response
func (o *NrAlbumBookListRelationDefault) WithStatusCode(code int) *NrAlbumBookListRelationDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the album book list relation default response
func (o *NrAlbumBookListRelationDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the album book list relation default response
func (o *NrAlbumBookListRelationDefault) WithPayload(payload *models.Error) *NrAlbumBookListRelationDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the album book list relation default response
func (o *NrAlbumBookListRelationDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrAlbumBookListRelationDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
