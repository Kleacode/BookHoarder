package usecases

import "github.com/lib/pq"

type BookRecord struct {
	Id     int64         `db:"id"`
	Title  string        `db:"title"`
	UserId int64         `db:"user_id"`
	TagIds pq.Int64Array `db:"tags_id"`
}

type HoarderRecord struct {
	UserId   int `db:"user_id"`
	BookId   int `db:"book_id"`
	StatusId int `db:"status_id"`
}
