package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseBody struct {
	StatusCode int
	Message    string
	DevMessage error
	Body       map[string]interface{}
}

func successResponse(c *gin.Context, Message string, Body map[string]interface{}) {
	response := ResponseBody{
		StatusCode: http.StatusOK,
		Message:    Message,
		Body:       Body,
	}
	c.JSON(http.StatusOK, response)
}
func InternalServerErrorResponse(c *gin.Context, Error error) {
	errorResponse := ResponseBody{
		StatusCode: 500,
		Message:    "Internal Server Error",
		DevMessage: Error,
		Body:       map[string]interface{}{},
	}
	c.JSON(http.StatusInternalServerError, errorResponse)
}
