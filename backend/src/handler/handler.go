package handler

import (
	api "back/src/generated"

	"github.com/gin-gonic/gin"
)

type ServiceInterface interface {
	GetBooks(c *gin.Context, params api.GetBooksParams) ([]api.ExistBook, error)
	GetBooksBookId(c *gin.Context, bookId int) (api.ExistBook, error)

	PostUserIdHoarder(c *gin.Context, book api.PostHoarder, userId int) (api.HoarderBook, error)
	PostUserIdHoarderBookId(c *gin.Context, book api.PostHoarder, userId int, bookId int) (api.HoarderBook, error)

	PatchUserIdBooksBookId(c *gin.Context, book api.PostBook, userId int, bookId int) (api.ExistBook, error)
	PatchUserIdHoarderBookId(c *gin.Context, book api.PostHoarder, userId int, bookId int) (api.HoarderBook, error)

	GetUserIdHoarder(c *gin.Context, userId int, params api.GetUserIdHoarderParams) ([]api.HoarderBook, error)

	DeleteUserIdBooksBookId(c *gin.Context, userId int, bookId int) error
	DeleteUserIdHoarderBookId(c *gin.Context, userId int, bookId int) error

	DeleteUserIdTagsTagId(c *gin.Context, userId int, tagId int) error
	GetTags(c *gin.Context, params api.GetTagsParams) ([]api.ExistTag, error)
	PostUserIdTags(c *gin.Context, userId int, taginfo api.TagInfo) (api.ExistTag, error)
}

type Handler struct {
	service ServiceInterface
}

func NewHandler(service ServiceInterface) api.ServerInterface {
	return &Handler{service: service}
}
