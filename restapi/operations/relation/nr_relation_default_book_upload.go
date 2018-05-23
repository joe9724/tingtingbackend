// Code generated by go-swagger; DO NOT EDIT.

package relation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	middleware "github.com/go-openapi/runtime/middleware"
	"fmt"
	_"os"
	"tingtingbackend/models"
	"tingtingbackend/var"
	_"time"
)

// NrRelationDefaultBookUploadHandlerFunc turns a function with the right signature into a relation default book upload handler
type NrRelationDefaultBookUploadHandlerFunc func(NrRelationDefaultBookUploadParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrRelationDefaultBookUploadHandlerFunc) Handle(params NrRelationDefaultBookUploadParams) middleware.Responder {
	return fn(params)
}

// NrRelationDefaultBookUploadHandler interface for that can handle valid relation default book upload params
type NrRelationDefaultBookUploadHandler interface {
	Handle(NrRelationDefaultBookUploadParams) middleware.Responder
}

// NewNrRelationDefaultBookUpload creates a new http.Handler for the relation default book upload operation
func NewNrRelationDefaultBookUpload(ctx *middleware.Context, handler NrRelationDefaultBookUploadHandler) *NrRelationDefaultBookUpload {
	return &NrRelationDefaultBookUpload{Context: ctx, Handler: handler}
}

/*NrRelationDefaultBookUpload swagger:route POST /relation/default/book/upload Relation relationDefaultBookUpload

编辑专辑下的书本集合

编辑专辑下的书本集合

*/
type NrRelationDefaultBookUpload struct {
	Context *middleware.Context
	Handler NrRelationDefaultBookUploadHandler
}

func (o *NrRelationDefaultBookUpload) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrRelationDefaultBookUploadParams()

	//fmt.Println("Params.grade is",Params.Body.Grade)

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	//grade,bookId,startTime,status
	var ok RelationDefaultBookUploadOK
	var response models.InlineResponse2003
	var status models.Response

	var msg string
	var code int64

	msg = "ok"
	code = 200

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}


	db.Exec("insert into book_default_grade_relation(bookId,grade_range,startTime,status) values(?,?,?,?)",Params.Body.BookId,Params.Body.GradeRange,Params.Body.StartTime,Params.Body.Status)

	status.UnmarshalBinary([]byte(_var.Response200(code,msg)))
	response.Status = &status
	ok.SetPayload(&response)
	o.Context.Respond(rw, r, route.Produces, route, ok)

}
