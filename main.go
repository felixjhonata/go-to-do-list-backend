package main

import (
	"context"
	"log"
	"os"

	"felixj/to-do-list/todolist"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	if err := pool.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.GET("/api/to-do-list", todolist.GetToDoList(pool))
	router.POST("/api/to-do-list", todolist.AddToDoItem(pool))
	router.PUT("/api/to-do-list/:id", todolist.UpdateToDoItem(pool))
	router.DELETE("/api/to-do-list/:id", todolist.DeleteToDoItem(pool))

	router.Run("localhost:8080")
}