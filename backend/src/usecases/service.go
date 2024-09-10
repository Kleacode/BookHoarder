package usecases

import (
	"back/src/domain"
	api "back/src/generated"
	"back/src/generated/models"
	"back/src/handler"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null/v8"
)

type RepositoryInterface interface {
	GetBooks(c *gin.Context, params *api.GetBooksParams) ([]models.Book, error)
	GetBooksBookId(c *gin.Context, bookId int) (models.Book, error)
	GetHoarders(c *gin.Context, userId int, params *api.GetUserIdHoarderParams) ([]domain.ExistHoarderRecord, error)
	GetTags(c *gin.Context, params *api.GetTagsParams) ([]models.Tag, error)

	DeleteUserIdBooksBookId(c *gin.Context, userId int, bookId int) error
	DeleteUserIdHoarderHoarderId(c *gin.Context, userId int, hoarderId int) error
	DeleteUserIdTagsTagId(c *gin.Context, userId int, tagId int) error

	InsertTag(c *gin.Context, data *models.Tag) (models.Tag, error)
	InsertBook(c *gin.Context, data *models.Book) (models.Book, error)
	InsertHoarder(c *gin.Context, data *models.UserBookStatus) (models.UserBookStatus, error)
	UpsertHoarderTags(c *gin.Context, data []int, hoarderId int) error
}

type Service struct {
	repo RepositoryInterface
}

// DeleteUserIdBooksBookId implements handler.ServiceInterface.
func (s *Service) DeleteUserIdBooksBookId(c *gin.Context, userId int, bookId int) error {
	return s.repo.DeleteUserIdBooksBookId(c, userId, bookId)
}

// DeleteUserIdHoarderHoarderId implements handler.ServiceInterface.
func (s *Service) DeleteUserIdHoarderHoarderId(c *gin.Context, userId int, hoarderId int) error {
	return s.repo.DeleteUserIdHoarderHoarderId(c, userId, hoarderId)
}

// DeleteUserIdTagsTagId implements handler.ServiceInterface.
func (s *Service) DeleteUserIdTagsTagId(c *gin.Context, userId int, tagId int) error {
	return s.repo.DeleteUserIdTagsTagId(c, userId, tagId)
}

// GetBook implements handler.ServiceInterface.
func (s *Service) GetBook(c *gin.Context, bookID int) (api.ExistBook, error) {
	book, err := s.repo.GetBooksBookId(c, bookID)
	if err != nil {
		return api.ExistBook{}, err
	}
	return api.ExistBook{BookId: &book.ID, Title: &book.Title.String, UserId: &book.UserID}, nil
}

// GetBooks implements handler.ServiceInterface.
func (s *Service) GetBooks(c *gin.Context, params *api.GetBooksParams) ([]api.ExistBook, error) {
	books, err := s.repo.GetBooks(c, params)
	if err != nil {
		return nil, err
	}
	var result []api.ExistBook
	for _, b := range books {
		result = append(result, api.ExistBook{BookId: &b.ID, Title: &b.Title.String, UserId: &b.UserID})
	}
	return result, nil
}

// GetHoarderBooks implements handler.ServiceInterface.
func (s *Service) GetHoarderBooks(c *gin.Context, userId int, params *api.GetUserIdHoarderParams) ([]api.ExistHoarderBook, error) {
	hoarders, err := s.repo.GetHoarders(c, userId, params)
	if err != nil {
		return nil, err
	}
	var result []api.ExistHoarderBook
	for _, h := range hoarders {
		r, err := h.ToExistHoarderBook()
		if err != nil {
			return nil, err
		}
		result = append(result, r)
	}
	return result, nil
}

// GetTags implements handler.ServiceInterface.
func (s *Service) GetTags(c *gin.Context, params *api.GetTagsParams) ([]api.ExistTag, error) {
	tags, err := s.repo.GetTags(c, params)
	if err != nil {
		return nil, err
	}

	var result []api.ExistTag
	for _, t := range tags {
		result = append(result, api.ExistTag{TagId: &t.ID, Name: &t.Name.String, UserId: &t.UserID})
	}
	return result, nil
}

// PatchUserIdBooksBookId implements handler.ServiceInterface.
func (s *Service) PatchUserIdBooksBookId(c *gin.Context, userId int, bookId int, data *api.Book) (api.ExistBook, error) {
	panic("unimplemented")
}

// PatchUserIdHoarderHoarderId implements handler.ServiceInterface.
func (s *Service) PatchUserIdHoarderHoarderId(c *gin.Context, userId int, hoarderId int, data *api.PostHoarderExist) (api.ExistHoarderBook, error) {
	panic("unimplemented")
}

// PostUserIdHoarder implements handler.ServiceInterface.
func (s *Service) PostUserIdHoarder(c *gin.Context, userId int, data *api.PostHoarderNew) (api.ExistHoarderBook, error) {
	var newbook models.Book = models.Book{Title: null.NewString(*data.Book.Title, true), UserID: userId}
	b, err := s.repo.InsertBook(c, &newbook)
	if err != nil {
		return api.ExistHoarderBook{}, err
	}

	result, err := s.PostUserIdHoarderBookId(c, userId, b.ID, &api.PostHoarderExist{Status: data.Status, Tags: data.Tags})
	if err != nil {
		return api.ExistHoarderBook{}, err
	}
	return result, nil
}

// PostUserIdHoarderBookId implements handler.ServiceInterface.
func (s *Service) PostUserIdHoarderBookId(c *gin.Context, userId int, bookId int, data *api.PostHoarderExist) (api.ExistHoarderBook, error) {
	statusID, err := domain.ConvertStautstoId(*data.Status)
	if err != nil {
		return api.ExistHoarderBook{}, err
	}
	var new models.UserBookStatus = models.UserBookStatus{UserID: userId, BookID: bookId, StatusID: statusID}
	result, err := s.repo.InsertHoarder(c, &new)
	if err != nil {
		return api.ExistHoarderBook{}, err
	}

	var tagids []int
	for _, t := range *data.Tags {
		tagids = append(tagids, *t.Id)
	}
	err = s.repo.UpsertHoarderTags(c, tagids, result.ID)
	if err != nil {
		return api.ExistHoarderBook{}, err
	}

	return api.ExistHoarderBook{}, nil
}

// PostUserIdTags implements handler.ServiceInterface.
func (s *Service) PostUserIdTags(c *gin.Context, userId int, data *api.TagInfo) (api.ExistTag, error) {
	tag, err := s.repo.InsertTag(c, &models.Tag{Name: null.NewString(*data.Name, true), UserID: userId})
	if err != nil {
		return api.ExistTag{}, err
	}

	return api.ExistTag{TagId: &tag.ID, Name: &tag.Name.String, UserId: &tag.UserID}, nil
}

func NewService(repo RepositoryInterface) handler.ServiceInterface {
	return &Service{repo: repo}
}
