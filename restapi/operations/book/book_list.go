// Code generated by go-swagger; DO NOT EDIT.

package book

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

// BookListHandlerFunc turns a function with the right signature into a book list handler
type BookListHandlerFunc func(BookListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn BookListHandlerFunc) Handle(params BookListParams) middleware.Responder {
	return fn(params)
}

// BookListHandler interface for that can handle valid book list params
type BookListHandler interface {
	Handle(BookListParams) middleware.Responder
}

// NewBookList creates a new http.Handler for the book list operation
func NewBookList(ctx *middleware.Context, handler BookListHandler) *BookList {
	return &BookList{Context: ctx, Handler: handler}
}

/*BookList swagger:route GET /book/list Book bookList

获取书本列表

获取书本列表

*/
type BookList struct {
	Context *middleware.Context
	Handler BookListHandler
}

func (o *BookList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewBookListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok BookListOK
	var response models.InlineResponse20021
	var bookList models.InlineResponse20021BookList
	var count int64

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer db.Close()
	//query
	if Params.Memberid !=nil{
		fmt.Println("get fav booklist",*(Params.Memberid))
		db.Raw("select books.id, books.name,books.icon,fav_book.time from books left join fav_book on books.id = fav_book.book_id where books.status=0 and fav_book.member_id=?",*(Params.Memberid)).Limit(*(Params.PageSize)).Offset(*(Params.PageIndex) * (*(Params.PageSize))).Find(&bookList)
		db.Raw("select books.id, books.name,books.icon,fav_book.time from books left join fav_book on books.id = fav_book.book_id where books.status=0 and fav_book.member_id=?",*(Params.Memberid)).Count(&count)
		fmt.Println("favbookcount is",count)
		//db.Table("albums").Select("albums.id, albums.name").Joins("left join fav_album on albums.id = fav_album.album_id").Where("fav_album.member_id =?", *(Params.Memberid)).Where("albums.status=?", 0).Where("fav_album.status =?", 0).Count(&count)
		//db.Table("albums").Select("albums.id, albums.name").Joins("left join fav_album on albums.id = fav_album.album_id").Where("fav_album.member_id =?", *(Params.Memberid)).Where("albums.status=?", 0).Where("fav_album.status =?", 0).Limit(*(Params.PageSize)).Offset(*(Params.PageIndex) * (*(Params.PageSize))).Find(&albumList)
	}else {
		if Params.Keyword != nil && Params.AlbumID != nil {
			if (*Params.Keyword == " ") {
				db.Raw("select id,name  FROM books where  status=0 and id not in (select bookId from album_book_relation  where status=0 and albumId = ? )", *(Params.AlbumID)).Limit(*(Params.PageSize)).Offset(*(Params.PageIndex) * (*(Params.PageSize))).Find(&bookList)
				db.Raw("select id,name  FROM books where status=0 and  id not in (select bookId from album_book_relation  where status=0 and albumId = ? )", *(Params.AlbumID)).Count(&count)
			} else {
				db.Raw("select id,name  FROM books where status=0 and name like '%" + *(Params.Keyword)+"%' and id not in (select bookId from album_book_relation  where status=0 and albumId = ? )", *(Params.AlbumID)).Limit(*(Params.PageSize)).Offset(*(Params.PageIndex) * (*(Params.PageSize))).Find(&bookList)
				db.Raw("select id,name  FROM books where status=0 and name like '%" + *(Params.Keyword)+"%' and id not in (select bookId from album_book_relation  where status=0 and albumId = ? )", *(Params.AlbumID)).Count(&count)
			}
			//db.Table("books").Select("books.id, books.name").Joins("left join album_book_relation on books.id = album_book_relation.bookId").Where("books.id is null").Find(&bookList)
			fmt.Println("1")
			//db.Where(map[string]interface{}{"status":0}).Where("name like ?","%"+*(Params.Keyword)+"%").Not("id",).Find(&bookList).Offset(*(Params.PageIndex)*(*(Params.PageSize)))
		} else {
			if Params.AlbumID != nil {
				fmt.Println("2")
				db.Table("books").Select("album_book_relation.id, books.name, album_book_relation.order").Joins("left join album_book_relation on books.id = album_book_relation.bookId").Order("album_book_relation.order").Where("album_book_relation.albumId =?", *Params.AlbumID).Where("books.status=?", 0).Where("album_book_relation.status =?", 0).Count(&count)
				db.Table("books").Select("album_book_relation.id, books.name, album_book_relation.order").Joins("left join album_book_relation on books.id = album_book_relation.bookId").Order("album_book_relation.order").Where("album_book_relation.albumId =?", *Params.AlbumID).Where("books.status=?", 0).Where("album_book_relation.status =?", 0).Limit(*(Params.PageSize)).Offset(*(Params.PageIndex) * (*(Params.PageSize))).Find(&bookList)
			} else {
				fmt.Println("3")
				db.Table("books").Where(map[string]interface{}{"status": 0}).Count(&count)
				db.Table("books").Where(map[string]interface{}{"status": 0}).Limit(*(Params.PageSize)).Offset(*(Params.PageIndex) * (*(Params.PageSize))).Find(&bookList)
			}

		}
	}

	//data
	response.BookList = bookList

	//status
	var status models.Response
	status.UnmarshalBinary([]byte(_var.Response200(200,"ok")))
	response.Status = &status
	response.Status.TotalCount = count

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}
