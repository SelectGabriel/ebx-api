package main

import (
	"EBX-API/pkg/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.RegisterRoutes(r)

	fmt.Println("Starting the server...")
	r.Run(":8089")
}
