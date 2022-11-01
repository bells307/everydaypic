package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Middleware обработки ошибок
func ErrorHandler(c *gin.Context) {
	c.Next()

	var errors []string
	for _, err := range c.Errors {
		errors = append(errors, err.Error())
	}

	if len(errors) != 0 {
		c.JSON(http.StatusInternalServerError, map[string][]string{"errors": errors})
	}
}
