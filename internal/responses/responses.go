package responses

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type singleItemResponse struct {
	Data interface{} `json:"data"`
}

type multiItemResponse struct {
	Data interface{} `json:"data"`
	Pagination
}

type Pagination struct {
	CurrentPage int `json:"currentPage"`
	PerPage     int `json:"perPage"`
	PageCount   int `json:"pageCount"`
	TotalCount  int `json:"totalCount"`
}

func SingleItemResponse(c *gin.Context, data interface{}) {
	resp := singleItemResponse{
		Data: data,
	}

	c.JSON(http.StatusOK, resp)
}

func MultiItemResponse(c *gin.Context, data interface{}, pagination Pagination) {
	resp := multiItemResponse{
		Data:       data,
		Pagination: pagination,
	}

	c.JSON(http.StatusOK, resp)
}
