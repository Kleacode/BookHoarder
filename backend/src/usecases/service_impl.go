package usecases

import (
	api "back/src/generated"

	"github.com/gin-gonic/gin"
)

// DeleteUserIdBooksBookId implements handler.ServiceInterface.
func (s Service) DeleteUserIdBooksBookId(c *gin.Context, userId int, bookId int) error {
	err := s.repo.DeleteUserIdBooksBookId(c, userId, bookId)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUserIdHoarderBookId implements handler.ServiceInterface.
func (s Service) DeleteUserIdHoarderBookId(c *gin.Context, userId int, bookId int) error {
	err := s.repo.DeleteUserIdHoarderBookId(c, userId, bookId)
	if err != nil {
		return err
	}
	return nil
}

// GetBooks implements handler.ServiceInterface.
func (s Service) GetBooks(c *gin.Context, params api.GetBooksParams) ([]api.Book, error) {
	res, err := s.repo.GetBooks(c, params)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetBooksBookId implements handler.ServiceInterface.
func (s Service) GetBooksBookId(c *gin.Context, bookId int) (api.Book, error) {
	res, err := s.repo.GetBooksBookId(c, bookId)
	if err != nil {
		return api.Book{}, err
	}
	return res, nil
}

// GetUserIdHoarder implements handler.ServiceInterface.
func (s Service) GetUserIdHoarder(c *gin.Context, userId int, params api.GetUserIdHoarderParams) ([]api.Book, error) {
	res, err := s.repo.GetUserIdHoarder(c, userId, params)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// PatchUserIdBooksBookId implements handler.ServiceInterface.
func (s Service) PatchUserIdBooksBookId(c *gin.Context, book api.Book) (api.Book, error) {
	res, err := s.repo.PatchUserIdBooksBookId(c, BookRecord{
		Id:     book.BookId,
		UserId: book.UserId,
		Title:  book.Title,
		TagIds: *book.TagIds,
	})
	if err != nil {
		return api.Book{}, err
	}
	return res, nil
}

// PatchUserIdHoarderBookId implements handler.ServiceInterface.
func (s Service) PatchUserIdHoarderBookId(c *gin.Context, book api.Book) (api.Book, error) {
	res, err := s.repo.PatchUserIdHoarderBookId(c, HoarderRecord{
		UserId:   int(book.UserId),
		BookId:   int(book.BookId),
		StatusId: 2,
	})
	if err != nil {
		return api.Book{}, err
	}
	return res, nil
}

// PostUserIdHoarder implements handler.ServiceInterface.
func (s Service) PostUserIdHoarder(c *gin.Context, book api.Book) (api.Book, error) {
	res, err := s.repo.PostUserIdHoarder(c, BookRecord{
		Id:     book.BookId,
		UserId: book.UserId,
		Title:  book.Title,
		TagIds: *book.TagIds,
	}, 2)
	if err != nil {
		return api.Book{}, err
	}
	return res, nil
}

// PostUserIdHoarderBookId implements handler.ServiceInterface.
func (s Service) PostUserIdHoarderBookId(c *gin.Context, book api.Book) (api.Book, error) {
	res, err := s.repo.PostUserIdHoarderBookId(c, BookRecord{
		Id:     book.BookId,
		UserId: book.UserId,
		Title:  book.Title,
		TagIds: *book.TagIds,
	}, 2)
	if err != nil {
		return api.Book{}, err
	}
	return res, nil
}
