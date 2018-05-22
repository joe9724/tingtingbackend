// Code generated by go-swagger; DO NOT EDIT.

package web

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

// WebDetailHandlerFunc turns a function with the right signature into a web detail handler
type WebDetailHandlerFunc func(WebDetailParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WebDetailHandlerFunc) Handle(params WebDetailParams) middleware.Responder {
	return fn(params)
}

// WebDetailHandler interface for that can handle valid web detail params
type WebDetailHandler interface {
	Handle(WebDetailParams) middleware.Responder
}

// NewWebDetail creates a new http.Handler for the web detail operation
func NewWebDetail(ctx *middleware.Context, handler WebDetailHandler) *WebDetail {
	return &WebDetail{Context: ctx, Handler: handler}
}

/*WebDetail swagger:route GET /web/detail Web webDetail

列表

web列表

*/
type WebDetail struct {
	Context *middleware.Context
	Handler WebDetailHandler
}

func (o *WebDetail) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWebDetailParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok WebDetailOK
	var response models.InlineResponse2001513

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer db.Close()

	var web models.Web
    db.Raw("select id,content,url,title from web where id=?",Params.WebID).Find(&web)
	//status
	var status models.Response
	status.UnmarshalBinary([]byte(_var.Response200(200,"ok")))
	response.Status = &status
	response.Web = &web
	var Ds models.DashBoard
	db.Raw("select * from dashboard where id=1").Find(&Ds)
	//get hotalbums
	var hotalbums []models.HotAlbum
	db.Raw("select name,play_count as play_count from albums where status=0 order by play_count desc limit 0,3").Find(&hotalbums)
	var allcount int64
	db.Raw("select sum(play_count) as allcount from albums where status=0").Count(&allcount)
	fmt.Println("allcount is",allcount)
	for i:=0; i<len(hotalbums);i++  {
		fmt.Println("playcount is",hotalbums[i].PlayCount)
		hotalbums[i].Percent = hotalbums[i].PlayCount*100/allcount
	}
	Ds.HotAlbums = hotalbums
	response.Dashboard = Ds
	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}
