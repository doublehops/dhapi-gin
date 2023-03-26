package handlers

import (
	"fmt"
	"github.com/doublehops/dh-api/internal/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
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

type UpdateUserObj struct {
	Username string `json:"username" binding:"required,min=3"`
	Country  string `json:"country" binding:"required,min=3,max=3"`
}

func Validate[T any](c *gin.Context, obj *T) []responses.ValidationField {
	var errors []responses.ValidationField

	var u UpdateUserObj
	if bindErrors := c.ShouldBindJSON(obj); bindErrors != nil {
		for _, v := range bindErrors.(validator.ValidationErrors) {
			thisError := responses.ValidationField{}
			field, _ := reflect.TypeOf(&u).Elem().FieldByName(v.Field())
			fieldName, _ := field.Tag.Lookup("json")
			thisError[fieldName] = []string{v.Error()}

			errors = append(errors, thisError)
		}
		fmt.Printf(">> bindErrors: %+v <<\n", bindErrors)
	}

	return errors
}

// UpdateUser - Validation error example.
func UpdateUser(c *gin.Context) {
	user := UpdateUserObj{}
	//if err := c.ShouldBindJSON(&user); err != nil {
	//	c.AbortWithStatusJSON(http.StatusBadRequest,
	//		gin.H{
	//			"error":   "VALIDATEERR-1",
	//			"message": "Invalid inputs. Please check your inputs"})
	//	return
	//}

	errors := Validate(c, &user)
	//errors := []responses.ValidationField{
	//	{
	//		"emailAddress": []string{"email address not valid"},
	//	},
	//}

	responses.ValidationErrorResponse(c, 422, errors)
}
