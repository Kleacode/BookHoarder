package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type User struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

func main() {

	fmt.Println("hello world")

	db, err := sqlx.Connect("postgres", "host=db user=user password=example dbname=test sslmode=disable")
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
