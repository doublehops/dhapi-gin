package main

import (
	"github.com/doublehops/dh-api/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	//r.Use(logger.Logger()) // add middleware for all endpoints.
	r.Use(gin.Recovery())
	routes.GetRoutes(r)

	r.Run(":8080")
}
