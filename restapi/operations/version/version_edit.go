// Code generated by go-swagger; DO NOT EDIT.

package version

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// VersionEditHandlerFunc turns a function with the right signature into a version edit handler
type VersionEditHandlerFunc func(VersionEditParams) middleware.Responder

// Handle executing the request and returning a response
func (fn VersionEditHandlerFunc) Handle(params VersionEditParams) middleware.Responder {
	return fn(params)
}

// VersionEditHandler interface for that can handle valid version edit params
type VersionEditHandler interface {
	Handle(VersionEditParams) middleware.Responder
}

// NewVersionEdit creates a new http.Handler for the version edit operation
func NewVersionEdit(ctx *middleware.Context, handler VersionEditHandler) *VersionEdit {
	return &VersionEdit{Context: ctx, Handler: handler}
}

/*VersionEdit swagger:route POST /version/edit Version versionEdit

编辑一个版本

*/
type VersionEdit struct {
	Context *middleware.Context
	Handler VersionEditHandler
}

func (o *VersionEdit) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewVersionEditParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
