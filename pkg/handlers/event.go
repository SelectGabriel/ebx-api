package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Event(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "all good here!",
	})
}
