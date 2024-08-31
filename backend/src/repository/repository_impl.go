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
func (r Repository) GetBooks(c context.Context, params api.GetBooksParams) ([]usecases.BookRecord, error) {
	var books []usecases.BookRecord
	err := r.db.Select(&books, "SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	return books, err
}

// GetBooksBookId implements usecases.RepositoryInterface.
func (r Repository) GetBooksBookId(c context.Context, bookId int) (usecases.BookRecord, error) {
	var book usecases.BookRecord
	err := r.db.Get(&book, "SELECT * FROM books WHERE id = $1", bookId)
	if err != nil {
		return usecases.BookRecord{}, err
	}
	return book, nil
}

// GetUserIdHoarder implements usecases.RepositoryInterface.
func (r Repository) GetUserIdHoarder(c context.Context, userId int, params api.GetUserIdHoarderParams) ([]usecases.UserHoarderRecord, error) {
	var books []usecases.UserHoarderRecord
	err := r.db.Select(&books, `SELECT books.*, user_book_status.status_id FROM books 
	INNER JOIN user_book_status
	ON books.id = user_book_status.book_id
	WHERE books.user_id = $1`, userId)
	if err != nil {
		return nil, err
	}
	return books, err
}

// PatchUserIdBooksBookId implements usecases.RepositoryInterface.
func (r Repository) PatchUserIdBooksBookId(c context.Context, data usecases.BookRecord) (api.ExistBook, error) {
	result, err := r.db.NamedExec("UPDATE books SET title=:title, tags_id=:tags_id WHERE id=:id AND user_id=:user_id ", data)
	if err != nil {
		return api.ExistBook{}, err
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		return api.ExistBook{}, domain.ErrorNotFound
	}
	return data.ToExistBook(), nil
}

// PatchUserIdHoarderBookId implements usecases.RepositoryInterface.
func (r Repository) PatchUserIdHoarderBookId(c context.Context, data usecases.HoarderRecord) (api.HoarderBook, error) {
	result, err := r.db.NamedExec("UPDATE user_book_status SET status_id=:status_id WHERE user_id=:user_id AND book_id=:book_id", data)
	if err != nil {
		return api.HoarderBook{}, err
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		return api.HoarderBook{}, domain.ErrorNotFound
	}
	// TODO
	return api.HoarderBook{}, nil
}

// PostUserIdHoarder implements usecases.RepositoryInterface.
func (r Repository) PostUserIdHoarder(c context.Context, data usecases.BookRecord, statusId int) (api.HoarderBook, error) {
	var createdId int
	err := r.db.Get(&createdId, "INSERT INTO books (title, user_id, tags_id) VALUES ($1, $2, $3) RETURNING id", data.Title, data.UserId, pq.Array(data.TagIds))
	if err != nil {
		return api.HoarderBook{}, err
	}

	if createdId == 0 {
		return api.HoarderBook{}, domain.ErrorFailedInsert
	}

	data.Id = int64(createdId)
	_, err = r.PostUserIdHoarderBookId(c, data, statusId)
	if err != nil {
		return api.HoarderBook{}, err
	}

	return data.ToHoarderBook(int64(statusId))
}

// PostUserIdHoarderBookId implements usecases.RepositoryInterface.
func (r Repository) PostUserIdHoarderBookId(c context.Context, data usecases.BookRecord, statusId int) (api.HoarderBook, error) {
	var createdId int
	err := r.db.Get(&createdId, "INSERT INTO user_book_status (user_id, book_id, status_id) VALUES ($1, $2, $3) RETURNING id", data.UserId, data.Id, statusId)
	if err != nil {
		return api.HoarderBook{}, err
	}

	if createdId == 0 {
		return api.HoarderBook{}, domain.ErrorFailedInsert
	}

	return data.ToHoarderBook(int64(statusId))
}

// DeleteUserIdTagsTagId implements usecases.RepositoryInterface.
func (r Repository) DeleteUserIdTagsTagId(c context.Context, userId int, tagId int) error {
	result, err := r.db.Exec("DELETE FROM tags WHERE id = $1 AND user_id = $2", tagId, userId)
	if err != nil {
		return err
	}
	row, _ := result.RowsAffected()
	if row == 0 {
		return domain.ErrorNotFound
	}
	return nil
}

// GetTags implements usecases.RepositoryInterface.
func (r Repository) GetTags(c context.Context, params api.GetTagsParams) ([]usecases.TagRecord, error) {
	var tags []usecases.TagRecord
	err := r.db.Select(&tags, "SELECT * FROM tags")
	if err != nil {
		return nil, err
	}
	return tags, err
}

// PostUserIdTags implements usecases.RepositoryInterface.
func (r Repository) PostUserIdTags(c context.Context, data usecases.TagRecord) (api.ExistTag, error) {
	var createdId int
	err := r.db.Get(&createdId, "INSERT INTO tags (user_id, name) VALUES ($1, $2) RETURNING id", data.UserId, data.Name)
	if err != nil {
		return api.ExistTag{}, err
	}

	if createdId == 0 {
		return api.ExistTag{}, domain.ErrorFailedInsert
	}

	data.Id = int64(createdId)
	return data.ToExistTag(), nil
}
