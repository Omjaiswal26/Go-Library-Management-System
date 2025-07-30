package utils

import (
	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, status int, success bool, message string, data interface{}) {
	c.JSON(status, gin.H{"success": success, "message": message, "data": data})
}

func SuccessResponse(c *gin.Context, message string, data interface{}) {
	Response(c, 200, true, message, data)
}

func BadRequestResponse(c *gin.Context,) {
	Response(c, 400, false, "Invalid Payload", nil)
}

func InternalServerErrorResponse(c *gin.Context) {
	Response(c, 500, false, "Internal Server Error", nil)
}

func NotFoundResponse(c *gin.Context) {
	Response(c, 404, false, "Not found", nil)
}