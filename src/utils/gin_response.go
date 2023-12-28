package utils

import "github.com/gin-gonic/gin"

func HandleErrorResponse(ctx *gin.Context, statusCode int, errorMessage string) {
	ctx.AbortWithStatusJSON(statusCode, gin.H{
		"error": errorMessage,
	})
}

func RespondWithData(ctx *gin.Context, statusCode int, data gin.H) {
	ctx.JSON(statusCode, gin.H{
		"data": data,
	})
}
