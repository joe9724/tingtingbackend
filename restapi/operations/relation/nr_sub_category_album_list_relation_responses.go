// Code generated by go-swagger; DO NOT EDIT.

package relation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "tingtingbackend/models"
)

// SubCategoryAlbumListRelationOKCode is the HTTP code returned for type SubCategoryAlbumListRelationOK
const SubCategoryAlbumListRelationOKCode int = 200

/*SubCategoryAlbumListRelationOK 登录成功，返回登录信息

swagger:response subCategoryAlbumListRelationOK
*/
type SubCategoryAlbumListRelationOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20027 `json:"body,omitempty"`
}

// NewSubCategoryAlbumListRelationOK creates SubCategoryAlbumListRelationOK with default headers values
func NewSubCategoryAlbumListRelationOK() *SubCategoryAlbumListRelationOK {
	return &SubCategoryAlbumListRelationOK{}
}

// WithPayload adds the payload to the sub category album list relation o k response
func (o *SubCategoryAlbumListRelationOK) WithPayload(payload *models.InlineResponse20027) *SubCategoryAlbumListRelationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the sub category album list relation o k response
func (o *SubCategoryAlbumListRelationOK) SetPayload(payload *models.InlineResponse20027) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SubCategoryAlbumListRelationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrSubCategoryAlbumListRelationDefault error

swagger:response subCategoryAlbumListRelationDefault
*/
type NrSubCategoryAlbumListRelationDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrSubCategoryAlbumListRelationDefault creates NrSubCategoryAlbumListRelationDefault with default headers values
func NewNrSubCategoryAlbumListRelationDefault(code int) *NrSubCategoryAlbumListRelationDefault {
	if code <= 0 {
		code = 500
	}

	return &NrSubCategoryAlbumListRelationDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the sub category album list relation default response
func (o *NrSubCategoryAlbumListRelationDefault) WithStatusCode(code int) *NrSubCategoryAlbumListRelationDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the sub category album list relation default response
func (o *NrSubCategoryAlbumListRelationDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the sub category album list relation default response
func (o *NrSubCategoryAlbumListRelationDefault) WithPayload(payload *models.Error) *NrSubCategoryAlbumListRelationDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the sub category album list relation default response
func (o *NrSubCategoryAlbumListRelationDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrSubCategoryAlbumListRelationDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
