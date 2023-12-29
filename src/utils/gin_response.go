package utils

import "github.com/gin-gonic/gin"

func HandleErrorResponse(ctx *gin.Context, statusCode int, errorMessage string) {
	ctx.AbortWithStatusJSON(statusCode, gin.H{
		"success": false,
		"error": gin.H{
			"code": statusCode,
			"message": errorMessage,
		},
	})
}

func RespondWithData(ctx *gin.Context, statusCode int, data gin.H) {
	ctx.JSON(statusCode, gin.H{
		"success": true,
		"data": data,
	})
}
