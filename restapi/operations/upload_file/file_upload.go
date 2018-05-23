// Code generated by go-swagger; DO NOT EDIT.

package upload_file

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	"io/ioutil"
	"github.com/go-openapi/runtime/middleware"
	"fmt"
	_ "os"
	"runtime"
	"time"
	"strings"
	"tingtingbackend/models"
	"tingtingbackend/var"
	"strconv"
)

// FileUploadHandlerFunc turns a function with the right signature into a file upload handler
type FileUploadHandlerFunc func(FileUploadParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FileUploadHandlerFunc) Handle(params FileUploadParams) middleware.Responder {
	return fn(params)
}

// FileUploadHandler interface for that can handle valid file upload params
type FileUploadHandler interface {
	Handle(FileUploadParams) middleware.Responder
}

// NewFileUpload creates a new http.Handler for the file upload operation
func NewFileUpload(ctx *middleware.Context, handler FileUploadHandler) *FileUpload {
	return &FileUpload{Context: ctx, Handler: handler}
}

/*FileUpload swagger:route POST /file/upload UploadFile fileUpload

上传文件

*/
type FileUpload struct {
	Context *middleware.Context
	Handler FileUploadHandler
}
type SeedModel1 struct{
	Seed int64 `json:"seed"`
}
func (o *FileUpload) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewFileUploadParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok FileUploadOK
	var response models.InlineResponse200191
	var status models.Response

	var msg string
	var code int64

	var filename string
	filename = strconv.FormatInt((time.Now().UnixNano()), 10)

	//fmt.Println("filename is", *Params.Filename)

	//fmt.Println("Param.type is", Params.Type)

	var contentType string
	//var file []byte
	code = 200
	msg = "ok"

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer db.Close()

	if (Params.File != nil) {
		file, err := ioutil.ReadAll(Params.File)
		if err != nil {
			fmt.Println("err upload:", err.Error())
		}
		contentType = http.DetectContentType(file)
		fmt.Println("contentType is", contentType)
		fmt.Println("file.size is", len(file))

		if contentType == "application/octet-stream" || contentType == "video/mp4" {
			fmt.Println("10")
			fmt.Println("contentType is", contentType)
			if (runtime.GOOS == "windows") {
				err1 := ioutil.WriteFile(filename+".m4a", file, 0644)
				if err1 != nil {
					fmt.Println(err1.Error())
				}
			} else {
				fmt.Println("11")
				fmt.Println("file.size2 is", len(file))
				err1 := ioutil.WriteFile("/root/go/src/resource/mp3/"+filename+".m4a", file, 0644)
				if err1 != nil {
					fmt.Println(err1.Error())
				}
			}
			code = 200
			msg = "ok"
			response.URL = _var.GetResourceDomain("m4a") + filename + ".m4a"
			if(Params.Filename != nil){ //==nil说明是批量上传章节
			title := strings.Split(*(Params.Filename),".")
			db.Exec("insert into chapters(name,status,summary,url,time) values(?,?,?,?,?)",title[0],0,title[0],response.URL,time.Now().UnixNano() / 1000000000)
			//判断是否传了书本id参数
			if Params.BookId != nil {
				//将新添加的章节映射到传入书本id
				fmt.Println("add relation bookid is",Params.BookId)
				var tempChapterId int64
				db.Raw("select id from chapters where name=? order by id desc limit 0,1",title[0]).Find(tempChapterId)
				//添加映射关系
				var seedmodel SeedModel1
				db.Raw("select max(`order`)+1 as seed from book_chapter_relation").First(&seedmodel)
				db.Exec("insert into book_chapter_relation(bookId,chapterId,`order`) values(?,?,?)",*(Params.BookId),tempChapterId,seedmodel.Seed)
				fmt.Println("insert into book_chapter_relation(bookId,chapterId,`order`) values(?,?,?)",*(Params.BookId),tempChapterId,seedmodel.Seed)
				db.Exec("update books set clips_number=clips_number+1 where id=?",*(Params.BookId))
				fmt.Println("update books set clips_number=clips_number+1 where id=?",*(Params.BookId))

			}

			}
		} else {
			if (Params.Type == nil) {
				fmt.Println("1")
				//save
				var lower string
				lower = strings.ToLower(contentType)
				fmt.Println("lower is", lower)
				fmt.Println("file.size is", len(file))
				if (true) {
					if (runtime.GOOS == "windows") {
						err1 := ioutil.WriteFile(filename+".jpg", file, 0644)
						if err1 != nil {
							fmt.Println(err1.Error())
						}
					} else {
						err1 := ioutil.WriteFile("/root/go/src/resource/image/icon/"+filename+".jpg", file, 0644)
						if err1 != nil {
							fmt.Println(err1.Error())
						}
					}
					response.URL = _var.GetResourceDomain("icon") + filename + ".jpg"
					code = 200
					msg = "ok"
				} else {
					fmt.Println("2")
					code = 401
					msg = "image format need jpg or png"
				}
			} else {
				fmt.Println("3")
				//如果有icon
				if (*Params.Type) == "icon" {
					fmt.Println("4")
					if (true) {
						fmt.Println("5")
						if (runtime.GOOS == "windows") {
							err1 := ioutil.WriteFile(filename+".jpg", file, 0644)
							if err1 != nil {
								fmt.Println(err1.Error())
							}
						} else {
							err1 := ioutil.WriteFile("/root/go/src/resource/image/icon/"+filename+".jpg", file, 0644)
							if err1 != nil {
								fmt.Println(err1.Error())
							}
						}
						response.URL = _var.GetResourceDomain("icon") + filename + ".jpg"
						code = 200
						msg = "ok"
					} else {
						fmt.Println("6")
						code = 402
						msg = "image format need jpg or png"
					}
				} else if (*Params.Type) == "cover" {
					fmt.Println("7")
					if (true) {
						fmt.Println("8")
						if (runtime.GOOS == "windows") {
							err1 := ioutil.WriteFile(filename+".jpg", file, 0644)
							if err1 != nil {
								fmt.Println(err1.Error())
							}
						} else {
							err1 := ioutil.WriteFile("/root/go/src/resource/image/cover/"+filename+".jpg", file, 0644)
							if err1 != nil {
								fmt.Println(err1.Error())
							}
						}
						code = 200
						msg = "ok"
						response.URL = _var.GetResourceDomain("cover") + filename + ".jpg"
					} else {
						fmt.Println("9")
						code = 403
						msg = "image format need jpg or png "
					}

				}

			}

		}
	}

	if (Params.Filename!=nil){
		response.OriginName = *(Params.Filename)
	}

	status.UnmarshalBinary([]byte(_var.Response200(code, msg)))
	response.Return = &status

	ok.SetPayload(&response)
	o.Context.Respond(rw, r, route.Produces, route, ok)

}
