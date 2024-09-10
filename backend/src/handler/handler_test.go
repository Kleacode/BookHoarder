package handler

import (
	api "back/src/generated"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type MockUsecase struct {
	mock.Mock
}

// DeleteUserIdBooksBookId implements ServiceInterface.
func (m *MockUsecase) DeleteUserIdBooksBookId(c *gin.Context, userId int, bookId int) error {
	panic("unimplemented")
}

// DeleteUserIdHoarderHoarderId implements ServiceInterface.
func (m *MockUsecase) DeleteUserIdHoarderHoarderId(c *gin.Context, userId int, hoarderId int) error {
	panic("unimplemented")
}

// DeleteUserIdTagsTagId implements ServiceInterface.
func (m *MockUsecase) DeleteUserIdTagsTagId(c *gin.Context, userId int, tagId int) error {
	panic("unimplemented")
}

// GetBook implements ServiceInterface.
func (m *MockUsecase) GetBook(c *gin.Context, bookID int) (api.ExistBook, error) {
	panic("unimplemented")
}

// GetBooks implements ServiceInterface.
func (m *MockUsecase) GetBooks(c *gin.Context, params *api.GetBooksParams) ([]api.ExistBook, error) {
	panic("unimplemented")
}

// GetHoarderBooks implements ServiceInterface.
func (m *MockUsecase) GetHoarderBooks(c *gin.Context, userId int, params *api.GetUserIdHoarderParams) ([]api.ExistHoarderBook, error) {
	panic("unimplemented")
}

// GetTags implements ServiceInterface.
func (m *MockUsecase) GetTags(c *gin.Context, params *api.GetTagsParams) ([]api.ExistTag, error) {
	panic("unimplemented")
}

// PatchUserIdBooksBookId implements ServiceInterface.
func (m *MockUsecase) PatchUserIdBooksBookId(c *gin.Context, userId int, bookId int, data *api.Book) (api.ExistBook, error) {
	panic("unimplemented")
}

// PatchUserIdHoarderHoarderId implements ServiceInterface.
func (m *MockUsecase) PatchUserIdHoarderHoarderId(c *gin.Context, userId int, hoarderId int, data *api.PostHoarderExist) (api.ExistHoarderBook, error) {
	panic("unimplemented")
}

// PostUserIdHoarder implements ServiceInterface.
func (m *MockUsecase) PostUserIdHoarder(c *gin.Context, userId int, data *api.PostHoarderNew) (api.ExistHoarderBook, error) {
	panic("unimplemented")
}

// PostUserIdHoarderBookId implements ServiceInterface.
func (m *MockUsecase) PostUserIdHoarderBookId(c *gin.Context, userId int, bookId int, data *api.PostHoarderExist) (api.ExistHoarderBook, error) {
	panic("unimplemented")
}

// PostUserIdTags implements ServiceInterface.
func (m *MockUsecase) PostUserIdTags(c *gin.Context, userId int, data *api.TagInfo) (api.ExistTag, error) {
	panic("unimplemented")
}

var _ ServiceInterface = (*MockUsecase)(nil)
