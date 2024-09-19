package handler

import (
	"back/src/domain"
	api "back/src/generated"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServiceInterface interface {
	GetBooks(c *gin.Context, params *api.GetBooksParams) ([]api.ExistBook, error)
	GetBook(c *gin.Context, bookID int) (api.ExistBook, error)
	GetUserTags(c *gin.Context, userID int, params *api.GetUserIdTagsParams) ([]api.ExistTag, error)
	GetHoarderBooks(c *gin.Context, userId int, params *api.GetUserIdHoarderParams) ([]api.ExistHoarderBook, error)

	DeleteUserIdBooksBookId(c *gin.Context, userId int, bookId int) error
	DeleteUserIdHoarderHoarderId(c *gin.Context, userId int, hoarderId int) error
	DeleteUserIdTagsTagId(c *gin.Context, userId int, tagId int) error

	PatchUserIdBooksBookId(c *gin.Context, userId int, bookId int, data *api.PostBook) (api.ExistBook, error)
	PatchUserIdHoarderHoarderId(c *gin.Context, userId int, hoarderId int, data *api.PostHoarderExist) (api.ExistHoarderBook, error)

	PostUserIdHoarder(c *gin.Context, userId int, data *api.PostHoarderNew) (api.ExistHoarderBook, error)
	PostUserIdHoarderBookId(c *gin.Context, userId int, bookId int, data *api.PostHoarderExist) (api.ExistHoarderBook, error)
	PostUserIdTags(c *gin.Context, userId int, data *api.TagInfo) (api.ExistTag, error)
}

type Handler struct {
	service ServiceInterface
}

// GetUserIdTags implements api.ServerInterface.
func (h *Handler) GetUserIdTags(c *gin.Context, userId int, params api.GetUserIdTagsParams) {
	tags, err := h.service.GetUserTags(c, userId, &params)
	if err != nil {
		c.JSON(domain.GetErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, tags)
}

// DeleteUserIdBooksBookId implements api.ServerInterface.
func (h *Handler) DeleteUserIdBooksBookId(c *gin.Context, userId int, bookId int) {
	err := h.service.DeleteUserIdBooksBookId(c, userId, bookId)
	if err != nil {
		c.JSON(domain.GetErrorResponse(err))
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// DeleteUserIdHoarderHoarderId implements api.ServerInterface.
func (h *Handler) DeleteUserIdHoarderHoarderId(c *gin.Context, userId int, hoarderId int) {
	err := h.service.DeleteUserIdBooksBookId(c, userId, hoarderId)
	if err != nil {
		c.JSON(domain.GetErrorResponse(err))
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// DeleteUserIdTagsTagId implements api.ServerInterface.
func (h *Handler) DeleteUserIdTagsTagId(c *gin.Context, userId int, tagId int) {
	err := h.service.DeleteUserIdTagsTagId(c, userId, tagId)
	if err != nil {
		c.JSON(domain.GetErrorResponse(err))
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// GetBooks implements api.ServerInterface.
func (h *Handler) GetBooks(c *gin.Context, params api.GetBooksParams) {
	books, err := h.service.GetBooks(c, &params)
	if err != nil {
		c.JSON(domain.GetErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, books)
}

// GetBooksBookId implements api.ServerInterface.
func (h *Handler) GetBooksBookId(c *gin.Context, bookId int) {
	book, err := h.service.GetBook(c, bookId)
	if err != nil {
		c.JSON(domain.GetErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, book)
}

// GetUserIdHoarder implements api.ServerInterface.
func (h *Handler) GetUserIdHoarder(c *gin.Context, userId int, params api.GetUserIdHoarderParams) {
	books, err := h.service.GetHoarderBooks(c, userId, &params)
	if err != nil {
		c.JSON(domain.GetErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, books)
}

// PatchUserIdBooksBookId implements api.ServerInterface.
func (h *Handler) PatchUserIdBooksBookId(c *gin.Context, userId int, bookId int) {
	var data api.PostBook
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(domain.GetErrorResponse(err))
		return
	}

	result, err := h.service.PatchUserIdBooksBookId(c, userId, bookId, &data)
	if err != nil {
		c.JSON(domain.GetErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, result)
}

// PatchUserIdHoarderHoarderId implements api.ServerInterface.
func (h *Handler) PatchUserIdHoarderHoarderId(c *gin.Context, userId int, hoarderId int) {
	var data api.PostHoarderExist
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(domain.GetErrorResponse(err))
		return
	}

	result, err := h.service.PatchUserIdHoarderHoarderId(c, userId, hoarderId, &data)
	if err != nil {
		c.JSON(domain.GetErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, result)
}

// PostUserIdHoarder implements api.ServerInterface.
func (h *Handler) PostUserIdHoarder(c *gin.Context, userId int) {
	var data api.PostHoarderNew
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(domain.GetErrorResponse(err))
		return
	}

	result, err := h.service.PostUserIdHoarder(c, userId, &data)
	if err != nil {
		c.JSON(domain.GetErrorResponse(err))
		return
	}
	c.JSON(http.StatusCreated, result)
}

// PostUserIdHoarderBookId implements api.ServerInterface.
func (h *Handler) PostUserIdHoarderBookId(c *gin.Context, userId int, bookId int) {
	var data api.PostHoarderExist
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(domain.GetErrorResponse(err))
		return
	}

	result, err := h.service.PostUserIdHoarderBookId(c, userId, bookId, &data)
	if err != nil {
		c.JSON(domain.GetErrorResponse(err))
		return
	}
	c.JSON(http.StatusCreated, result)
}

// PostUserIdTags implements api.ServerInterface.
func (h *Handler) PostUserIdTags(c *gin.Context, userId int) {
	var data api.TagInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(domain.GetErrorResponse(err))
		return
	}

	result, err := h.service.PostUserIdTags(c, userId, &data)
	if err != nil {
		c.JSON(domain.GetErrorResponse(err))
		return
	}
	c.JSON(http.StatusCreated, result)
}

func NewHandler(service ServiceInterface) api.ServerInterface {
	return &Handler{service: service}
}
