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
	var count int64

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer db.Close()
	//query
	//db.Where(map[string]interface{}{"status":0}).Limit(*(Params.PageSize)).Offset(*(Params.PageIndex)*(*(Params.PageSize))).Find(&albumList)
	//db.Table("albums").Where(map[string]interface{}{"status":0}).Count(&count)
	//判断是否是用户收藏的专辑列表请求还是其他
	if Params.Memberid !=nil{
		fmt.Println("get fav albumlist",*(Params.Memberid))
		db.Raw("select albums.id, albums.name,albums.icon,fav_album.time from albums left join fav_album on albums.id = fav_album.album_id where fav_album.member_id=?",*(Params.Memberid)).Limit(*(Params.PageSize)).Offset(*(Params.PageIndex) * (*(Params.PageSize))).Find(&albumList)
		db.Raw("select albums.id, albums.name,albums.icon,fav_album.timefrom albums left join fav_album on albums.id = fav_album.album_id where fav_album.member_id=?",*(Params.Memberid)).Count(&count)
		//db.Table("albums").Select("albums.id, albums.name").Joins("left join fav_album on albums.id = fav_album.album_id").Where("fav_album.member_id =?", *(Params.Memberid)).Where("albums.status=?", 0).Where("fav_album.status =?", 0).Count(&count)
		//db.Table("albums").Select("albums.id, albums.name").Joins("left join fav_album on albums.id = fav_album.album_id").Where("fav_album.member_id =?", *(Params.Memberid)).Where("albums.status=?", 0).Where("fav_album.status =?", 0).Limit(*(Params.PageSize)).Offset(*(Params.PageIndex) * (*(Params.PageSize))).Find(&albumList)
	}else {
		if Params.Keyword != nil && Params.Categoryid != nil {
			if (*Params.Keyword == " ") {
				db.Raw("select id,name  FROM albums where  status=0 and id not in (select albumId from category_album_relation  where status=0 and albumId = ? )", *(Params.Categoryid)).Limit(*(Params.PageSize)).Offset(*(Params.PageIndex) * (*(Params.PageSize))).Find(&albumList)
				db.Raw("select id,name  FROM albums where status=0 and  id not in (select albumId from category_album_relation  where status=0 and albumId = ? )", *(Params.Categoryid)).Count(&count)
			} else {
				db.Raw("select id,name  FROM albums where status=0 and name like '%" + *(Params.Keyword)+"%' and id not in (select albumId from category_album_relation  where status=0 and albumId = ? )", *(Params.Categoryid)).Limit(*(Params.PageSize)).Offset(*(Params.PageIndex) * (*(Params.PageSize))).Find(&albumList)
				db.Raw("select id,name  FROM albums where status=0 and name like '%" + *(Params.Keyword)+"%' and id not in (select albumId from category_album_relation  where status=0 and albumId = ? )", *(Params.Categoryid)).Count(&count)
			}
			//db.Table("albums").Select("albums.id, albums.name").Joins("left join category_album_relation on albums.id = category_album_relation.albumId").Where("albums.id is null").Find(&bookList)
			fmt.Println("1")
			//db.Where(map[string]interface{}{"status":0}).Where("name like ?","%"+*(Params.Keyword)+"%").Not("id",).Find(&bookList).Offset(*(Params.PageIndex)*(*(Params.PageSize)))
		} else {
			if Params.Categoryid != nil {
				fmt.Println("2")
				db.Table("albums").Select("albums.id, albums.name").Joins("left join category_album_relation on albums.id = category_album_relation.albumId").Where("category_album_relation.categoryId =?", *Params.Categoryid).Where("albums.status=?", 0).Where("category_album_relation.status =?", 0).Count(&count)
				db.Table("albums").Select("albums.id, albums.name").Joins("left join category_album_relation on albums.id = category_album_relation.albumId").Where("category_album_relation.categoryId =?", *Params.Categoryid).Where("albums.status=?", 0).Where("category_album_relation.status =?", 0).Limit(*(Params.PageSize)).Offset(*(Params.PageIndex) * (*(Params.PageSize))).Find(&albumList)
			} else {
				fmt.Println("3")
				db.Table("albums").Where(map[string]interface{}{"status": 0}).Count(&count)
				db.Table("albums").Where(map[string]interface{}{"status": 0}).Limit(*(Params.PageSize)).Offset(*(Params.PageIndex) * (*(Params.PageSize))).Find(&albumList)
			}

		}
	}
	//data
	response.AlbumList = albumList
	fmt.Println("size is",len(albumList))
    //fmt.Println("haspushed is",albumList[0].HasPushed)
	//status
	var status models.Response
	status.UnmarshalBinary([]byte(_var.Response200(200,"ok")))
	response.Status = &status
	response.Status.TotalCount = count

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}
