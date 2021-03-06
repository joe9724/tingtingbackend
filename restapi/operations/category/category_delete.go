// Code generated by go-swagger; DO NOT EDIT.

package category

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

// CategoryDeleteHandlerFunc turns a function with the right signature into a category delete handler
type CategoryDeleteHandlerFunc func(CategoryDeleteParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CategoryDeleteHandlerFunc) Handle(params CategoryDeleteParams) middleware.Responder {
	return fn(params)
}

// CategoryDeleteHandler interface for that can handle valid category delete params
type CategoryDeleteHandler interface {
	Handle(CategoryDeleteParams) middleware.Responder
}

// NewCategoryDelete creates a new http.Handler for the category delete operation
func NewCategoryDelete(ctx *middleware.Context, handler CategoryDeleteHandler) *CategoryDelete {
	return &CategoryDelete{Context: ctx, Handler: handler}
}

/*CategoryDelete swagger:route GET /category/delete Category categoryDelete

删除分类

删除分类

*/
type CategoryDelete struct {
	Context *middleware.Context
	Handler CategoryDeleteHandler
}

func (o *CategoryDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCategoryDeleteParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok CategoryDeleteOK
	var response models.InlineResponse20018

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	//query
	//sql := "update sub_category_items set status=1 where id=?",*(Params.CategoryID)
	fmt.Println("sql is ","update sub_category_items set status=1 where id=?",*(Params.CategoryID))
	db.Exec("update sub_category_items set status=1 where id=?",*(Params.CategoryID))
	//var category models.Category
	//db.Table("sub_category_items").Where("id=?",Params.CategoryID).First(&category)
	//data
	//response.Data = &category

	//status
	var status models.Response
	status.UnmarshalBinary([]byte(_var.Response200(200,"ok")))
	response.Status = &status

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}
