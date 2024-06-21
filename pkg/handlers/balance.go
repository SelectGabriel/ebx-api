package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Balance(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "all good here!",
	})
}
