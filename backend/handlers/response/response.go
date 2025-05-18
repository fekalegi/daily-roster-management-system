package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Code         int         `json:"code"`
	Status       string      `json:"status"`
	ErrorMessage string      `json:"error_message,omitempty"`
	Data         interface{} `json:"data,omitempty"`
}

func JSON(c *gin.Context, code int, data interface{}) {
	c.JSON(code, APIResponse{
		Code:         code,
		Status:       "success",
		ErrorMessage: "",
		Data:         data,
	})
}

func Created(c *gin.Context, data interface{}) {
	JSON(c, http.StatusCreated, data)
}

func Error(c *gin.Context, code int, message string) {
	c.JSON(code, APIResponse{
		Code:         code,
		Status:       "error",
		ErrorMessage: message,
		Data:         nil,
	})
}
