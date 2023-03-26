package handlers

import (
	"github.com/doublehops/dh-api/internal/responses"

	"github.com/gin-gonic/gin"
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

	p := responses.PaginationType{
		CurrentPage: 1,
		PerPage:     10,
		PageCount:   22,
		TotalCount:  229,
	}

	responses.MultiItemResponse(c, users, p)
}

// UpdateUser - Validation error example.
func UpdateUser(c *gin.Context) {
	errors := []responses.ValidationField{
		{
			"emailAddress": []string{"email address not valid"},
		},
	}

	responses.ValidationErrorResponse(c, 422, errors)
}
