package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()	
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "all good here!",
		})
	})
	r.POST("/reset", allGood)
	r.GET("/balance", allGood)
	r.POST("/event", allGood)

	fmt.Println("Starting the server...")
	r.Run(":8089") 
}

func allGood(c *gin.Context) {
	helloWord := "Hello World"
	c.String(http.StatusOK, " %s", helloWord)
}