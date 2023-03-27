package handlers

import (
	"errors"
	"fmt"
	"github.com/doublehops/dh-api/internal/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
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
	Country  string `json:"country" binding:"min=3,max=3"`
}

func GetErrorMessages() map[string]map[string]string {
	return map[string]map[string]string{
		"username": {"required": "This field is required"},
		"country":  {"required": "This field is required", "min": "Must be at least 3 characters long"},
	}
}

func getStandardErrors(tag string) string {
	errorMap := map[string]string{
		"required": "this is a required field",
	}

	if _, ok := errorMap[tag]; ok {
		return errorMap[tag]
	}

	return ""
}

func getErrorMessage(field, tag, ginError string) string {
	messages := GetErrorMessages()
	if _, ok := messages[field]; ok {
		if _, ok = messages[field][tag]; ok {
			return messages[field][tag]
		}
	}

	stdError := getStandardErrors(tag)
	if stdError != "" {
		return stdError
	}

	return parseGinError(ginError)
}

func parseGinError(str string) string {
	pos := strings.Index(str, "Error:")
	substr := str[pos+6:]

	return substr
}

func Validate[T any](c *gin.Context, obj *T) []responses.ValidationField {
	var errs []responses.ValidationField

	err := c.BindQuery(&obj)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			for _, fe := range ve {
				thisError := responses.ValidationField{}
				field, _ := reflect.TypeOf(obj).Elem().FieldByName(fe.Field())
				fieldName, _ := field.Tag.Lookup("json")
				thisError[fieldName] = []string{getErrorMessage(fieldName, fe.Tag(), fe.Error())}
				errs = append(errs, thisError)
			}
		}
	}

	return errs
}

func ValidateSSS[T any](c *gin.Context, obj *T) []responses.ValidationField {
	var errors []responses.ValidationField

	if bindErrors := c.ShouldBindJSON(obj); bindErrors != nil {
		for _, v := range bindErrors.(validator.ValidationErrors) {
			thisError := responses.ValidationField{}
			field, _ := reflect.TypeOf(obj).Elem().FieldByName(v.Field())
			fieldName, _ := field.Tag.Lookup("json")
			//thisErr := v.(fieldError)
			thisError[fieldName] = []string{v.Error()}
			//thisError[fieldName] = []string{getErrorMessage(fieldName, thisError.Tag())}

			errors = append(errors, thisError)
		}
		fmt.Printf(">> bindErrors: %+v <<\n", bindErrors)
	}

	return errors
}

// UpdateUser - Validation error example.
func UpdateUser(c *gin.Context) {
	var user UpdateUserObj

	_ = c.ShouldBindJSON(&user)

	validationErrors := Validate(c, &user)

	responses.ValidationErrorResponse(c, 422, validationErrors)
}
