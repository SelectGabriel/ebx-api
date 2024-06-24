package handlers

import (
	"EBX-API/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Event(c *gin.Context) {
	var event struct {
		Type        string `json:"type"`
		Destination string `json:"destination,omitempty"`
		Origin      string `json:"origin,omitempty"`
		Amount      int    `json:"amount"`
	}

	if err := c.BindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//checa operação e chega validade
	switch event.Type {
	case "deposit":
		models.Accounts[event.Destination] += event.Amount
		c.JSON(http.StatusCreated, gin.H{"destination": gin.H{"id": event.Destination, "balance": models.Accounts[event.Destination]}})
	case "withdraw":
		balance, exists := models.Accounts[event.Origin]
		if !exists || balance < event.Amount {
			c.JSON(http.StatusNotFound, 0)
			return
		}
		models.Accounts[event.Origin] -= event.Amount
		c.JSON(http.StatusCreated, gin.H{"origin": gin.H{"id": event.Origin, "balance": models.Accounts[event.Origin]}})
	case "transfer":
		originBalance, originExists := models.Accounts[event.Origin]
		if !originExists || originBalance < event.Amount {
			c.JSON(http.StatusNotFound, 0)
			return
		}
		models.Accounts[event.Origin] -= event.Amount
		models.Accounts[event.Destination] += event.Amount
		c.JSON(http.StatusCreated, gin.H{
			"origin":      gin.H{"id": event.Origin, "balance": models.Accounts[event.Origin]},
			"destination": gin.H{"id": event.Destination, "balance": models.Accounts[event.Destination]},
		})
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event type"})
	}
}