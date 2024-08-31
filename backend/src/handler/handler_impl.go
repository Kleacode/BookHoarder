package handler

import (
	"back/src/domain"
	api "back/src/generated"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteUserIdBooksBookId implements api.ServerInterface.
func (h *Handler) DeleteUserIdBooksBookId(c *gin.Context, userId int, bookId int) {
	err := h.service.DeleteUserIdBooksBookId(c, userId, bookId)
	if err != nil {
		c.JSON(domain.GetStatusCode(err), err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// DeleteUserIdHoarderBookId implements api.ServerInterface.
func (h *Handler) DeleteUserIdHoarderBookId(c *gin.Context, userId int, bookId int) {
	err := h.service.DeleteUserIdHoarderBookId(c, userId, bookId)
	if err != nil {
		c.JSON(domain.GetStatusCode(err), err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// GetBooks implements api.ServerInterface.
func (h *Handler) GetBooks(c *gin.Context, params api.GetBooksParams) {
	books, err := h.service.GetBooks(c, params)
	if err != nil {
		c.JSON(domain.GetStatusCode(err), err.Error())
		return
	}
	c.JSON(http.StatusOK, books)
}

// GetBooksBookId implements api.ServerInterface.
func (h *Handler) GetBooksBookId(c *gin.Context, bookId int) {
	book, err := h.service.GetBooksBookId(c, bookId)
	if err != nil {
		c.JSON(domain.GetStatusCode(err), err.Error())
		return
	}
	c.JSON(http.StatusOK, book)
}

// GetUserIdHoarder implements api.ServerInterface.
func (h *Handler) GetUserIdHoarder(c *gin.Context, userId int, params api.GetUserIdHoarderParams) {
	books, err := h.service.GetUserIdHoarder(c, userId, params)
	if err != nil {
		c.JSON(domain.GetStatusCode(err), err.Error())
		return
	}
	c.JSON(http.StatusOK, books)
}

// PatchUserIdBooksBookId implements api.ServerInterface.
func (h *Handler) PatchUserIdBooksBookId(c *gin.Context, userId int, bookId int) {
	var book api.PostBook
	if err := c.BindJSON(&book); err != nil {
		c.JSON(domain.GetStatusCode(err), err.Error())
		return
	}

	patch, err := h.service.PatchUserIdBooksBookId(c, book, userId, bookId)

	if err != nil {
		c.JSON(domain.GetStatusCode(err), err.Error())
		return
	}

	c.JSON(http.StatusOK, patch)
}

// PatchUserIdHoarderBookId implements api.ServerInterface.
func (h *Handler) PatchUserIdHoarderBookId(c *gin.Context, userId int, bookId int) {
	var book api.PostHoarder
	if err := c.BindJSON(&book); err != nil {
		return
	}

	patch, err := h.service.PatchUserIdHoarderBookId(c, book, userId, bookId)

	if err != nil {
		c.JSON(domain.GetStatusCode(err), err.Error())
		return
	}

	c.JSON(http.StatusOK, patch)
}

// PostUserIdHoarder implements api.ServerInterface.
func (h *Handler) PostUserIdHoarder(c *gin.Context, userId int) {
	var book api.PostHoarder
	if err := c.BindJSON(&book); err != nil {
		return
	}

	post, err := h.service.PostUserIdHoarder(c, book, userId)
	if err != nil {
		c.JSON(domain.GetStatusCode(err), err.Error())
		return
	}

	c.JSON(http.StatusCreated, post)
}

// PostUserIdHoarderBookId implements api.ServerInterface.
func (h *Handler) PostUserIdHoarderBookId(c *gin.Context, userId int, bookId int) {
	var book api.PostHoarder
	if err := c.BindJSON(&book); err != nil {
		return
	}

	post, err := h.service.PostUserIdHoarderBookId(c, book, userId, bookId)
	if err != nil {
		c.JSON(domain.GetStatusCode(err), err.Error())
		return
	}

	c.JSON(http.StatusCreated, post)
}

// DeleteUserIdTagsTagId implements api.ServerInterface.
func (h *Handler) DeleteUserIdTagsTagId(c *gin.Context, userId int, tagId int) {
	err := h.service.DeleteUserIdTagsTagId(c, userId, tagId)
	if err != nil {
		c.JSON(domain.GetStatusCode(err), err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// GetTags implements api.ServerInterface.
func (h *Handler) GetTags(c *gin.Context, params api.GetTagsParams) {
	tags, err := h.service.GetTags(c, params)
	if err != nil {
		c.JSON(domain.GetStatusCode(err), err.Error())
		return
	}
	c.JSON(http.StatusOK, tags)
}

// PostUserIdTags implements api.ServerInterface.
func (h *Handler) PostUserIdTags(c *gin.Context, userId int) {
	var taginfo api.TagInfo
	if err := c.BindJSON(&taginfo); err != nil {
		return
	}

	tag, err := h.service.PostUserIdTags(c, userId, taginfo)
	if err != nil {
		c.JSON(domain.GetStatusCode(err), err.Error())
		return
	}

	c.JSON(http.StatusCreated, tag)
}
