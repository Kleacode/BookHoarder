package usecases

import (
	api "back/src/generated"

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
		Details: &api.NewBook{
			UserId: &r.UserId,
			BookInfo: &api.BookInfo{
				Title:  &r.Title,
				TagIds: (*[]int64)(&r.TagIds),
			},
		},
	}
}

func (r *BookRecord) ToHoarderBook(StatusId int64) api.HoarderBook {
	return api.HoarderBook{
		StatusId: &StatusId,
		Details: &api.ExistBook{
			BookId: &r.Id,
			Details: &api.NewBook{
				UserId: &r.UserId,
				BookInfo: &api.BookInfo{
					Title:  &r.Title,
					TagIds: (*[]int64)(&r.TagIds),
				},
			},
		},
	}
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

func (r *UserHoarderRecord) ToHoarderBook() api.HoarderBook {
	return api.HoarderBook{
		StatusId: &r.StatusId,
		Details: &api.ExistBook{
			BookId: &r.Id,
			Details: &api.NewBook{
				UserId: &r.UserId,
				BookInfo: &api.BookInfo{
					Title:  &r.Title,
					TagIds: (*[]int64)(&r.TagIds),
				},
			},
		},
	}
}
