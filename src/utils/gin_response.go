package utils

import "github.com/gin-gonic/gin"

type SuccessResponse struct {
	Success bool                   `json:"success"`
	Data    map[string]interface{} `json:"data"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Success bool  `json:"success"`
	Error   Error `json:"error"`
}

func HandleErrorResponse(ctx *gin.Context, statusCode int, errorMessage string) {
	res := ErrorResponse{
		Success: false,
		Error: Error{
			Code:    statusCode,
			Message: errorMessage,
		},
	}
	ctx.AbortWithStatusJSON(statusCode, res)
}

func RespondWithData(ctx *gin.Context, statusCode int, data gin.H) {
	res := SuccessResponse{
		Success: true,
		Data:    data,
	}
	ctx.JSON(statusCode, res)
}
