package usecases

import (
	api "back/src/generated"
	"back/src/handler"
	"context"
)

type RepositoryInterface interface {
	GetBooks(c context.Context, params api.GetBooksParams) ([]BookRecord, error)
	GetBooksBookId(c context.Context, bookId int) (BookRecord, error)

	PostUserIdHoarder(c context.Context, data BookRecord, statusId int) (api.HoarderBook, error)
	PostUserIdHoarderBookId(c context.Context, data BookRecord, statusId int) (api.HoarderBook, error)

	PatchUserIdBooksBookId(c context.Context, data BookRecord) (api.ExistBook, error)
	PatchUserIdHoarderBookId(c context.Context, data HoarderRecord) (api.HoarderBook, error)

	GetUserIdHoarder(c context.Context, userId int, params api.GetUserIdHoarderParams) ([]UserHoarderRecord, error)

	DeleteUserIdBooksBookId(c context.Context, userId int, bookId int) error
	DeleteUserIdHoarderBookId(c context.Context, userId int, bookId int) error
}

type Service struct {
	repo RepositoryInterface
}

func NewService(repo RepositoryInterface) handler.ServiceInterface {
	return Service{repo: repo}
}
