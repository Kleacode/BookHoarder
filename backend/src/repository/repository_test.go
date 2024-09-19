package repository

import (
	"back/src/domain"
	api "back/src/generated"
	"back/src/generated/models"
	"back/src/usecases"
	"fmt"
	"log"
	"net/http/httptest"
	"os"
	"reflect"
	"sort"
	"testing"
	"time"

	"github.com/DATA-DOG/go-txdb"
	"github.com/gin-gonic/gin"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/null/v8"

	_ "github.com/lib/pq"
)

var (
	db       *sqlx.DB
	repo     usecases.RepositoryInterface
	fixtures *testfixtures.Loader
)

func strPtr(v string) *string {
	return &v
}

func setupFixtures() {
	password := os.Getenv("POSTGRES_PASSWORD")
	username := os.Getenv("POSTGRES_USER")
	dbname := os.Getenv("POSTGRES_TEST_DB")
	hostname := os.Getenv("HOST_NAME")
	dataSource := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", hostname, username, password, dbname)

	var err error
	db, err = sqlx.Connect("postgres", dataSource)
	if err != nil {
		log.Fatalf("Could not connect database: %s", err)
	}

	fixtures, err = testfixtures.New(
		testfixtures.Database(db.DB),                  // You database connection
		testfixtures.Dialect("postgres"),              // Available: "postgresql", "timescaledb", "mysql", "mariadb", "sqlite" and "sqlserver"
		testfixtures.Directory("./testdata/fixtures"), // The directory containing the YAML files
	)
	if err != nil {
		log.Fatalf("Failed setup fixtures: %s", err)
	}
}

func prepareTestDatabase() {
	if err := fixtures.Load(); err != nil {
		log.Fatalf("Failed load fixtures: %s", err)
	}
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	setupFixtures()
	prepareTestDatabase()

	txdb.Register("txdb", "postgres", "root@/txdb_test")
	repo = NewRepository(db)
	m.Run()
	defer db.Close()
}

func Test_GetBooks(t *testing.T) {
	testcases := []struct {
		name string
		args api.GetBooksParams
		want []models.Book
	}{
		{
			name: "全体検索",
			args: api.GetBooksParams{},
			want: []models.Book{
				{
					ID:        1,
					Title:     null.String{String: "book_title", Valid: true},
					UserID:    1,
					CreatedAt: time.Date(2020, time.December, 31, 23, 59, 59, 0, time.FixedZone("", 0)),
					UpdatedAt: time.Date(2020, time.December, 31, 23, 59, 59, 0, time.FixedZone("", 0)),
				},
				{
					ID:        2,
					Title:     null.String{String: "初期登録本", Valid: true},
					UserID:    2,
					CreatedAt: time.Date(2020, time.December, 31, 23, 59, 59, 0, time.FixedZone("", 0)),
					UpdatedAt: time.Date(2020, time.December, 31, 23, 59, 59, 0, time.FixedZone("", 0)),
				},
				{
					ID:        3,
					Title:     null.String{String: "book", Valid: true},
					UserID:    2,
					CreatedAt: time.Date(2020, time.December, 31, 23, 59, 59, 0, time.FixedZone("", 0)),
					UpdatedAt: time.Date(2020, time.December, 31, 23, 59, 59, 0, time.FixedZone("", 0)),
				},
			},
		},
		{
			name: "文字列検索1",
			args: api.GetBooksParams{Title: strPtr("book")},
			want: []models.Book{
				{
					ID:        1,
					Title:     null.String{String: "book_title", Valid: true},
					UserID:    1,
					CreatedAt: time.Date(2020, time.December, 31, 23, 59, 59, 0, time.FixedZone("", 0)),
					UpdatedAt: time.Date(2020, time.December, 31, 23, 59, 59, 0, time.FixedZone("", 0)),
				},
				{
					ID:        3,
					Title:     null.String{String: "book", Valid: true},
					UserID:    2,
					CreatedAt: time.Date(2020, time.December, 31, 23, 59, 59, 0, time.FixedZone("", 0)),
					UpdatedAt: time.Date(2020, time.December, 31, 23, 59, 59, 0, time.FixedZone("", 0)),
				},
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			books, err := repo.GetBooks(ctx, &tt.args)

			if err != nil {
				t.Errorf("error: %v", err)
			}

			sort.SliceStable(books, func(i, j int) bool {
				return books[i].ID < books[j].ID
			})

			if !reflect.DeepEqual(books, tt.want) {
				t.Errorf("mismatch result:%v want:%v", books, tt.want)
			}
		})
	}
}

func Test_GetBookBookId(t *testing.T) {
	testcases := []struct {
		name    string
		args    int
		want    models.Book
		wantErr error
	}{
		{
			name: "ok1",
			args: 1,
			want: models.Book{

				ID:     1,
				Title:  null.String{String: "book_title", Valid: true},
				UserID: 1,
			},
			wantErr: nil,
		},
		{
			name: "ok2",
			args: 3,
			want: models.Book{

				ID:     3,
				Title:  null.String{String: "book", Valid: true},
				UserID: 2,
			},
			wantErr: nil,
		},
		{
			name:    "ng",
			args:    -1,
			want:    models.Book{},
			wantErr: domain.ErrInternal,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			book, err := repo.GetBooksBookId(ctx, tt.args)

			if tt.wantErr != nil {
				if tt.wantErr != err {
					t.Errorf("mismatch error result:%v want:%v", err, tt.wantErr)
				}
			} else {
				if !reflect.DeepEqual(book, tt.want) {
					t.Errorf("mismatch result:%v want:%v", book, tt.want)
				}
			}

		})
	}
}

func Test_GetTags(t *testing.T) {
	testcases := []struct {
		name   string
		userID int
		param  api.GetUserIdTagsParams
		want   []models.Tag
	}{
		{
			name:   "タグの取得1",
			userID: 1,
			param:  api.GetUserIdTagsParams{},
			want: []models.Tag{
				{
					ID:     1,
					Name:   null.String{String: "tag1", Valid: true},
					UserID: 1,
				},
			},
		},
		{
			name:   "タグの取得2",
			userID: 2,
			param:  api.GetUserIdTagsParams{},
			want: []models.Tag{
				{
					ID:     2,
					Name:   null.String{String: "tag2", Valid: true},
					UserID: 2,
				},
				{
					ID:     3,
					Name:   null.String{String: "tag3", Valid: true},
					UserID: 2,
				},
			},
		},
		{
			name:   "タグの取得3",
			userID: 3,
			param:  api.GetUserIdTagsParams{},
			want: []models.Tag{
				{
					ID:     4,
					Name:   null.String{String: "3t_ttt", Valid: true},
					UserID: 3,
				},
			},
		},
		{
			name:   "検索1",
			userID: 2,
			param:  api.GetUserIdTagsParams{Name: strPtr("3")},
			want: []models.Tag{
				{
					ID:     3,
					Name:   null.String{String: "tag3", Valid: true},
					UserID: 2,
				},
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			result, err := repo.GetUserTags(ctx, tt.userID, &tt.param)

			if err != nil {
				t.Errorf("error: %v", err)
			}
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("result: %v, want: %v", result, tt.want)
			}
		})
	}
}

func Test_Series_Book(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	var initbooks []models.Book
	var err error

	initbooks, err = repo.GetBooks(ctx, &api.GetBooksParams{})
	if err != nil {
		t.Errorf("error")
	}

	post1 := models.Book{Title: null.NewString("new book1", true), UserID: 1}
	book1, err := repo.InsertBook(ctx, &post1)
	if err != nil {
		t.Errorf("%v", err)
	}

	initbooks = append(initbooks, book1)
	books, err := repo.GetBooks(ctx, &api.GetBooksParams{})

	if !reflect.DeepEqual(initbooks, books) {
		t.Errorf("mismatch %v, %v", initbooks, books)
	}

	post2 := models.Book{Title: null.NewString("new book2", true), UserID: 2}
	book2, err := repo.InsertBook(ctx, &post2)

	initbooks = append(initbooks, book2)
	books, err = repo.GetBooks(ctx, &api.GetBooksParams{})
	if !reflect.DeepEqual(initbooks, books) {
		t.Errorf("mismatch %v, %v", initbooks, books)
	}

	repo.DeleteUserIdBooksBookId(ctx, book2.UserID, book2.ID)
	initbooks = initbooks[:len(initbooks)-1]
	books, err = repo.GetBooks(ctx, &api.GetBooksParams{})
	if !reflect.DeepEqual(initbooks, books) {
		t.Errorf("mismatch %v, %v", initbooks, books)
	}

	repo.DeleteUserIdBooksBookId(ctx, book1.UserID, book1.ID)
	initbooks = initbooks[:len(initbooks)-1]
	books, err = repo.GetBooks(ctx, &api.GetBooksParams{})
	if !reflect.DeepEqual(initbooks, books) {
		t.Errorf("mismatch %v, %v", initbooks, books)
	}
}

func Test_Series_Tag(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	var inittags []models.Tag
	var err error

	inittags, err = repo.GetUserTags(ctx, 1, &api.GetUserIdTagsParams{})
	if err != nil {
		t.Errorf("error")
	}

	post1 := models.Tag{Name: null.NewString("new tag1", true), UserID: 1}
	tag1, err := repo.InsertTag(ctx, &post1)

	inittags = append(inittags, tag1)
	tags, err := repo.GetUserTags(ctx, 1, &api.GetUserIdTagsParams{})
	if !reflect.DeepEqual(inittags, tags) {
		t.Errorf("mismatch %v, %v", inittags, tags)
	}

	post2 := models.Tag{Name: null.NewString("new tag2", true), UserID: 1}
	tag2, err := repo.InsertTag(ctx, &post2)

	inittags = append(inittags, tag2)
	tags, err = repo.GetUserTags(ctx, 1, &api.GetUserIdTagsParams{})
	if !reflect.DeepEqual(inittags, tags) {
		t.Errorf("mismatch %v, %v", inittags, tags)
	}

	repo.DeleteUserIdTagsTagId(ctx, tag2.UserID, tag2.ID)
	inittags = inittags[:len(inittags)-1]
	tags, err = repo.GetUserTags(ctx, 1, &api.GetUserIdTagsParams{})
	if !reflect.DeepEqual(inittags, tags) {
		t.Errorf("mismatch %v, %v", inittags, tags)
	}

	repo.DeleteUserIdTagsTagId(ctx, tag1.UserID, tag1.ID)
	inittags = inittags[:len(inittags)-1]
	tags, err = repo.GetUserTags(ctx, 1, &api.GetUserIdTagsParams{})
	if !reflect.DeepEqual(inittags, tags) {
		t.Errorf("mismatch %v, %v", inittags, tags)
	}
}
