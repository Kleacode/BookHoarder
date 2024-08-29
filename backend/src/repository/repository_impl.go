package repository

import (
	"back/src/domain"
	api "back/src/generated"
	"back/src/usecases"
	"context"

	"github.com/lib/pq"
)

// DeleteUserIdBooksBookId implements usecases.RepositoryInterface.
func (r Repository) DeleteUserIdBooksBookId(c context.Context, userId int, bookId int) error {
	result, err := r.db.Exec("DELETE FROM books WHERE id = $1 AND user_id = $2", bookId, userId)
	if err != nil {
		return err
	}
	row, _ := result.RowsAffected()
	if row == 0 {
		return domain.ErrorNotFound
	}
	return nil
}

// DeleteUserIdHoarderBookId implements usecases.RepositoryInterface.
func (r Repository) DeleteUserIdHoarderBookId(c context.Context, userId int, bookId int) error {
	result, err := r.db.Exec("DELETE FROM user_book_status WHERE book_id = $1 AND user_id = $2", bookId, userId)
	if err != nil {
		return err
	}
	row, _ := result.RowsAffected()
	if row == 0 {
		return domain.ErrorNotFound
	}
	return nil
}

// GetBooks implements usecases.RepositoryInterface.
func (r Repository) GetBooks(c context.Context, params api.GetBooksParams) ([]api.Book, error) {
	var books []usecases.BookRecord
	var result []api.Book

	err := r.db.Select(&books, "SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	for _, e := range books {
		result = append(result, api.Book{BookId: e.Id, UserId: e.UserId, Title: e.Title, TagIds: (*[]int64)(&e.TagIds)})
	}
	return result, err
}

// GetBooksBookId implements usecases.RepositoryInterface.
func (r Repository) GetBooksBookId(c context.Context, bookId int) (api.Book, error) {
	var book usecases.BookRecord
	err := r.db.Get(&book, "SELECT * FROM books WHERE id = $1", bookId)
	if err != nil {
		return api.Book{}, err
	}
	return api.Book{BookId: int64(bookId), Title: book.Title, TagIds: (*[]int64)(&book.TagIds)}, nil
}

// GetUserIdHoarder implements usecases.RepositoryInterface.
func (r Repository) GetUserIdHoarder(c context.Context, userId int, params api.GetUserIdHoarderParams) ([]api.Book, error) {
	var books []usecases.BookRecord
	err := r.db.Select(&books, `SELECT * FROM books WHERE id IN (SELECT book_id FROM user_book_status WHERE user_id = $1)`, userId)
	if err != nil {
		return nil, err
	}
	var result []api.Book
	for _, e := range books {
		result = append(result, api.Book{
			BookId: int64(e.Id),
			Title:  e.Title,
			TagIds: (*[]int64)(&e.TagIds),
		})
	}
	return result, err
}

// PatchUserIdBooksBookId implements usecases.RepositoryInterface.
func (r Repository) PatchUserIdBooksBookId(c context.Context, data usecases.BookRecord) (api.Book, error) {
	result, err := r.db.NamedExec("UPDATE books SET title=:title, tags_id=:tags_id WHERE id=:id AND user_id=:user_id ", data)
	if err != nil {
		return api.Book{}, err
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		return api.Book{}, domain.ErrorNotFound
	}
	return api.Book{}, nil
}

// PatchUserIdHoarderBookId implements usecases.RepositoryInterface.
func (r Repository) PatchUserIdHoarderBookId(c context.Context, data usecases.HoarderRecord) (api.Book, error) {
	result, err := r.db.NamedExec("UPDATE user_book_status SET status_id=:status_id WHERE user_id=:user_id AND book_id=:book_id", data)
	if err != nil {
		return api.Book{}, err
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		return api.Book{}, domain.ErrorNotFound
	}
	return api.Book{}, nil
}

// PostUserIdHoarder implements usecases.RepositoryInterface.
func (r Repository) PostUserIdHoarder(c context.Context, data usecases.BookRecord, statusId int) (api.Book, error) {
	var createdId int
	err := r.db.Get(&createdId, "INSERT INTO books (title, user_id, tags_id) VALUES ($1, $2, $3) RETURNING id", data.Title, data.UserId, pq.Array(data.TagIds))
	if err != nil {
		return api.Book{}, err
	}

	if createdId == 0 {
		return api.Book{}, domain.ErrorFailedInsert
	}

	data.Id = int64(createdId)
	_, err = r.PostUserIdHoarderBookId(c, data, statusId)
	if err != nil {
		return api.Book{}, err
	}
	return api.Book{}, nil
}

// PostUserIdHoarderBookId implements usecases.RepositoryInterface.
func (r Repository) PostUserIdHoarderBookId(c context.Context, data usecases.BookRecord, statusId int) (api.Book, error) {
	var createdId int
	err := r.db.Get(&createdId, "INSERT INTO user_book_status (user_id, book_id, status_id) VALUES ($1, $2, $3) RETURNING id", data.UserId, data.Id, statusId)
	if err != nil {
		return api.Book{}, err
	}

	if createdId == 0 {
		return api.Book{}, domain.ErrorFailedInsert
	}

	return api.Book{}, nil
}
