// Code generated by go-swagger; DO NOT EDIT.

package album

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AlbumListHandlerFunc turns a function with the right signature into a album list handler
type AlbumListHandlerFunc func(AlbumListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AlbumListHandlerFunc) Handle(params AlbumListParams) middleware.Responder {
	return fn(params)
}

// AlbumListHandler interface for that can handle valid album list params
type AlbumListHandler interface {
	Handle(AlbumListParams) middleware.Responder
}

// NewAlbumList creates a new http.Handler for the album list operation
func NewAlbumList(ctx *middleware.Context, handler AlbumListHandler) *AlbumList {
	return &AlbumList{Context: ctx, Handler: handler}
}

/*AlbumList swagger:route GET /album/list Album albumList

获取专辑列表

获取专辑列表

*/
type AlbumList struct {
	Context *middleware.Context
	Handler AlbumListHandler
}

func (o *AlbumList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAlbumListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
