// Code generated by go-swagger; DO NOT EDIT.

package chapter

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

// ChapterListHandlerFunc turns a function with the right signature into a chapter list handler
type ChapterListHandlerFunc func(ChapterListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ChapterListHandlerFunc) Handle(params ChapterListParams) middleware.Responder {
	return fn(params)
}

// ChapterListHandler interface for that can handle valid chapter list params
type ChapterListHandler interface {
	Handle(ChapterListParams) middleware.Responder
}

// NewChapterList creates a new http.Handler for the chapter list operation
func NewChapterList(ctx *middleware.Context, handler ChapterListHandler) *ChapterList {
	return &ChapterList{Context: ctx, Handler: handler}
}

/*ChapterList swagger:route GET /chapter/list Chapter chapterList

章节列表

章节列表

*/
type ChapterList struct {
	Context *middleware.Context
	Handler ChapterListHandler
}

func (o *ChapterList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewChapterListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok ChapterListOK
	var response models.InlineResponse20023
	var chapterlist models.InlineResponse20023Chapters
	var count int64

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer db.Close()
	//query
	if Params.Keyword !=nil && Params.BookID!=nil{
		if(*Params.Keyword == " ") {
			fmt.Println("1")
			db.Raw("select id,name  FROM chapters where status=0 and  id not in (select chapterId from book_chapter_relation  where bookId = ? and status=0)", *(Params.BookID)).Find(&chapterlist)
			db.Raw("select id,name  FROM chapters where status=0 and  id not in (select chapterId from book_chapter_relation  where bookId = ? and status=0)", *(Params.BookID)).Count(&count)
		}else{
			fmt.Println("2")
			db.Raw("select id,name  FROM chapters where status=0 and name like '%" + *(Params.Keyword)+"%' and id not in (select chapterId from book_chapter_relation  where bookId = ? and status=0)", *(Params.BookID)).Find(&chapterlist)
			db.Raw("select id,name  FROM chapters where status=0 and name like '%" + *(Params.Keyword)+"%' and id not in (select chapterId from book_chapter_relation  where bookId = ? and status=0)", *(Params.BookID)).Count(&count)
		}
	}else{
		if Params.BookID !=nil{
			fmt.Println("3")
			db.Raw("select book_chapter_relation.id, chapters.name,book_chapter_relation.`order` from chapters left join book_chapter_relation on chapters.id = book_chapter_relation.chapterId where book_chapter_relation.bookId =? and chapters.status=0 and book_chapter_relation.status=0",*Params.BookID).Order("book_chapter_relation.order").Limit(*(Params.PageSize)).Offset(*(Params.PageIndex) * (*(Params.PageSize))).Find(&chapterlist)
			//db.Table("chapters").Select("chapters.id, chapters.name,book_chapter_relation.order").Joins("left join book_chapter_relation on chapters.id = book_chapter_relation.chapterId").Where("book_chapter_relation.bookId =?",*Params.BookID).Where("chapters.status=?",0).Where("book_chapter_relation.status=?",0).Order("book_chapter_relation").Find(&chapterlist)
			db.Table("chapters").Select("chapters.id, chapters.name").Joins("left join book_chapter_relation on chapters.id = book_chapter_relation.chapterId").Where("book_chapter_relation.bookId =?",*Params.BookID).Where("chapters.status=?",0).Where("book_chapter_relation.status=?",0).Count(&count)
		}else{
			if (Params.Key != nil){
				fmt.Println("4")
				db.Raw("select chapters.id,chapters.name,chapters.play_count,chapters.status,chapters.time,books.name as book_name from chapters left join book_chapter_relation on chapters.id = book_chapter_relation.chapterId left join books on book_chapter_relation.bookId=books.id where chapters.status=0 and chapters.name like '%" + *(Params.Key)+"%'").Limit(*(Params.PageSize)).Offset(*(Params.PageIndex)*(*(Params.PageSize))).Find(&chapterlist)
				//db.Raw("select count(id) from chapters where status=0").Count(&count)
				db.Raw("select count(chapters.id) from chapters left join book_chapter_relation on chapters.id = book_chapter_relation.chapterId left join books on book_chapter_relation.bookId=books.id where chapters.status=0 and chapters.name like '%" + *(Params.Key)+"%'").Count(&count)
				//db.Raw("select * from chapters where status=0 and name like '%" + *(Params.Key)+"%' order by id desc").Limit(*(Params.PageSize)).Offset(*(Params.PageIndex)*(*(Params.PageSize))).Find(&chapterlist)
				//fmt.Println("select * from chapters where status=0 and name like '%" + *(Params.Key)+"%' order by id desc").Count(&count)
				//db.Raw("select count(id) from chapters where status=0 and name like '%" + *(Params.Key)+"%' ").Count(&count)
				fmt.Println("select chapters.id,chapters.name,chapters.play_count,chapters.status,chapters.time,books.name as book_name from chapters left join book_chapter_relation on chapters.id = book_chapter_relation.chapterId left join books on book_chapter_relation.bookId=books.id where chapters.status=0 and chapters.name like '%" + *(Params.Key)+"%'")
				//db.Table("chapters").Where(map[string]interface{}{"status":0}).Count(&count)
			}else{
				fmt.Println("5")
				//db.Table("chapters").Where(map[string]interface{}{"status":0}).Order("id desc").Limit(*(Params.PageSize)).Offset(*(Params.PageIndex)*(*(Params.PageSize))).Find(&chapterlist)
				//db.Table("chapters").Where(map[string]interface{}{"status":0}).Count(&count)
				db.Raw("select chapters.id,chapters.name,chapters.play_count,chapters.status,chapters.time,books.name as book_name from chapters left join book_chapter_relation on chapters.id = book_chapter_relation.chapterId left join books on book_chapter_relation.bookId=books.id where chapters.status=0").Limit(*(Params.PageSize)).Offset(*(Params.PageIndex)*(*(Params.PageSize))).Find(&chapterlist)
				//db.Raw("select count(id) from chapters where status=0").Count(&count)
				db.Raw("select count(chapters.id) from chapters left join book_chapter_relation on chapters.id = book_chapter_relation.chapterId left join books on book_chapter_relation.bookId=books.id where chapters.status=0").Count(&count)
			}

		}

	}
	//db.Where(map[string]interface{}{"status":0}).Find(&chapterlist).Limit(*(Params.PageSize)).Offset(*(Params.PageIndex)*(*(Params.PageSize)))
	//data
	response.Chapters = chapterlist

	//status
	var status models.Response
	status.UnmarshalBinary([]byte(_var.Response200(200,"ok")))
	response.Status = &status
	response.Status.TotalCount = count

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}
