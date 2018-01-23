// Code generated by go-swagger; DO NOT EDIT.

package album

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	"io/ioutil"
	middleware "github.com/go-openapi/runtime/middleware"
	"fmt"
	_"os"
	"runtime"
	"time"
	"strings"
	"tingtingbackend/models"
	"tingtingbackend/var"
	"strconv"
)

// AlbumUploadHandlerFunc turns a function with the right signature into a album upload handler
type AlbumUploadHandlerFunc func(AlbumUploadParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AlbumUploadHandlerFunc) Handle(params AlbumUploadParams) middleware.Responder {
	return fn(params)
}

// AlbumUploadHandler interface for that can handle valid album upload params
type AlbumUploadHandler interface {
	Handle(AlbumUploadParams) middleware.Responder
}

// NewAlbumUpload creates a new http.Handler for the album upload operation
func NewAlbumUpload(ctx *middleware.Context, handler AlbumUploadHandler) *AlbumUpload {
	return &AlbumUpload{Context: ctx, Handler: handler}
}

/*AlbumUpload swagger:route POST /album/upload Album albumUpload

添加一个专辑

*/
type AlbumUpload struct {
	Context *middleware.Context
	Handler AlbumUploadHandler
}

func (o *AlbumUpload) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAlbumUploadParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok AlbumUploadOK
	var response models.InlineResponse20016
	var status models.Response
	var album models.Album
	var msg string
	var code int64

	var filename string
	filename = strconv.FormatInt((time.Now().Unix()),10)

	fmt.Println("filename is",filename)

	if(Params.IconUrl != ""){
		fmt.Println(Params.IconUrl)
		album.Icon = Params.IconUrl
	}

	//如果有icon
	if (Params.Icon!=nil) {
		icon, err := ioutil.ReadAll(Params.Icon)
		if err != nil {
			fmt.Println("err upload:", err.Error())
		}
		fmt.Println(len(icon))
		// Always returns a valid content-type and "application/octet-stream" if no others seemed to match.
		contentType := http.DetectContentType(icon)
		fmt.Println("contentType is", contentType)

		//save
		var lower string
		lower = strings.ToLower(contentType)
		if(strings.Contains(lower,"jp")||(strings.Contains(lower,"pn"))) {
			if (runtime.GOOS == "windows") {
				err1 := ioutil.WriteFile(filename+".jpg", icon, 0644)
				if err1 != nil {
					fmt.Println(err1.Error())
				}
			} else {
				err1 := ioutil.WriteFile("/root/go/src/resource/image/icon/"+filename+".jpg", icon, 0644)
				if err1 != nil {
					fmt.Println(err1.Error())
				}
			}
			album.Icon = "http://tingting-resource.bitekun.xin/resource/image/icon/"+filename+".jpg"
			code = 200
			msg = "ok"
		}else{
			code = 401
			msg = "image format need jpg or png"
		}
	}

	//如果有cover
	if (Params.Cover!=nil) {
		cover, err := ioutil.ReadAll(Params.Cover)
		if err != nil {
			fmt.Println("err upload:", err.Error())
		}
		//fmt.Println(len(icon))
		// Always returns a valid content-type and "application/octet-stream" if no others seemed to match.
		contentType := http.DetectContentType(cover)
		//fmt.Println("contentType is", contentType)

		//save
		var lower string
		lower = strings.ToLower(contentType)
		if(strings.Contains(lower,"jp")||(strings.Contains(lower,"pn"))) {
			if(runtime.GOOS == "windows") {
				err1 := ioutil.WriteFile(filename+".jpg", cover, 0644)
				if err1 != nil {
					fmt.Println(err1.Error())
				}
			}else{
				err1 := ioutil.WriteFile("/root/go/src/resource/image/cover/"+filename+".jpg", cover, 0644)
				if err1 != nil {
					fmt.Println(err1.Error())
				}
			}
			code = 200
			msg = "ok"
			album.Cover = "http://tingting-resource.bitekun.xin/resource/image/cover/"+filename+".jpg"
		}else{
			code = 402
			msg = "image format need jpg or png "
		}

	}

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer db.Close()
	if(&(Params.Summary)!=nil){
		fmt.Println("Summary is",Params.Summary)
		album.Summary = *(Params.Summary)
		// album.Summary = "summary"
	}

	album.HasPushed = 0
	album.Name = Params.Title
	//album.User_id = *(Params.MemberID)
	db.Create(&album)
	status.UnmarshalBinary([]byte(_var.Response200(code,msg)))
	response.Return = &status
	ok.SetPayload(&response)
	o.Context.Respond(rw, r, route.Produces, route, ok)

}
