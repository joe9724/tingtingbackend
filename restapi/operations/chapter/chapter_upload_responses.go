// Code generated by go-swagger; DO NOT EDIT.

package chapter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/tingtingbackend/models"
)

// ChapterUploadOKCode is the HTTP code returned for type ChapterUploadOK
const ChapterUploadOKCode int = 200

/*ChapterUploadOK 上传成功，返回成功信息

swagger:response chapterUploadOK
*/
type ChapterUploadOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20015 `json:"body,omitempty"`
}

// NewChapterUploadOK creates ChapterUploadOK with default headers values
func NewChapterUploadOK() *ChapterUploadOK {
	return &ChapterUploadOK{}
}

// WithPayload adds the payload to the chapter upload o k response
func (o *ChapterUploadOK) WithPayload(payload *models.InlineResponse20015) *ChapterUploadOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the chapter upload o k response
func (o *ChapterUploadOK) SetPayload(payload *models.InlineResponse20015) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ChapterUploadOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ChapterUploadDefault error

swagger:response chapterUploadDefault
*/
type ChapterUploadDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewChapterUploadDefault creates ChapterUploadDefault with default headers values
func NewChapterUploadDefault(code int) *ChapterUploadDefault {
	if code <= 0 {
		code = 500
	}

	return &ChapterUploadDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the chapter upload default response
func (o *ChapterUploadDefault) WithStatusCode(code int) *ChapterUploadDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the chapter upload default response
func (o *ChapterUploadDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the chapter upload default response
func (o *ChapterUploadDefault) WithPayload(payload *models.Error) *ChapterUploadDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the chapter upload default response
func (o *ChapterUploadDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ChapterUploadDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
