package usecases

import (
	api "back/src/generated"
	"back/src/handler"
	"context"
)

type RepositoryInterface interface {
	GetBooks(c context.Context, params api.GetBooksParams) (api.Books, error)
	GetBooksBookId(c context.Context, bookId int) (api.Book, error)

	PostUserIdHoarder(c context.Context, data BookRecord, statusId int) (api.Book, error)
	PostUserIdHoarderBookId(c context.Context, data BookRecord, statusId int) (api.Book, error)

	PatchUserIdBooksBookId(c context.Context, data BookRecord) (api.Book, error)
	PatchUserIdHoarderBookId(c context.Context, data HoarderRecord) (api.Book, error)

	GetUserIdHoarder(c context.Context, userId int, params api.GetUserIdHoarderParams) (api.Books, error)

	DeleteUserIdBooksBookId(c context.Context, userId int, bookId int) error
	DeleteUserIdHoarderBookId(c context.Context, userId int, bookId int) error
}

type Service struct {
	repo RepositoryInterface
}

func NewService(repo RepositoryInterface) handler.ServiceInterface {
	return Service{repo: repo}
}
