// Code generated by go-swagger; DO NOT EDIT.

package relation

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
	"strings"
	"strconv"
)

// BookChapterListRelationEditHandlerFunc turns a function with the right signature into a book chapter list relation edit handler
type BookChapterListRelationEditHandlerFunc func(BookChapterListRelationEditParams) middleware.Responder

// Handle executing the request and returning a response
func (fn BookChapterListRelationEditHandlerFunc) Handle(params BookChapterListRelationEditParams) middleware.Responder {
	return fn(params)
}

// BookChapterListRelationEditHandler interface for that can handle valid book chapter list relation edit params
type BookChapterListRelationEditHandler interface {
	Handle(BookChapterListRelationEditParams) middleware.Responder
}

// NewBookChapterListRelationEdit creates a new http.Handler for the book chapter list relation edit operation
func NewBookChapterListRelationEdit(ctx *middleware.Context, handler BookChapterListRelationEditHandler) *BookChapterListRelationEdit {
	return &BookChapterListRelationEdit{Context: ctx, Handler: handler}
}

/*BookChapterListRelationEdit swagger:route POST /relation/book/chapterList/edit Relation bookChapterListRelationEdit

编辑书的章节集合

编辑书的章节集合

*/
type BookChapterListRelationEdit struct {
	Context *middleware.Context
	Handler BookChapterListRelationEditHandler
}

type SeedModel struct{
	Seed int64 `json:"seed"`
}

func (o *BookChapterListRelationEdit) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewBookChapterListRelationEditParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok BookChapterListRelationEditOK
	var response models.InlineResponse20018


	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}

	chapters := *(Params.Body.ChapterIds)

	if (*(Params.Body.ActionCode) == 0){ //添加映射
		//先解析出bookis集合,样式 1,2,3,4,
		if (!strings.Contains(chapters,",")){
			var seedmodel SeedModel
			db.Raw("select max(`order`)+1 as seed from book_chapter_relation").First(&seedmodel)
			fmt.Println("seed is",seedmodel.Seed)
			db.Exec("insert into book_chapter_relation(bookId,chapterId,`order`) values(?,?,?)",*(Params.Body.BookID),chapters,seedmodel.Seed)
			db.Exec("update books set clips_number=clips_number+1 where id=?",*(Params.Body.BookID))
			fmt.Println("1")
		}else{
			temp := strings.Split(chapters,",")
			for k:=0;k< len(temp);k++ {
				//id 自增
				var seedmodel SeedModel
				db.Raw("select max(`order`)+1 as seed from book_chapter_relation").First(&seedmodel)
				fmt.Println("seed is",seedmodel.Seed)
				db.Exec("insert into book_chapter_relation(bookId,chapterId,`order`) values(?,?,?)",*(Params.Body.BookID),temp[k],seedmodel.Seed)
				fmt.Println("insert into book_chapter_relation(bookId,chapterId) values(?,?)",Params.Body.BookID,temp[k])
			}
			db.Exec("update books set clips_number=clips_number+"+strconv.Itoa(len(temp))+" where id="+strconv.FormatInt(*(Params.Body.BookID),10))
		}
	}else{ //去除映射
		db.Exec("delete from book_chapter_relation where bookId=? and id=?",*(Params.Body.BookID),chapters)
		fmt.Println("delete from book_chapter_relation where bookId=? and id=?",Params.Body.BookID,chapters)
		db.Exec("update books set clips_number=clips_number-1 where id=?",*(Params.Body.BookID))
	}

	var status models.Response
	status.UnmarshalBinary([]byte(_var.Response200(200,"ok")))
	response.Status = &status

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}
