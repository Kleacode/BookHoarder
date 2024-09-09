// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// Defines values for Status.
const (
	Done Status = "done"
	Todo Status = "todo"
	Wip  Status = "wip"
)

// Book defines model for book.
type Book struct {
	Title *string `json:"title,omitempty"`
}

// Error defines model for error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// ExistBook defines model for exist_book.
type ExistBook struct {
	BookId *int    `json:"bookId,omitempty"`
	Title  *string `json:"title,omitempty"`
	UserId *int    `json:"userId,omitempty"`
}

// ExistHoarderBook defines model for exist_hoarder_book.
type ExistHoarderBook struct {
	Book      *ExistBook `json:"book,omitempty"`
	HoarderId *int       `json:"hoarderId,omitempty"`
	Status    *Status    `json:"status,omitempty"`
	Tags      *Tags      `json:"tags,omitempty"`
}

// ExistTag defines model for exist_tag.
type ExistTag struct {
	Name   *string `json:"name,omitempty"`
	TagId  *int    `json:"tagId,omitempty"`
	UserId *int    `json:"userId,omitempty"`
}

// HoarderBook defines model for hoarder_book.
type HoarderBook struct {
	Book   *ExistBook `json:"book,omitempty"`
	Status *Status    `json:"status,omitempty"`
	Tags   *Tags      `json:"tags,omitempty"`
}

// NewBook defines model for new_book.
type NewBook struct {
	Title  *string `json:"title,omitempty"`
	UserId *int    `json:"userId,omitempty"`
}

// PatchHoarder defines model for patch_hoarder.
type PatchHoarder struct {
	Status *Status `json:"status,omitempty"`
	Tags   *Tags   `json:"tags,omitempty"`
}

// PostBook defines model for post_book.
type PostBook = Book

// PostHoarderExist defines model for post_hoarder_exist.
type PostHoarderExist struct {
	Status *Status `json:"status,omitempty"`
	Tags   *Tags   `json:"tags,omitempty"`
}

// PostHoarderNew defines model for post_hoarder_new.
type PostHoarderNew struct {
	Book   *Book   `json:"book,omitempty"`
	Status *Status `json:"status,omitempty"`
	Tags   *Tags   `json:"tags,omitempty"`
}

// Status defines model for status.
type Status string

// Tag defines model for tag.
type Tag struct {
	Id   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// TagInfo defines model for tag_info.
type TagInfo struct {
	Name *string `json:"name,omitempty"`
}

// Tags defines model for tags.
type Tags = []Tag

// BookId defines model for bookId.
type BookId = int

// HoarderId defines model for hoarderId.
type HoarderId = int

// TagId defines model for tagId.
type TagId = int

// UserId defines model for userId.
type UserId = int

// DefaultResponse defines model for default_response.
type DefaultResponse = Error

// GetBooksParams defines parameters for GetBooks.
type GetBooksParams struct {
	Title *string `form:"title,omitempty" json:"title,omitempty"`
}

// GetTagsParams defines parameters for GetTags.
type GetTagsParams struct {
	Name *string `form:"name,omitempty" json:"name,omitempty"`
}

// GetUserIdHoarderParams defines parameters for GetUserIdHoarder.
type GetUserIdHoarderParams struct {
	Status *string `form:"status,omitempty" json:"status,omitempty"`
	Tags   *[]int  `form:"tags,omitempty" json:"tags,omitempty"`
}

// PatchUserIdBooksBookIdJSONRequestBody defines body for PatchUserIdBooksBookId for application/json ContentType.
type PatchUserIdBooksBookIdJSONRequestBody = PostBook

// PostUserIdHoarderJSONRequestBody defines body for PostUserIdHoarder for application/json ContentType.
type PostUserIdHoarderJSONRequestBody = PostHoarderNew

// PostUserIdHoarderBookIdJSONRequestBody defines body for PostUserIdHoarderBookId for application/json ContentType.
type PostUserIdHoarderBookIdJSONRequestBody = PostHoarderExist

// PatchUserIdHoarderHoarderIdJSONRequestBody defines body for PatchUserIdHoarderHoarderId for application/json ContentType.
type PatchUserIdHoarderHoarderIdJSONRequestBody = PatchHoarder

// PostUserIdTagsJSONRequestBody defines body for PostUserIdTags for application/json ContentType.
type PostUserIdTagsJSONRequestBody = TagInfo

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// 登録されている本を取得する
	// (GET /books)
	GetBooks(c *gin.Context, params GetBooksParams)
	// 特定の本1冊の情報を取得する
	// (GET /books/{bookId})
	GetBooksBookId(c *gin.Context, bookId BookId)
	// 登録されているタグを取得する
	// (GET /tags)
	GetTags(c *gin.Context, params GetTagsParams)
	// ユーザーが登録した本を削除する
	// (DELETE /{userId}/books/{bookId})
	DeleteUserIdBooksBookId(c *gin.Context, userId UserId, bookId BookId)
	// ユーザーが登録した本の情報を更新する
	// (PATCH /{userId}/books/{bookId})
	PatchUserIdBooksBookId(c *gin.Context, userId UserId, bookId BookId)
	// ユーザーの積読リストから、積読の一覧を取得する。
	// (GET /{userId}/hoarder)
	GetUserIdHoarder(c *gin.Context, userId UserId, params GetUserIdHoarderParams)
	// 本を新しく登録する。その本をユーザーの積読リストに積読として登録する。
	// (POST /{userId}/hoarder)
	PostUserIdHoarder(c *gin.Context, userId UserId)
	// ユーザーの積読リストに既に登録済みの本から積読を登録する
	// (POST /{userId}/hoarder/{bookId})
	PostUserIdHoarderBookId(c *gin.Context, userId UserId, bookId BookId)
	// ユーザーの積読リストにある積読を削除する
	// (DELETE /{userId}/hoarder/{hoarderId})
	DeleteUserIdHoarderHoarderId(c *gin.Context, userId UserId, hoarderId HoarderId)
	// ユーザーの積読リストにある積読の状態を更新する
	// (PATCH /{userId}/hoarder/{hoarderId})
	PatchUserIdHoarderHoarderId(c *gin.Context, userId UserId, hoarderId HoarderId)
	// タグを新しく登録する
	// (POST /{userId}/tags)
	PostUserIdTags(c *gin.Context, userId UserId)
	// タグを削除する
	// (DELETE /{userId}/tags/{tagId})
	DeleteUserIdTagsTagId(c *gin.Context, userId UserId, tagId TagId)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetBooks operation middleware
func (siw *ServerInterfaceWrapper) GetBooks(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetBooksParams

	// ------------- Optional query parameter "title" -------------

	err = runtime.BindQueryParameter("form", true, false, "title", c.Request.URL.Query(), &params.Title)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter title: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetBooks(c, params)
}

// GetBooksBookId operation middleware
func (siw *ServerInterfaceWrapper) GetBooksBookId(c *gin.Context) {

	var err error

	// ------------- Path parameter "bookId" -------------
	var bookId BookId

	err = runtime.BindStyledParameterWithOptions("simple", "bookId", c.Param("bookId"), &bookId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter bookId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetBooksBookId(c, bookId)
}

// GetTags operation middleware
func (siw *ServerInterfaceWrapper) GetTags(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetTagsParams

	// ------------- Optional query parameter "name" -------------

	err = runtime.BindQueryParameter("form", true, false, "name", c.Request.URL.Query(), &params.Name)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter name: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetTags(c, params)
}

// DeleteUserIdBooksBookId operation middleware
func (siw *ServerInterfaceWrapper) DeleteUserIdBooksBookId(c *gin.Context) {

	var err error

	// ------------- Path parameter "userId" -------------
	var userId UserId

	err = runtime.BindStyledParameterWithOptions("simple", "userId", c.Param("userId"), &userId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter userId: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "bookId" -------------
	var bookId BookId

	err = runtime.BindStyledParameterWithOptions("simple", "bookId", c.Param("bookId"), &bookId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter bookId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteUserIdBooksBookId(c, userId, bookId)
}

// PatchUserIdBooksBookId operation middleware
func (siw *ServerInterfaceWrapper) PatchUserIdBooksBookId(c *gin.Context) {

	var err error

	// ------------- Path parameter "userId" -------------
	var userId UserId

	err = runtime.BindStyledParameterWithOptions("simple", "userId", c.Param("userId"), &userId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter userId: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "bookId" -------------
	var bookId BookId

	err = runtime.BindStyledParameterWithOptions("simple", "bookId", c.Param("bookId"), &bookId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter bookId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PatchUserIdBooksBookId(c, userId, bookId)
}

// GetUserIdHoarder operation middleware
func (siw *ServerInterfaceWrapper) GetUserIdHoarder(c *gin.Context) {

	var err error

	// ------------- Path parameter "userId" -------------
	var userId UserId

	err = runtime.BindStyledParameterWithOptions("simple", "userId", c.Param("userId"), &userId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter userId: %w", err), http.StatusBadRequest)
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetUserIdHoarderParams

	// ------------- Optional query parameter "status" -------------

	err = runtime.BindQueryParameter("form", true, false, "status", c.Request.URL.Query(), &params.Status)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter status: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "tags" -------------

	err = runtime.BindQueryParameter("form", false, false, "tags", c.Request.URL.Query(), &params.Tags)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter tags: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetUserIdHoarder(c, userId, params)
}

// PostUserIdHoarder operation middleware
func (siw *ServerInterfaceWrapper) PostUserIdHoarder(c *gin.Context) {

	var err error

	// ------------- Path parameter "userId" -------------
	var userId UserId

	err = runtime.BindStyledParameterWithOptions("simple", "userId", c.Param("userId"), &userId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter userId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostUserIdHoarder(c, userId)
}

// PostUserIdHoarderBookId operation middleware
func (siw *ServerInterfaceWrapper) PostUserIdHoarderBookId(c *gin.Context) {

	var err error

	// ------------- Path parameter "userId" -------------
	var userId UserId

	err = runtime.BindStyledParameterWithOptions("simple", "userId", c.Param("userId"), &userId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter userId: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "bookId" -------------
	var bookId BookId

	err = runtime.BindStyledParameterWithOptions("simple", "bookId", c.Param("bookId"), &bookId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter bookId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostUserIdHoarderBookId(c, userId, bookId)
}

// DeleteUserIdHoarderHoarderId operation middleware
func (siw *ServerInterfaceWrapper) DeleteUserIdHoarderHoarderId(c *gin.Context) {

	var err error

	// ------------- Path parameter "userId" -------------
	var userId UserId

	err = runtime.BindStyledParameterWithOptions("simple", "userId", c.Param("userId"), &userId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter userId: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "hoarderId" -------------
	var hoarderId HoarderId

	err = runtime.BindStyledParameterWithOptions("simple", "hoarderId", c.Param("hoarderId"), &hoarderId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter hoarderId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteUserIdHoarderHoarderId(c, userId, hoarderId)
}

// PatchUserIdHoarderHoarderId operation middleware
func (siw *ServerInterfaceWrapper) PatchUserIdHoarderHoarderId(c *gin.Context) {

	var err error

	// ------------- Path parameter "userId" -------------
	var userId UserId

	err = runtime.BindStyledParameterWithOptions("simple", "userId", c.Param("userId"), &userId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter userId: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "hoarderId" -------------
	var hoarderId HoarderId

	err = runtime.BindStyledParameterWithOptions("simple", "hoarderId", c.Param("hoarderId"), &hoarderId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter hoarderId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PatchUserIdHoarderHoarderId(c, userId, hoarderId)
}

// PostUserIdTags operation middleware
func (siw *ServerInterfaceWrapper) PostUserIdTags(c *gin.Context) {

	var err error

	// ------------- Path parameter "userId" -------------
	var userId UserId

	err = runtime.BindStyledParameterWithOptions("simple", "userId", c.Param("userId"), &userId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter userId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostUserIdTags(c, userId)
}

// DeleteUserIdTagsTagId operation middleware
func (siw *ServerInterfaceWrapper) DeleteUserIdTagsTagId(c *gin.Context) {

	var err error

	// ------------- Path parameter "userId" -------------
	var userId UserId

	err = runtime.BindStyledParameterWithOptions("simple", "userId", c.Param("userId"), &userId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter userId: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "tagId" -------------
	var tagId TagId

	err = runtime.BindStyledParameterWithOptions("simple", "tagId", c.Param("tagId"), &tagId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter tagId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteUserIdTagsTagId(c, userId, tagId)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/books", wrapper.GetBooks)
	router.GET(options.BaseURL+"/books/:bookId", wrapper.GetBooksBookId)
	router.GET(options.BaseURL+"/tags", wrapper.GetTags)
	router.DELETE(options.BaseURL+"/:userId/books/:bookId", wrapper.DeleteUserIdBooksBookId)
	router.PATCH(options.BaseURL+"/:userId/books/:bookId", wrapper.PatchUserIdBooksBookId)
	router.GET(options.BaseURL+"/:userId/hoarder", wrapper.GetUserIdHoarder)
	router.POST(options.BaseURL+"/:userId/hoarder", wrapper.PostUserIdHoarder)
	router.POST(options.BaseURL+"/:userId/hoarder/:bookId", wrapper.PostUserIdHoarderBookId)
	router.DELETE(options.BaseURL+"/:userId/hoarder/:hoarderId", wrapper.DeleteUserIdHoarderHoarderId)
	router.PATCH(options.BaseURL+"/:userId/hoarder/:hoarderId", wrapper.PatchUserIdHoarderHoarderId)
	router.POST(options.BaseURL+"/:userId/tags", wrapper.PostUserIdTags)
	router.DELETE(options.BaseURL+"/:userId/tags/:tagId", wrapper.DeleteUserIdTagsTagId)
}
