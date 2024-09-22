package usecases

import (
	"back/src/domain"
	api "back/src/generated"
	"back/src/generated/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

// DeleteUserIdBooksBookId implements RepositoryInterface.
func (m *MockRepository) DeleteUserIdBooksBookId(c *gin.Context, userId int, bookId int) error {
	panic("unimplemented")
}

// DeleteUserIdHoarderHoarderId implements RepositoryInterface.
func (m *MockRepository) DeleteUserIdHoarderHoarderId(c *gin.Context, userId int, hoarderId int) error {
	panic("unimplemented")
}

// DeleteUserIdTagsTagId implements RepositoryInterface.
func (m *MockRepository) DeleteUserIdTagsTagId(c *gin.Context, userId int, tagId int) error {
	panic("unimplemented")
}

// GetBooks implements RepositoryInterface.
func (m *MockRepository) GetBooks(c *gin.Context, params *api.GetBooksParams) ([]models.Book, error) {
	panic("unimplemented")
}

// GetBooksBookId implements RepositoryInterface.
func (m *MockRepository) GetBooksBookId(c *gin.Context, bookId int) (models.Book, error) {
	panic("unimplemented")
}

// GetHoarders implements RepositoryInterface.
func (m *MockRepository) GetHoarders(c *gin.Context, userId int, params *api.GetUserIdHoarderParams) ([]domain.ExistHoarderRecord, error) {
	panic("unimplemented")
}

// GetUserTags implements RepositoryInterface.
func (m *MockRepository) GetUserTags(c *gin.Context, userId int, params *api.GetUserIdTagsParams) ([]models.Tag, error) {
	panic("unimplemented")
}

// InsertBook implements RepositoryInterface.
func (m *MockRepository) InsertBook(c *gin.Context, data *models.Book) (models.Book, error) {
	panic("unimplemented")
}

// InsertHoarder implements RepositoryInterface.
func (m *MockRepository) InsertHoarder(c *gin.Context, data *models.UserBookStatus) (models.UserBookStatus, error) {
	panic("unimplemented")
}

// InsertTag implements RepositoryInterface.
func (m *MockRepository) InsertTag(c *gin.Context, data *models.Tag) (models.Tag, error) {
	panic("unimplemented")
}

// UpdateBook implements RepositoryInterface.
func (m *MockRepository) UpdateBook(c *gin.Context, data *models.Book) (models.Book, error) {
	panic("unimplemented")
}

// UpsertHoarderTags implements RepositoryInterface.
func (m *MockRepository) UpsertHoarderTags(c *gin.Context, data []int, hoarderId int) error {
	panic("unimplemented")
}

var _ RepositoryInterface = (*MockRepository)(nil)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	m.Run()
}
