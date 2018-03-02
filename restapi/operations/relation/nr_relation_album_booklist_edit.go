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
	"strconv"
)

// NrRelationAlbumBooklistEditHandlerFunc turns a function with the right signature into a relation album booklist edit handler
type NrRelationAlbumBooklistEditHandlerFunc func(NrRelationAlbumBooklistEditParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrRelationAlbumBooklistEditHandlerFunc) Handle(params NrRelationAlbumBooklistEditParams) middleware.Responder {
	return fn(params)
}

// NrRelationAlbumBooklistEditHandler interface for that can handle valid relation album booklist edit params
type NrRelationAlbumBooklistEditHandler interface {
	Handle(NrRelationAlbumBooklistEditParams) middleware.Responder
}

// NewNrRelationAlbumBooklistEdit creates a new http.Handler for the relation album booklist edit operation
func NewNrRelationAlbumBooklistEdit(ctx *middleware.Context, handler NrRelationAlbumBooklistEditHandler) *NrRelationAlbumBooklistEdit {
	return &NrRelationAlbumBooklistEdit{Context: ctx, Handler: handler}
}

/*NrRelationAlbumBooklistEdit swagger:route POST /relation/album/booklist/edit Relation relationAlbumBooklistEdit

编辑专辑下的书本集合

编辑专辑下的书本集合

*/
type NrRelationAlbumBooklistEdit struct {
	Context *middleware.Context
	Handler NrRelationAlbumBooklistEditHandler
}

func (o *NrRelationAlbumBooklistEdit) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("edit_book_chapter_relation")
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrRelationAlbumBooklistEditParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok RelationAlbumBooklistEditOK
	var response models.InlineResponse20018


	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}

	books := *(Params.Body.BookIds)

	fmt.Println("Params.body.actionCode is",*(Params.Body.Action))
	
	if (*(Params.Body.Action) == 0){ //添加映射
		//先解析出bookis集合,样式 1,2,3,4,
		if (!strings.Contains(books,",")){
			fmt.Println("1")
			db.Exec("insert into album_book_relation(albumId,bookId) values(?,?)",Params.Body.AlbumID,books)
			db.Exec("update albums set books_number=books_number+1 where id=?",Params.Body.AlbumID)
		}else{
			temp := strings.Split(books,",")
			for k:=0;k< len(temp);k++ {
				db.Exec("insert into album_book_relation(albumId,bookId) values(?,?)",Params.Body.AlbumID,temp[k])
				fmt.Println("insert into album_book_relation(albumId,bookId) values(?,?)",Params.Body.AlbumID,temp[k])
			}
			db.Exec("update albums set books_number=albums_number+"+strconv.Itoa(len(temp))+" where id="+strconv.FormatInt(*(Params.Body.AlbumID),10))
			fmt.Println("2")
		}
	}else{ //去除映射
		db.Exec("delete from album_book_relation where albumId=? and bookId=?",Params.Body.AlbumID,books)
		db.Exec("update albums set books_number=books_number-1 where id=?",Params.Body.AlbumID)
		fmt.Println("3")
	}
	
	

	//query
	//db.Where(map[string]interface{}{"status":0}).Find(&albumList).Limit(*(Params.PageSize)).Offset(*(Params.PageIndex)*(*(Params.PageSize)))
	//data
	//response.AlbumList = albumList
	//fmt.Println("haspushed is",albumList[0].HasPushed)
	//status
	var status models.Response
	status.UnmarshalBinary([]byte(_var.Response200(200,"ok")))
	response.Status = &status

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}
