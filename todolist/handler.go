package todolist

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetToDoList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, getToDoItems())
}