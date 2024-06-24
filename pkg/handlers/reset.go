package handlers

import (
	"EBX-API/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func Reset(c *gin.Context) {
	//reseta estrutura
	models.Accounts = make(map[string]int)
	//não conehcia o c.String, demorei para acertar esse teste!1
	c.String(http.StatusOK, "OK")
	//
}
