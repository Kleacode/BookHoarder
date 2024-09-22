package repository

import (
	"back/src/domain"
	api "back/src/generated"
	"back/src/generated/models"
	"back/src/usecases"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

// DeleteUserIdBooksBookId implements usecases.RepositoryInterface.
func (r *Repository) DeleteUserIdBooksBookId(c *gin.Context, userId int, bookId int) error {
	result, err := r.db.ExecContext(c, "DELETE FROM books WHERE id = $1 AND user_id = $2", bookId, userId)
	if err != nil {
		return err
	}
	row, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if row == 0 {
		return domain.ErrNoAffected
	}
	return nil
}

// DeleteUserIdHoarderHoarderId implements usecases.RepositoryInterface.
func (r *Repository) DeleteUserIdHoarderHoarderId(c *gin.Context, userId int, hoarderId int) error {
	result, err := r.db.ExecContext(c, "DELETE FROM user_book_status WHERE id = $1 AND user_id = $2", hoarderId, userId)
	if err != nil {
		return err
	}
	row, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if row == 0 {
		return domain.ErrNoAffected
	}
	return nil
}

// DeleteUserIdTagsTagId implements usecases.RepositoryInterface.
func (r *Repository) DeleteUserIdTagsTagId(c *gin.Context, userId int, tagId int) error {
	result, err := r.db.ExecContext(c, "DELETE FROM tags WHERE id = $1 AND user_id = $2", tagId, userId)
	if err != nil {
		return err
	}
	row, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if row == 0 {
		return domain.ErrNoAffected
	}
	return nil
}

// GetBooks implements usecases.RepositoryInterface.
func (r *Repository) GetBooks(c *gin.Context, params *api.GetBooksParams) ([]models.Book, error) {
	var rows *sqlx.Rows
	var err error
	query := `SELECT * FROM books`
	if params.Title != nil {
		searchWord := "%" + *(params.Title) + "%"
		query += ` WHERE title LIKE $1`
		rows, err = r.db.QueryxContext(c, query, searchWord)
	} else {
		rows, err = r.db.QueryxContext(c, query)
	}

	if err != nil {
		return nil, domain.ErrInternal
	}

	defer rows.Close()
	var books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.StructScan(&book)
		if err != nil {
			return nil, domain.ErrInternal
		}
		books = append(books, book)
	}
	return books, nil
}

// GetBooksBookId implements usecases.RepositoryInterface.
func (r *Repository) GetBooksBookId(c *gin.Context, bookId int) (models.Book, error) {
	row := r.db.QueryRowxContext(c, "SELECT * FROM books WHERE id = $1", bookId)
	var book models.Book
	err := row.StructScan(&book)
	if err != nil {
		return models.Book{}, domain.ErrInternal
	}
	return book, nil
}

func (r *Repository) GetUserTags(c *gin.Context, userId int, params *api.GetUserIdTagsParams) ([]models.Tag, error) {
	var rows *sqlx.Rows
	var err error
	query := `SELECT * FROM tags`
	if params.Name != nil {
		searchWord := "%" + *params.Name + "%"
		query += ` WHERE user_id = $1 AND name LIKE $2`
		rows, err = r.db.QueryxContext(c, query, userId, searchWord)
	} else {
		query += ` WHERE user_id = $1`
		rows, err = r.db.QueryxContext(c, query, userId)
	}

	if err != nil {
		return nil, domain.ErrInternal
	}

	defer rows.Close()
	var tags []models.Tag
	for rows.Next() {
		var tag models.Tag
		err := rows.StructScan(&tag)
		if err != nil {
			return nil, domain.ErrInternal
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

// GetUserIdHoarder implements usecases.RepositoryInterface.
func (r *Repository) GetHoarders(c *gin.Context, userId int, params *api.GetUserIdHoarderParams) ([]domain.ExistHoarderRecord, error) {
	rows, err := r.db.QueryxContext(c, `SELECT
		ubs.id AS hoarder_id,
		ubs.status_id,
		books.id AS book_id,
		books.title,
		books.user_id AS book_user_id,
		array_to_string(ARRAY(SELECT unnest(array_agg(tags.id))), ',') AS tag_id_array,
		array_to_string(ARRAY(SELECT unnest(array_agg(tags.name))), ',') AS tag_name_array
	FROM user_book_status ubs 
		LEFT OUTER JOIN books ON ubs.book_id = books.id
		LEFT OUTER JOIN hoarder_tag ON ubs.id = hoarder_tag.hoarder_id
		LEFT OUTER JOIN tags ON tags.id = hoarder_tag.tag_id
	WHERE ubs.user_id = $1
	GROUP BY ubs.id, books.id
	`, userId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.ExistHoarderRecord
	for rows.Next() {
		var row domain.ExistHoarderRecord
		err := rows.StructScan(&row)
		if err != nil {
			return nil, domain.ErrInternal
		}
		result = append(result, row)
	}
	return result, nil
}

// PostUserIdTags implements usecases.RepositoryInterface.
func (r *Repository) InsertTag(c *gin.Context, data *models.Tag) (models.Tag, error) {
	var createdId int
	err := r.db.QueryRowxContext(c, `INSERT INTO tags (user_id, name, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id`,
		&data.UserID, &data.Name.String, &data.CreatedAt, &data.UpdatedAt).Scan(&createdId)
	if err != nil {
		return models.Tag{}, err
	}
	if createdId == 0 {
		return models.Tag{}, err
	}

	return models.Tag{ID: createdId, Name: data.Name, UserID: data.UserID}, nil
}

// InsertBook implements usecases.RepositoryInterface.
func (r *Repository) InsertBook(c *gin.Context, data *models.Book) (models.Book, error) {
	var createdId int
	err := r.db.QueryRowxContext(c, `INSERT INTO books (title, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id`,
		&data.Title.String, &data.UserID, &data.CreatedAt, &data.UpdatedAt).Scan(&createdId)
	if err != nil {
		return models.Book{}, err
	}
	if createdId == 0 {
		return models.Book{}, err
	}
	return models.Book{ID: createdId, Title: data.Title, UserID: data.UserID}, nil
}

func (r *Repository) InsertHoarder(c *gin.Context, data *models.UserBookStatus) (models.UserBookStatus, error) {
	var createdId int
	err := r.db.QueryRowxContext(c, `INSERT INTO user_book_status (user_id, book_id, status_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		&data.UserID, &data.BookID, &data.StatusID, &data.CreatedAt, &data.UpdatedAt).Scan(&createdId)
	if err != nil {
		return models.UserBookStatus{}, err
	}
	if createdId == 0 {
		return models.UserBookStatus{}, err
	}
	return models.UserBookStatus{ID: createdId, UserID: data.UserID, BookID: data.BookID, StatusID: data.StatusID}, nil
}

func (r *Repository) UpsertHoarderTags(c *gin.Context, data []int, hoarderId int) error {
	var existTagIDs []int
	err := r.db.SelectContext(c, &existTagIDs, `SELECT tag_id FROM hoarder_tag WHERE hoarder_id = $1`, &hoarderId)
	if err != nil {
		return err
	}
	exists := mapset.NewSet[int](existTagIDs...)
	news := mapset.NewSet[int](data...)
	deletes := exists.Difference(news)
	adds := news.Difference(exists)
	if !deletes.IsEmpty() {
		query, args, err := sqlx.In("DELETE FROM hoarder_tag WHERE hoarder_id = ? AND tag_id IN (?)", hoarderId, deletes.ToSlice())
		if err != nil {
			return err
		}
		query = r.db.Rebind(query)
		_, err = r.db.ExecContext(c, query, args...)
		if err != nil {
			return err
		}
	}

	if !adds.IsEmpty() {
		type addarg struct {
			HoarderID int       `db:"hoarder_id"`
			TagID     int       `db:"tag_id"`
			CreatedAt time.Time `db:"created_at"`
			UpdatedAt time.Time `db:"updated_at"`
		}
		var args []addarg
		created := time.Now()
		for _, e := range adds.ToSlice() {
			args = append(args, addarg{HoarderID: hoarderId, TagID: e, CreatedAt: created, UpdatedAt: created})
		}

		_, err = r.db.NamedExecContext(c, "INSERT INTO hoarder_tag (hoarder_id, tag_id, created_at, updated_at) VALUES(:hoarder_id, :tag_id, :created_at, :updated_at)", args)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Repository) UpdateBook(c *gin.Context, data *models.Book) (models.Book, error) {
	_, err := r.db.NamedExecContext(c, `UPDATE books SET (title = :title, updated_at = :updated_at) WHERE user_id = :user_id AND book_id = :book_id`, &data)
	if err != nil {
		return models.Book{}, err
	}
	return models.Book{}, nil
}

func NewRepository(db *sqlx.DB) usecases.RepositoryInterface {
	return &Repository{db: db}
}
