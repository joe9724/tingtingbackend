// Code generated by go-swagger; DO NOT EDIT.

package web

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	middleware "github.com/go-openapi/runtime/middleware"
	"fmt"
	"os"
	"tingtingbackend/models"
	"tingtingbackend/var"
	//"time"
	"strconv"
	"runtime"
)

// WebUploadHandlerFunc turns a function with the right signature into a web upload handler
type WebUploadHandlerFunc func(WebUploadParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WebUploadHandlerFunc) Handle(params WebUploadParams) middleware.Responder {
	return fn(params)
}

// WebUploadHandler interface for that can handle valid web upload params
type WebUploadHandler interface {
	Handle(WebUploadParams) middleware.Responder
}

// NewWebUpload creates a new http.Handler for the web upload operation
func NewWebUpload(ctx *middleware.Context, handler WebUploadHandler) *WebUpload {
	return &WebUpload{Context: ctx, Handler: handler}
}

/*WebUpload swagger:route POST /web/upload Web webUpload

添加一个web

*/
type WebUpload struct {
	Context *middleware.Context
	Handler WebUploadHandler
}

func (o *WebUpload) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWebUploadParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok WebUploadOK
	var response models.InlineResponse200
	var status models.Response
	/*var book models.Book*/
	var msg string
	var code int64

	msg = "ok"
	code = 200

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer db.Close()
	/*if(Params.IconUrl != ""){
		fmt.Println(Params.IconUrl)
		book.Icon = Params.IconUrl
	}*/
 //
	//tt:= int64(-1)
	fmt.Println("Params.id=",*(Params.ID))
	//添加正常
	db.Exec("update web set content=? where id=?",Params.Content,Params.ID)

	//生成网页

	if (runtime.GOOS == "windows") {
		f, err3 := os.Create("./"+strconv.FormatInt(*(Params.ID),10)+".html") //创建文件
		check(err3)
		defer f.Close()
		f.Sync()
	} else {
		//err1 := ioutil.WriteFile("/root/go/src/resource/image/icon/"+filename+".jpg", file, 0644)
		f, err3 := os.Create("/root/go/src/resource/html/"+strconv.FormatInt(*(Params.ID),10)+".html") //创建文件
		check(err3)
		defer f.Close()
		f.Sync()
	}


	//n2, err3 := f.Write(d1) //写入文件(字节数组)
	//check(err3)
	//fmt.Printf("写入 %d 个字节n", n2)
	/*n3, err3 := f.WriteString(*(Params.Content)) //写入文件(字节数组)
	fmt.Printf("写入 %d 个字节n", n3)*/


	status.UnmarshalBinary([]byte(_var.Response200(code,msg)))
	response.Status = &status
	ok.SetPayload(&response)
	o.Context.Respond(rw, r, route.Produces, route, ok)

}
func check(e error) {
	if e != nil {
		panic(e)
	}
}