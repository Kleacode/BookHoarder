package main

import (
	api "back/src/generated"
	"back/src/handler"
	"back/src/repository"
	"back/src/usecases"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	dbname := os.Getenv("POSTGRES_DEV_DB")
	hostname := os.Getenv("HOST_NAME")

	dataSource := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", hostname, username, password, dbname)

	db, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8001","http://localhost:3000"},
		AllowMethods:     []string{"GET", "DELETE", "PUT", "PATCH", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	repo := repository.NewRepository(db)
	service := usecases.NewService(repo)
	handler := handler.NewHandler(service)

	api.RegisterHandlers(r, handler)

	r.Run()
}
