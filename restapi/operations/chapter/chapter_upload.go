// Code generated by go-swagger; DO NOT EDIT.

package chapter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	middleware "github.com/go-openapi/runtime/middleware"
	"fmt"
	_"os"
	"tingtingbackend/models"
	"tingtingbackend/var"
	"time"
	"strings"
	"strconv"
)

// ChapterUploadHandlerFunc turns a function with the right signature into a chapter upload handler
type ChapterUploadHandlerFunc func(ChapterUploadParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ChapterUploadHandlerFunc) Handle(params ChapterUploadParams) middleware.Responder {
	return fn(params)
}

// ChapterUploadHandler interface for that can handle valid chapter upload params
type ChapterUploadHandler interface {
	Handle(ChapterUploadParams) middleware.Responder
}

// NewChapterUpload creates a new http.Handler for the chapter upload operation
func NewChapterUpload(ctx *middleware.Context, handler ChapterUploadHandler) *ChapterUpload {
	return &ChapterUpload{Context: ctx, Handler: handler}
}

/*ChapterUpload swagger:route POST /chapter/upload Chapter chapterUpload

添加一个章节

*/
type ChapterUpload struct {
	Context *middleware.Context
	Handler ChapterUploadHandler
}

func (o *ChapterUpload) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewChapterUploadParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok ChapterUploadOK
	var response models.InlineResponse20016
	var status models.Response
	var chapter models.Chapter
	var msg string
	var code int64

	msg = "ok"
	code = 200

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	if(Params.IconUrl != ""){
		fmt.Println(Params.IconUrl)
		chapter.Icon = &(Params.IconUrl)
	}

	//tt:= int64(-1)
	fmt.Println("Params.ChapterId=",*(Params.ChapterId))
	if(*(Params.ChapterId) == -1){ //新建
		fmt.Println("new")
		fmt.Println("Params.Summary is",Params.Summary)
		chapter.Summary = Params.Summary
		chapter.Name = &(Params.Title)
		chapter.Time = time.Now().UnixNano() / 1000000000
		chapter.URL = &(Params.Url)
		var pc  int64
		pc = 1
		chapter.PlayCount = &pc
		// 00:00:00 -> duration
		//var duration int64
		var tempduration string
		var s []string
		tempduration = Params.Duration
		fmt.Println("duration is",Params.Duration)
		s = strings.Split(tempduration,":")
		hour,_ := strconv.ParseInt(s[0],10,64)
		min,_ := strconv.ParseInt(s[1],10,64)
		sec,_ := strconv.ParseInt(s[2],10,64)
		var du int64
		du = hour*3600+min*60+sec*1
		chapter.Duration = 100000

		//chapter.AuthorName = Params.AuthorName
		//fmt.Println("author is",Params.AuthorName)
		//t := int64(-1)
		//chapter.chapter_Id = &t
		chapter.Status = *(Params.Status)
		if(Params.IconUrl != ""){
			chapter.Icon = &(Params.IconUrl)
		}
		//chapter.User_id = *(Params.MemberID)
		db.Table("chapters").Create(&chapter)
	}else{ //更新
		//fmt.Println("edit")
		//db.Table("sub_chapter_items").Where("id=?",*(Params.ChapterId)).Last(&chapter)
		if(Params.IconUrl != ""){
			fmt.Println("1",Params.IconUrl)
			db.Exec("update chapters set name=?,status=?,summary=?,icon=?,url=? where id=?",Params.Title,0,*(Params.Summary),Params.IconUrl,Params.Url,&(Params.ChapterId))
		}else{
			fmt.Println("2",Params.IconUrl,*(Params.Summary))
			summary := *(Params.Summary)
			db.Exec("update chapters set name=?,status=?,summary=?,url=? where id=?",Params.Title,0,summary,Params.Url,&(Params.ChapterId))
		}

	}

	status.UnmarshalBinary([]byte(_var.Response200(code,msg)))
	response.Return = &status
	ok.SetPayload(&response)
	o.Context.Respond(rw, r, route.Produces, route, ok)

}
