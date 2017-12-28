// Code generated by go-swagger; DO NOT EDIT.

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"github.com/tingtingbackend/restapi/operations"
	"github.com/tingtingbackend/restapi/operations/album"
	"github.com/tingtingbackend/restapi/operations/banner"
	"github.com/tingtingbackend/restapi/operations/book"
	"github.com/tingtingbackend/restapi/operations/category"
	"github.com/tingtingbackend/restapi/operations/chapter"
	"github.com/tingtingbackend/restapi/operations/feedback"
	"github.com/tingtingbackend/restapi/operations/member"
	"github.com/tingtingbackend/restapi/operations/msg"
	"github.com/tingtingbackend/restapi/operations/order"
	"github.com/tingtingbackend/restapi/operations/relation"
	"github.com/tingtingbackend/restapi/operations/report_err"
	"github.com/tingtingbackend/restapi/operations/user"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name  --spec ../swagger.json

func configureFlags(api *operations.TingtingBackendAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TingtingBackendAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.MultipartformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	api.RelationNrCategorySubCategoryListRelationHandler = relation.NrCategorySubCategoryListRelationHandlerFunc(func(params relation.NrCategorySubCategoryListRelationParams) middleware.Responder {
		return middleware.NotImplemented("operation relation.NrCategorySubCategoryListRelation has not yet been implemented")
	})
	api.RelationNrAlbumBookListRelationHandler = relation.NrAlbumBookListRelationHandlerFunc(func(params relation.NrAlbumBookListRelationParams) middleware.Responder {
		return middleware.NotImplemented("operation relation.NrAlbumBookListRelation has not yet been implemented")
	})
	api.MemberNrMemberRecordListHandler = member.NrMemberRecordListHandlerFunc(func(params member.NrMemberRecordListParams) middleware.Responder {
		return middleware.NotImplemented("operation member.NrMemberRecordList has not yet been implemented")
	})
	api.RelationNrRelationCategorySubCategoryListEditHandler = relation.NrRelationCategorySubCategoryListEditHandlerFunc(func(params relation.NrRelationCategorySubCategoryListEditParams) middleware.Responder {
		return middleware.NotImplemented("operation relation.NrRelationCategorySubCategoryListEdit has not yet been implemented")
	})
	api.RelationNrRelationAlbumBooklistEditHandler = relation.NrRelationAlbumBooklistEditHandlerFunc(func(params relation.NrRelationAlbumBooklistEditParams) middleware.Responder {
		return middleware.NotImplemented("operation relation.NrRelationAlbumBooklistEdit has not yet been implemented")
	})
	api.RelationNrRelationSubCategoryAlbumListEditHandler = relation.NrRelationSubCategoryAlbumListEditHandlerFunc(func(params relation.NrRelationSubCategoryAlbumListEditParams) middleware.Responder {
		return middleware.NotImplemented("operation relation.NrRelationSubCategoryAlbumListEdit has not yet been implemented")
	})
	api.RelationNrSubCategoryAlbumListRelationHandler = relation.NrSubCategoryAlbumListRelationHandlerFunc(func(params relation.NrSubCategoryAlbumListRelationParams) middleware.Responder {
		return middleware.NotImplemented("operation relation.NrSubCategoryAlbumListRelation has not yet been implemented")
	})
	api.UserNrUserDeleteHandler = user.NrUserDeleteHandlerFunc(func(params user.NrUserDeleteParams) middleware.Responder {
		return middleware.NotImplemented("operation user.NrUserDelete has not yet been implemented")
	})
	api.UserNrUserEditPassHandler = user.NrUserEditPassHandlerFunc(func(params user.NrUserEditPassParams) middleware.Responder {
		return middleware.NotImplemented("operation user.NrUserEditPass has not yet been implemented")
	})
	api.UserNrUserListHandler = user.NrUserListHandlerFunc(func(params user.NrUserListParams) middleware.Responder {
		return middleware.NotImplemented("operation user.NrUserList has not yet been implemented")
	})
	api.UserNrUserLoginHandler = user.NrUserLoginHandlerFunc(func(params user.NrUserLoginParams) middleware.Responder {
		return middleware.NotImplemented("operation user.NrUserLogin has not yet been implemented")
	})
	api.AlbumAlbumDeleteHandler = album.AlbumDeleteHandlerFunc(func(params album.AlbumDeleteParams) middleware.Responder {
		return middleware.NotImplemented("operation album.AlbumDelete has not yet been implemented")
	})
	api.AlbumAlbumDetailHandler = album.AlbumDetailHandlerFunc(func(params album.AlbumDetailParams) middleware.Responder {
		return middleware.NotImplemented("operation album.AlbumDetail has not yet been implemented")
	})
	api.AlbumAlbumEditHandler = album.AlbumEditHandlerFunc(func(params album.AlbumEditParams) middleware.Responder {
		return middleware.NotImplemented("operation album.AlbumEdit has not yet been implemented")
	})
	api.AlbumAlbumListHandler = album.AlbumListHandlerFunc(func(params album.AlbumListParams) middleware.Responder {
		return middleware.NotImplemented("operation album.AlbumList has not yet been implemented")
	})
	api.AlbumAlbumUploadHandler = album.AlbumUploadHandlerFunc(func(params album.AlbumUploadParams) middleware.Responder {
		return middleware.NotImplemented("operation album.AlbumUpload has not yet been implemented")
	})
	api.BannerBannerDetailHandler = banner.BannerDetailHandlerFunc(func(params banner.BannerDetailParams) middleware.Responder {
		return middleware.NotImplemented("operation banner.BannerDetail has not yet been implemented")
	})
	api.BannerBannerListHandler = banner.BannerListHandlerFunc(func(params banner.BannerListParams) middleware.Responder {
		return middleware.NotImplemented("operation banner.BannerList has not yet been implemented")
	})
	api.RelationBookChapterListRelationHandler = relation.BookChapterListRelationHandlerFunc(func(params relation.BookChapterListRelationParams) middleware.Responder {
		return middleware.NotImplemented("operation relation.BookChapterListRelation has not yet been implemented")
	})
	api.RelationBookChapterListRelationEditHandler = relation.BookChapterListRelationEditHandlerFunc(func(params relation.BookChapterListRelationEditParams) middleware.Responder {
		return middleware.NotImplemented("operation relation.BookChapterListRelationEdit has not yet been implemented")
	})
	api.BookBookDeleteHandler = book.BookDeleteHandlerFunc(func(params book.BookDeleteParams) middleware.Responder {
		return middleware.NotImplemented("operation book.BookDelete has not yet been implemented")
	})
	api.BookBookDetailHandler = book.BookDetailHandlerFunc(func(params book.BookDetailParams) middleware.Responder {
		return middleware.NotImplemented("operation book.BookDetail has not yet been implemented")
	})
	api.BookBookEditHandler = book.BookEditHandlerFunc(func(params book.BookEditParams) middleware.Responder {
		return middleware.NotImplemented("operation book.BookEdit has not yet been implemented")
	})
	api.BookBookListHandler = book.BookListHandlerFunc(func(params book.BookListParams) middleware.Responder {
		return middleware.NotImplemented("operation book.BookList has not yet been implemented")
	})
	api.BookBookUploadHandler = book.BookUploadHandlerFunc(func(params book.BookUploadParams) middleware.Responder {
		return middleware.NotImplemented("operation book.BookUpload has not yet been implemented")
	})
	api.CategoryCategoryDeleteHandler = category.CategoryDeleteHandlerFunc(func(params category.CategoryDeleteParams) middleware.Responder {
		return middleware.NotImplemented("operation category.CategoryDelete has not yet been implemented")
	})
	api.CategoryCategoryDetailHandler = category.CategoryDetailHandlerFunc(func(params category.CategoryDetailParams) middleware.Responder {
		return middleware.NotImplemented("operation category.CategoryDetail has not yet been implemented")
	})
	api.CategoryCategoryEditHandler = category.CategoryEditHandlerFunc(func(params category.CategoryEditParams) middleware.Responder {
		return middleware.NotImplemented("operation category.CategoryEdit has not yet been implemented")
	})
	api.CategoryCategorySubEditHandler = category.CategorySubEditHandlerFunc(func(params category.CategorySubEditParams) middleware.Responder {
		return middleware.NotImplemented("operation category.CategorySubEdit has not yet been implemented")
	})
	api.CategoryCategorySubUploadHandler = category.CategorySubUploadHandlerFunc(func(params category.CategorySubUploadParams) middleware.Responder {
		return middleware.NotImplemented("operation category.CategorySubUpload has not yet been implemented")
	})
	api.CategoryCategoryUploadHandler = category.CategoryUploadHandlerFunc(func(params category.CategoryUploadParams) middleware.Responder {
		return middleware.NotImplemented("operation category.CategoryUpload has not yet been implemented")
	})
	api.ChapterChapterDeleteHandler = chapter.ChapterDeleteHandlerFunc(func(params chapter.ChapterDeleteParams) middleware.Responder {
		return middleware.NotImplemented("operation chapter.ChapterDelete has not yet been implemented")
	})
	api.ChapterChapterDetailHandler = chapter.ChapterDetailHandlerFunc(func(params chapter.ChapterDetailParams) middleware.Responder {
		return middleware.NotImplemented("operation chapter.ChapterDetail has not yet been implemented")
	})
	api.ChapterChapterEditHandler = chapter.ChapterEditHandlerFunc(func(params chapter.ChapterEditParams) middleware.Responder {
		return middleware.NotImplemented("operation chapter.ChapterEdit has not yet been implemented")
	})
	api.ChapterChapterFavListHandler = chapter.ChapterFavListHandlerFunc(func(params chapter.ChapterFavListParams) middleware.Responder {
		return middleware.NotImplemented("operation chapter.ChapterFavList has not yet been implemented")
	})
	api.ChapterChapterHistoryListHandler = chapter.ChapterHistoryListHandlerFunc(func(params chapter.ChapterHistoryListParams) middleware.Responder {
		return middleware.NotImplemented("operation chapter.ChapterHistoryList has not yet been implemented")
	})
	api.ChapterChapterUploadHandler = chapter.ChapterUploadHandlerFunc(func(params chapter.ChapterUploadParams) middleware.Responder {
		return middleware.NotImplemented("operation chapter.ChapterUpload has not yet been implemented")
	})
	api.FeedbackFeedbackDetailHandler = feedback.FeedbackDetailHandlerFunc(func(params feedback.FeedbackDetailParams) middleware.Responder {
		return middleware.NotImplemented("operation feedback.FeedbackDetail has not yet been implemented")
	})
	api.FeedbackFeedbackListHandler = feedback.FeedbackListHandlerFunc(func(params feedback.FeedbackListParams) middleware.Responder {
		return middleware.NotImplemented("operation feedback.FeedbackList has not yet been implemented")
	})
	api.MemberMemberDetailHandler = member.MemberDetailHandlerFunc(func(params member.MemberDetailParams) middleware.Responder {
		return middleware.NotImplemented("operation member.MemberDetail has not yet been implemented")
	})
	api.MemberMemberListHandler = member.MemberListHandlerFunc(func(params member.MemberListParams) middleware.Responder {
		return middleware.NotImplemented("operation member.MemberList has not yet been implemented")
	})
	api.MemberMemberRecordDetailHandler = member.MemberRecordDetailHandlerFunc(func(params member.MemberRecordDetailParams) middleware.Responder {
		return middleware.NotImplemented("operation member.MemberRecordDetail has not yet been implemented")
	})
	api.MsgMsgDetailHandler = msg.MsgDetailHandlerFunc(func(params msg.MsgDetailParams) middleware.Responder {
		return middleware.NotImplemented("operation msg.MsgDetail has not yet been implemented")
	})
	api.MsgMsgSendListHandler = msg.MsgSendListHandlerFunc(func(params msg.MsgSendListParams) middleware.Responder {
		return middleware.NotImplemented("operation msg.MsgSendList has not yet been implemented")
	})
	api.MsgMsgUnsendListHandler = msg.MsgUnsendListHandlerFunc(func(params msg.MsgUnsendListParams) middleware.Responder {
		return middleware.NotImplemented("operation msg.MsgUnsendList has not yet been implemented")
	})
	api.OrderOrderDetailHandler = order.OrderDetailHandlerFunc(func(params order.OrderDetailParams) middleware.Responder {
		return middleware.NotImplemented("operation order.OrderDetail has not yet been implemented")
	})
	api.OrderOrderListHandler = order.OrderListHandlerFunc(func(params order.OrderListParams) middleware.Responder {
		return middleware.NotImplemented("operation order.OrderList has not yet been implemented")
	})
	api.ReportErrReportErrDetailHandler = report_err.ReportErrDetailHandlerFunc(func(params report_err.ReportErrDetailParams) middleware.Responder {
		return middleware.NotImplemented("operation report_err.ReportErrDetail has not yet been implemented")
	})
	api.ReportErrReportErrListHandler = report_err.ReportErrListHandlerFunc(func(params report_err.ReportErrListParams) middleware.Responder {
		return middleware.NotImplemented("operation report_err.ReportErrList has not yet been implemented")
	})
	api.UserSearchHandler = user.SearchHandlerFunc(func(params user.SearchParams) middleware.Responder {
		return middleware.NotImplemented("operation user.Search has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
