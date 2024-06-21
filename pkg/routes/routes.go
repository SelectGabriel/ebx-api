package routes

import (
	"EBX-API/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/ping", handlers.Ping)
	r.POST("/reset", handlers.Reset)
	r.GET("/balance", handlers.Balance)
	r.POST("/event", handlers.Event)
}
