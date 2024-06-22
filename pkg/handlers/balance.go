package handlers

import (
	"EBX-API/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Balance(c *gin.Context) {
	accountID := c.Query("account_id")

	balance, exists := models.Accounts[accountID] //acessa a estrutura das contas models.Accounts.
	if !exists {
		c.JSON(http.StatusNotFound, 0)
		return
	}

	c.JSON(http.StatusOK, balance)
}