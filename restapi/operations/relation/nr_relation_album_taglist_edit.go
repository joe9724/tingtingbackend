// Code generated by go-swagger; DO NOT EDIT.

package relation

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
	"strings"
)

// NrRelationAlbumTaglistEditHandlerFunc turns a function with the right signature into a relation album taglist edit handler
type NrRelationAlbumTaglistEditHandlerFunc func(NrRelationAlbumTaglistEditParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrRelationAlbumTaglistEditHandlerFunc) Handle(params NrRelationAlbumTaglistEditParams) middleware.Responder {
	return fn(params)
}

// NrRelationAlbumTaglistEditHandler interface for that can handle valid relation album taglist edit params
type NrRelationAlbumTaglistEditHandler interface {
	Handle(NrRelationAlbumTaglistEditParams) middleware.Responder
}

// NewNrRelationAlbumTaglistEdit creates a new http.Handler for the relation album taglist edit operation
func NewNrRelationAlbumTaglistEdit(ctx *middleware.Context, handler NrRelationAlbumTaglistEditHandler) *NrRelationAlbumTaglistEdit {
	return &NrRelationAlbumTaglistEdit{Context: ctx, Handler: handler}
}

/*NrRelationAlbumTaglistEdit swagger:route POST /relation/album/taglist/edit Relation relationAlbumTaglistEdit

编辑专辑下的书本集合

编辑专辑下的书本集合

*/
type NrRelationAlbumTaglistEdit struct {
	Context *middleware.Context
	Handler NrRelationAlbumTaglistEditHandler
}

func (o *NrRelationAlbumTaglistEdit) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrRelationAlbumTaglistEditParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok RelationAlbumTaglistEditOK
	var response models.InlineResponse20014


	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}

	books := *(Params.Body.BookIds)

	if (*(Params.Body.ActionCode) == 0){ //添加映射
		fmt.Println("add")
		//先解析出bookis集合,样式 1,2,3,4,
		if (!strings.Contains(books,",")){
			db.Exec("insert into tag_album_relation(albumId,tagId) values(?,?)",Params.Body.AlbumID,books)
		}else{
			temp := strings.Split(books,",")
			for k:=0;k< len(temp);k++ {
				db.Exec("insert into tag_album_relation(albumId,tagId) values(?,?)",Params.Body.AlbumID,temp[k])
				fmt.Println("insert into tag_album_relation(albumId,tagId) values(?,?)",Params.Body.AlbumID,temp[k])
			}
		}
	}else{ //去除映射
	    fmt.Println("remove")
		db.Exec("delete from tag_album_relation where albumId=? and tagId=?",Params.Body.AlbumID,books)
	}

	var status models.Response
	status.UnmarshalBinary([]byte(_var.Response200(200,"ok")))
	response.Status = &status

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}