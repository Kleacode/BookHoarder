package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type User struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

func main() {
	password := os.Getenv("POSTGRES_PASSWORD")
	username := os.Getenv("POSTGRES_USER")
	dbname := os.Getenv("POSTGRES_DB")
	hostname := os.Getenv("HOST_NAME")

	dataSource := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", hostname, username, password, dbname)

	db, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		log.Fatalln(err)
	}

	ty := []User{}
	err2 := db.Select(&ty, "SELECT * FROM users")
	if err2 != nil {
		log.Fatalln(err2)
	}
	for _, name := range ty {
		fmt.Println(name)
	}
}
