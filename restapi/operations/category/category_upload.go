// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	_"io/ioutil"
	middleware "github.com/go-openapi/runtime/middleware"
	"fmt"
	_"os"
	_"runtime"
	_"strings"
	"tingtingbackend/models"
	"tingtingbackend/var"
)

// CategoryUploadHandlerFunc turns a function with the right signature into a category upload handler
type CategoryUploadHandlerFunc func(CategoryUploadParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CategoryUploadHandlerFunc) Handle(params CategoryUploadParams) middleware.Responder {
	return fn(params)
}

// CategoryUploadHandler interface for that can handle valid category upload params
type CategoryUploadHandler interface {
	Handle(CategoryUploadParams) middleware.Responder
}

// NewCategoryUpload creates a new http.Handler for the category upload operation
func NewCategoryUpload(ctx *middleware.Context, handler CategoryUploadHandler) *CategoryUpload {
	return &CategoryUpload{Context: ctx, Handler: handler}
}

/*CategoryUpload swagger:route POST /category/upload Category categoryUpload

添加一个大类

*/
type CategoryUpload struct {
	Context *middleware.Context
	Handler CategoryUploadHandler
}

func (o *CategoryUpload) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCategoryUploadParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok CategoryUploadOK
	var response models.InlineResponse20016
	var status models.Response
	var category models.Category
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
		category.Icon = &(Params.IconUrl)
	}

	//tt:= int64(-1)
	fmt.Println("Params.CategoryId=",*(Params.CategoryId))
	if(*(Params.CategoryId) == -1){ //新建
	    fmt.Println("new")
		fmt.Println("Params.Summary is",Params.Summary)
		category.Summary = Params.Summary
		category.Name = &(Params.Title)
		t := int64(-1)
		category.Category_Id = &t
		category.Status = Params.Status
		if(Params.IconUrl != ""){
			category.Icon = &(Params.IconUrl)
		}
		//album.User_id = *(Params.MemberID)
		db.Table("sub_category_items").Create(&category)
	}else{ //更新
	    //fmt.Println("edit")
		//db.Table("sub_category_items").Where("id=?",*(Params.CategoryId)).Last(&category)
		if(Params.IconUrl != ""){
			fmt.Println("0",Params.Title)
			fmt.Println("1",Params.IconUrl)
			fmt.Println("2",*(Params.Summary))
			fmt.Println("3",*(Params.CategoryId))
			name := Params.Title
			summary := Params.Summary
			iconurl := Params.IconUrl
			categoryid := Params.CategoryId

			db.Exec("update sub_category_items set name=?,status=?,summary=?,icon=? where id=?",name,0,summary,iconurl,categoryid)
		}else{
			fmt.Println("4",Params.IconUrl,*(Params.Summary))
			summary := *(Params.Summary)
			db.Exec("update sub_category_items set name=?,status=?,summary=? where id=?",Params.Title,0,summary,*(Params.CategoryId))
		}

	}

	status.UnmarshalBinary([]byte(_var.Response200(code,msg)))
	response.Return = &status
	ok.SetPayload(&response)
	o.Context.Respond(rw, r, route.Produces, route, ok)

}
