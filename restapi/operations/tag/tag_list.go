// Code generated by go-swagger; DO NOT EDIT.

package tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/go-openapi/runtime/middleware"
	"tingtingbackend/models"
	"fmt"
	"tingtingbackend/var"
)

// TagListHandlerFunc turns a function with the right signature into a tag list handler
type TagListHandlerFunc func(TagListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn TagListHandlerFunc) Handle(params TagListParams) middleware.Responder {
	return fn(params)
}

// TagListHandler interface for that can handle valid tag list params
type TagListHandler interface {
	Handle(TagListParams) middleware.Responder
}

// NewTagList creates a new http.Handler for the tag list operation
func NewTagList(ctx *middleware.Context, handler TagListHandler) *TagList {
	return &TagList{Context: ctx, Handler: handler}
}

/*TagList swagger:route GET /tag/list Tag tagList

获取专辑列表

获取专辑列表

*/
type TagList struct {
	Context *middleware.Context
	Handler TagListHandler
}

func (o *TagList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewTagListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok TagListOK
	var response models.InlineResponse20034
	var albumList models.InlineResponse20034AlbumList
	var count int64

	db, err := _var.OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	}
	//query
	if Params.Keyword != nil && Params.AlbumId != nil {
		if(*Params.Keyword == " ") {
			fmt.Println("1")
			db.Raw("select id,name  FROM tags where  id not in (select tagId from tag_album_relation  where albumId = ? )", *(Params.AlbumId)).Limit(*(Params.PageSize)).Offset(*(Params.PageIndex)*(*(Params.PageSize))).Find(&albumList)
			db.Raw("select id,name  FROM tags where  id not in (select tagId from tag_album_relation  where albumId = ? )", *(Params.AlbumId)).Count(&count)
		}else{
			fmt.Println("2")
			db.Raw("select id,name  FROM tags where name like '%" + *(Params.Keyword)+"%' and id not in (select tagId from tag_album_relation  where albumId = ? )", *(Params.AlbumId)).Limit(*(Params.PageSize)).Offset(*(Params.PageIndex)*(*(Params.PageSize))).Find(&albumList)
			db.Raw("select id,name  FROM tags where name like '%" + *(Params.Keyword)+"%' and id not in (select tagId from tag_album_relation  where albumId = ? )", *(Params.AlbumId)).Count(&count)
		}
	} else {
		if (Params.AlbumId != nil) { //获取专辑对应的id
			fmt.Println("3")
			db.Table("tags").Select("tags.id, tags.name").Joins("left join tag_album_relation on tags.id = tag_album_relation.tagId").Where("tag_album_relation.albumId =?", *Params.AlbumId).Limit(*(Params.PageSize)).Offset(*(Params.PageIndex) * (*(Params.PageSize))).Find(&albumList)
			db.Table("tags").Select("tags.id, tags.name").Joins("left join tag_album_relation on tags.id = tag_album_relation.tagId").Where("tag_album_relation.albumId =?", *Params.AlbumId).Count(&count)
		} else {
			fmt.Println("4")
			db.Where(map[string]interface{}{"status": 0}).Limit(*(Params.PageSize)).Offset(*(Params.PageIndex) * (*(Params.PageSize))).Find(&albumList)
			db.Table("albums").Where(map[string]interface{}{"status": 0}).Count(&count)
		}
	}
	//data
	response.AlbumList = albumList
	//fmt.Println("size is",len(albumList))
	//fmt.Println("haspushed is",albumList[0].HasPushed)
	//status
	var status models.Response
	status.UnmarshalBinary([]byte(_var.Response200(200, "ok")))
	response.Status = &status
	response.Status.TotalCount = count

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}
