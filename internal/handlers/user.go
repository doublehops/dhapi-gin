package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/doublehops/dh-api/internal/responses"
)

type User struct {
	Username     string `json:"username"`
	EmailAddress string `json:"emailAddress"`
}

func GetUser(c *gin.Context) {
	user := User{
		Username:     c.MustGet("username").(string),
		EmailAddress: c.MustGet("emailAddress").(string),
	}

	responses.SingleItemResponse(c, user)
}

func ListUser(c *gin.Context) {
	users := []User{
		{
			Username:     "Alice",
			EmailAddress: "alice@example.com",
		},
		{
			Username:     "Bob",
			EmailAddress: "bob@example.com",
		},
		{
			Username:     "Carol",
			EmailAddress: "carol@example.com",
		},
	}

	p := responses.Pagination{
		CurrentPage: 1,
		PerPage:     10,
		PageCount:   22,
		TotalCount:  229,
	}

	responses.MultiItemResponse(c, users, p)
}
