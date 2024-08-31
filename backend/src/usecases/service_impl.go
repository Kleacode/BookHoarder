package usecases

import (
	api "back/src/generated"

	"github.com/gin-gonic/gin"
)

// DeleteUserIdBooksBookId implements handler.ServiceInterface.
func (s Service) DeleteUserIdBooksBookId(c *gin.Context, userId int, bookId int) error {
	err := s.repo.DeleteUserIdBooksBookId(c, userId, bookId)
	return err
}

// DeleteUserIdHoarderBookId implements handler.ServiceInterface.
func (s Service) DeleteUserIdHoarderBookId(c *gin.Context, userId int, bookId int) error {
	err := s.repo.DeleteUserIdHoarderBookId(c, userId, bookId)
	return err
}

// GetBooks implements handler.ServiceInterface.
func (s Service) GetBooks(c *gin.Context, params api.GetBooksParams) ([]api.ExistBook, error) {
	books, err := s.repo.GetBooks(c, params)
	if err != nil {
		return nil, err
	}

	var result []api.ExistBook
	for _, e := range books {
		result = append(result, e.ToExistBook())
	}
	return result, err
}

// GetBooksBookId implements handler.ServiceInterface.
func (s Service) GetBooksBookId(c *gin.Context, bookId int) (api.ExistBook, error) {
	book, err := s.repo.GetBooksBookId(c, bookId)
	if err != nil {
		return api.ExistBook{}, err
	}
	return book.ToExistBook(), nil
}

// GetUserIdHoarder implements handler.ServiceInterface.
func (s Service) GetUserIdHoarder(c *gin.Context, userId int, params api.GetUserIdHoarderParams) ([]api.HoarderBook, error) {
	books, err := s.repo.GetUserIdHoarder(c, userId, params)
	if err != nil {
		return nil, err
	}

	var result []api.HoarderBook
	for _, e := range books {
		r, err := e.ToHoarderBook()
		if err != nil {
			return nil, err
		}
		result = append(result, r)
	}
	return result, nil
}

// PatchUserIdBooksBookId implements handler.ServiceInterface.
func (s Service) PatchUserIdBooksBookId(c *gin.Context, book api.PostBook, userId int, bookId int) (api.ExistBook, error) {
	result, err := s.repo.PatchUserIdBooksBookId(c, BookRecord{
		Id:     int64(bookId),
		UserId: int64(userId),
		Title:  *book.Title,
		TagIds: *book.TagIds,
	})

	if err != nil {
		return api.ExistBook{}, err
	}
	return result, err
}

// PatchUserIdHoarderBookId implements handler.ServiceInterface.
func (s Service) PatchUserIdHoarderBookId(c *gin.Context, book api.PostHoarder, userId int, bookId int) (api.HoarderBook, error) {
	statusId, err := ConvertStautstoId(*book.Status)
	if err != nil {
		return api.HoarderBook{}, err
	}

	result, err := s.repo.PatchUserIdHoarderBookId(c, HoarderRecord{
		UserId:   userId,
		BookId:   bookId,
		StatusId: statusId,
	})
	if err != nil {
		return api.HoarderBook{}, err
	}
	return result, nil
}

// PostUserIdHoarder implements handler.ServiceInterface.
func (s Service) PostUserIdHoarder(c *gin.Context, book api.PostHoarder, userId int) (api.HoarderBook, error) {
	statusId, err := ConvertStautstoId(*book.Status)
	if err != nil {
		return api.HoarderBook{}, err
	}

	result, err := s.repo.PostUserIdHoarder(c, BookRecord{
		//Id: unused
		Title:  *book.Title,
		UserId: int64(userId),
		TagIds: *book.TagIds,
	}, statusId)
	if err != nil {
		return api.HoarderBook{}, err
	}
	return result, nil
}

// PostUserIdHoarderBookId implements handler.ServiceInterface.
func (s Service) PostUserIdHoarderBookId(c *gin.Context, book api.PostHoarder, userId int, bookId int) (api.HoarderBook, error) {
	statusId, err := ConvertStautstoId(*book.Status)
	if err != nil {
		return api.HoarderBook{}, err
	}

	result, err := s.repo.PostUserIdHoarderBookId(c, BookRecord{
		Id:     int64(bookId),
		Title:  *book.Title,
		TagIds: *book.TagIds,
		UserId: int64(userId),
	}, statusId)
	if err != nil {
		return api.HoarderBook{}, err
	}
	return result, nil
}

// DeleteUserIdTagsTagId implements handler.ServiceInterface.
func (s Service) DeleteUserIdTagsTagId(c *gin.Context, userId int, tagId int) error {
	err := s.repo.DeleteUserIdTagsTagId(c, userId, tagId)
	return err
}

// GetTags implements handler.ServiceInterface.
func (s Service) GetTags(c *gin.Context, params api.GetTagsParams) ([]api.ExistTag, error) {
	tags, err := s.repo.GetTags(c, params)
	if err != nil {
		return nil, err
	}

	var result []api.ExistTag
	for _, e := range tags {
		result = append(result, e.ToExistTag())
	}
	return result, nil
}

// PostUserIdTags implements handler.ServiceInterface.
func (s Service) PostUserIdTags(c *gin.Context, userId int, taginfo api.TagInfo) (api.ExistTag, error) {
	tag, err := s.repo.PostUserIdTags(c, TagRecord{
		// Id: unused,
		Name:   *taginfo.Name,
		UserId: int64(userId),
	})
	if err != nil {
		return api.ExistTag{}, err
	}

	return tag, nil
}
