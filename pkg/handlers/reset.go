package handlers

import (
	"EBX-API/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Reset(c *gin.Context) {
	//reseta estrutura
	models.Accounts = make(map[string]int)
	c.Status(http.StatusOK)
}
