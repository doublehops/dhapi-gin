package customauth

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("username", "john")
		c.Set("emailAddress", "john@example.com")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print("latency: ", latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
