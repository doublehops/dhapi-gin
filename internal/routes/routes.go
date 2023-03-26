package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/doublehops/dh-api/internal/handlers"
	"github.com/doublehops/dh-api/internal/middleware/customauth"

	"github.com/gin-gonic/gin"
)

func GetRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	v1routes(v1)
}

func v1routes(rg *gin.RouterGroup) {
	user := rg.Group("/user")

	user.GET("", handlers.ListUser)
	user.GET("/bobby", customauth.Auth(), handlers.GetUser)
	user.PUT("", handlers.UpdateUser)

	user.GET("/middleware-test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		log.Println(example)
		c.JSON(http.StatusOK, fmt.Sprintf("user: %s", example))
	})

	user.GET("/by-id/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, fmt.Sprintf("user: %s", c.Param("id")))
	})
}
