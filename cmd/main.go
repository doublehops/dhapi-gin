package main

import (
	"github.com/doublehops/dh-api/internal/routes"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func main() {
	r := gin.New()
	//r.Use(logger.Logger())
	r.Use(gin.Recovery())
	routes.GetRoutes(r)

	r.Run(":8080")
}
