package todolist

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type todoInput struct {
	Task     string `json:"task" binding:"required"`
	Assignee string `json:"assignee" binding:"required"`
}

func GetToDoList(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		items, err := listToDoItems(c.Request.Context(), pool)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch todos"})
			return
		}
		c.IndentedJSON(http.StatusOK, items)
	}
}

func AddToDoItem(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in todoInput
		if err := c.ShouldBindJSON(&in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		item, err := createToDoItem(c.Request.Context(), pool, in.Task, in.Assignee)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create todo"})
			return
		}
		c.IndentedJSON(http.StatusCreated, item)
	}
}

func UpdateToDoItem(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		var in todoInput
		if err := c.ShouldBindJSON(&in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		item, err := editToDoItem(c.Request.Context(), pool, id, in.Task, in.Assignee)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update todo"})
			return
		}
		c.IndentedJSON(http.StatusOK, item)
	}
}

func DeleteToDoItem(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		n, err := removeToDoItem(c.Request.Context(), pool, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete todo"})
			return
		}
		if n == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
			return
		}
		c.Status(http.StatusNoContent)
	}
}