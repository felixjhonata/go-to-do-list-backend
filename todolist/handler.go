package todolist

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

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