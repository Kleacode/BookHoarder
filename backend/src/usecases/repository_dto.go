package usecases

import (
	api "back/src/generated"
	"errors"

	"github.com/lib/pq"
)

type BookRecord struct {
	Id     int64         `db:"id"`
	Title  string        `db:"title"`
	UserId int64         `db:"user_id"`
	TagIds pq.Int64Array `db:"tags_id"`
}

func (r *BookRecord) ToExistBook() api.ExistBook {
	return api.ExistBook{
		BookId: &r.Id,
		UserId: &r.UserId,
		Title:  &r.Title,
		TagIds: (*[]int64)(&r.TagIds),
	}
}

func (r *BookRecord) ToHoarderBook(StatusId int64) (api.HoarderBook, error) {
	status, err := ConvertIdtoStatus(StatusId)
	if err != nil {
		return api.HoarderBook{}, err
	}

	return api.HoarderBook{
		Status: &status,
		BookId: &r.Id,
		UserId: &r.UserId,
		Title:  &r.Title,
		TagIds: (*[]int64)(&r.TagIds),
	}, nil
}

type HoarderRecord struct {
	UserId   int `db:"user_id"`
	BookId   int `db:"book_id"`
	StatusId int `db:"status_id"`
}

type UserHoarderRecord struct {
	Id       int64         `db:"id"`
	Title    string        `db:"title"`
	UserId   int64         `db:"user_id"`
	TagIds   pq.Int64Array `db:"tags_id"`
	StatusId int64         `db:"status_id"`
}

func (r *UserHoarderRecord) ToHoarderBook() (api.HoarderBook, error) {
	status, err := ConvertIdtoStatus(r.StatusId)
	if err != nil {
		return api.HoarderBook{}, err
	}

	return api.HoarderBook{
		Status: &status,
		BookId: &r.Id,
		UserId: &r.UserId,
		Title:  &r.Title,
		TagIds: (*[]int64)(&r.TagIds),
	}, nil
}

type TagRecord struct {
	Id     int64  `db:"id"`
	Name   string `db:"name"`
	UserId int64  `db:"user_id"`
}

func (r *TagRecord) ToExistTag() api.ExistTag {
	return api.ExistTag{
		TagId:  &r.Id,
		Name:   &r.Name,
		UserId: &r.UserId,
	}
}

func ConvertStautstoId(s api.Status) (int, error) {
	switch s {
	case api.Todo:
		return 1, nil
	case api.Wip:
		return 2, nil
	case api.Done:
		return 3, nil
	default:
		return -1, errors.New("invalid status")
	}
}

func ConvertIdtoStatus(i int64) (api.Status, error) {
	switch i {
	case 1:
		return api.Todo, nil
	case 2:
		return api.Wip, nil
	case 3:
		return api.Done, nil
	default:
		return "", errors.New("invalid status")
	}
}
