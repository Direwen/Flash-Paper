package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error,omitempty"`
}

func SendSuccess(c *gin.Context, code int, data interface{}) {
	c.JSON(
		code,
		Response{
			Success: true,
			Data:    data,
		},
	)
}

func SendMessage(c *gin.Context, code int, message string) {
	c.JSON(
		code,
		Response{
			Success: true,
			Message: message,
		},
	)
}

func SendError(c *gin.Context, code int, err error) {
	c.JSON(
		code,
		Response{
			Success: false,
			Error:   err.Error(),
		},
	)
}
