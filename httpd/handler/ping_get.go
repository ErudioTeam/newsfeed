package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingGet(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"hello": "Found me",
	})
}

func PingGetv1() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"hello": "Found me",
		})
	}
}
