package handler

import (
	api "back/src/generated"

	"github.com/gin-gonic/gin"
)

type ServiceInterface interface {
	GetBooks(c *gin.Context, params api.GetBooksParams) (api.Books, error)
	GetBooksBookId(c *gin.Context, bookId int) (api.Book, error)

	PostUserIdHoarder(c *gin.Context, book api.Book) (api.Book, error)
	PostUserIdHoarderBookId(c *gin.Context, book api.Book) (api.Book, error)

	PatchUserIdBooksBookId(c *gin.Context, book api.Book) (api.Book, error)
	PatchUserIdHoarderBookId(c *gin.Context, book api.Book) (api.Book, error)

	GetUserIdHoarder(c *gin.Context, userId int, params api.GetUserIdHoarderParams) (api.Books, error)

	DeleteUserIdBooksBookId(c *gin.Context, userId int, bookId int) error
	DeleteUserIdHoarderBookId(c *gin.Context, userId int, bookId int) error
}

type Handler struct {
	service ServiceInterface
}

func NewHandler(service ServiceInterface) api.ServerInterface {
	return &Handler{service: service}
}
