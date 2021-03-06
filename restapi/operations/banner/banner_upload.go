// Code generated by go-swagger; DO NOT EDIT.

package banner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	middleware "github.com/go-openapi/runtime/middleware"
	"fmt"
	_"os"
	"tingtingbackend/models"
	"tingtingbackend/var"
)

// BannerUploadHandlerFunc turns a function with the right signature into a banner upload handler
type BannerUploadHandlerFunc func(BannerUploadParams) middleware.Responder

// Handle executing the request and returning a response
func (fn BannerUploadHandlerFunc) Handle(params BannerUploadParams) middleware.Responder {
	return fn(params)
}

// BannerUploadHandler interface for that can handle valid banner upload params
type BannerUploadHandler interface {
	Handle(BannerUploadParams) middleware.Responder
}

// NewBannerUpload creates a new http.Handler for the banner upload operation
func NewBannerUpload(ctx *middleware.Context, handler BannerUploadHandler) *BannerUpload {
	return &BannerUpload{Context: ctx, Handler: handler}
}

/*BannerUpload swagger:route POST /banner/upload Banner bannerUpload

添加一个banner

*/
type BannerUpload struct {
	Context *middleware.Context
	Handler BannerUploadHandler
}

func (o *BannerUpload) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewBannerUploadParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok BannerUploadOK
	var response models.InlineResponse2002
	var status models.Response
	var banner models.Banner
	var msg string
	var code int64

	msg = "ok"
	code = 200

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	if(Params.CoverUrl != ""){
		fmt.Println(Params.CoverUrl)
		banner.Cover = &(Params.CoverUrl)
	}

	//tt:= int64(-1)
	fmt.Println("Params.BookId=",*(Params.BannerID))
	if(*(Params.BannerID) == -1){ //新建
		banner.Type = Params.Type
		banner.Name = Params.Name
		banner.Jumpurl = Params.WebURL
		banner.Status = *(Params.Status)
		banner.Jumpid = Params.TargetID
		if(Params.CoverUrl != ""){
			banner.Cover = &(Params.CoverUrl)
		}
		db.Exec("insert into banners(name,status,type,jumpid,jumpurl,cover) values(?,?,?,?,?,?)",Params.Name,*(Params.Status),*(Params.Type),*(Params.TargetID),Params.WebURL,Params.CoverUrl)
		//book.User_id = *(Params.MemberID)
		//db.Table("banners").Create(&banner)
	}else{ //更新
		//fmt.Println("edit")
		//db.Table("sub_book_items").Where("id=?",*(Params.BookId)).Last(&book)
		if(Params.CoverUrl != ""){
			db.Exec("update banners set name=?,status=?,type=?,cover=?,jumpid=?,jumpurl=? where id=?",Params.Name,*(Params.Status),*(Params.Type),Params.CoverUrl,*(Params.TargetID),Params.WebURL,*(Params.BannerID))
		}else{
			db.Exec("update banners set name=?,status=?,type=?,jumpid=?,jumpurl=? where id=?",Params.Name,*(Params.Status),*(Params.Type),*(Params.TargetID),Params.WebURL,*(Params.BannerID))
		}

	}

	status.UnmarshalBinary([]byte(_var.Response200(code,msg)))
	response.Status = &status
	ok.SetPayload(&response)
	o.Context.Respond(rw, r, route.Produces, route, ok)

}
