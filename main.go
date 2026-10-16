package main

import (
	"felixj/to-do-list/todolist"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/api/to-do-list", todolist.GetToDoList)

	router.Run("localhost:8080")
}