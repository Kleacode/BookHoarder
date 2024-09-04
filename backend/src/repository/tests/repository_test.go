package tests

import (
	api "back/src/generated"
	"back/src/repository"
	"back/src/usecases"
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

var (
	db       *sqlx.DB
	fixtures *testfixtures.Loader
)

func SetupFixtures() {
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
		testfixtures.Database(db.DB),         // You database connection
		testfixtures.Dialect("postgres"),     // Available: "postgresql", "timescaledb", "mysql", "mariadb", "sqlite" and "sqlserver"
		testfixtures.Directory("./fixtures"), // The directory containing the YAML files
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

func Test_GetBooks(t *testing.T) {
	SetupFixtures()
	prepareTestDatabase()
	ctx := context.Background()

	want := []usecases.BookRecord{
		{
			Id:     1,
			Title:  "book_title",
			UserId: 1,
			TagIds: nil,
		},
		{
			Id:     2,
			Title:  "初期登録本",
			UserId: 2,
			TagIds: nil,
		},
		{
			Id:     3,
			Title:  "book",
			UserId: 2,
			TagIds: pq.Int64Array{1, 2},
		},
	}

	repo := repository.NewRepository(db)
	books, err := repo.GetBooks(ctx, api.GetBooksParams{})

	if err != nil {
		t.Fail()
	}
	if !reflect.DeepEqual(want, books) {
		t.Fail()
	}
}
