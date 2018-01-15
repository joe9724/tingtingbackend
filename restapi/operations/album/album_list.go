// Code generated by go-swagger; DO NOT EDIT.

package album

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	_"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	middleware "github.com/go-openapi/runtime/middleware"
	"tingtingbackend/models"
	"fmt"
	"tingtingbackend/var"
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

	var ok AlbumListOK
	var response models.InlineResponse20020
	var albumList models.InlineResponse20020AlbumList

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	//query
	db.Where(map[string]interface{}{"status":0}).Find(&albumList).Limit(*(Params.PageSize)).Offset(*(Params.PageIndex)*(*(Params.PageSize)))
	//data
	response.AlbumList = albumList
    fmt.Println("haspushed is",albumList[0].HasPushed)
	//status
	var status models.Response
	status.UnmarshalBinary([]byte(_var.Response200(200,"ok")))
	response.Status = &status

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}
