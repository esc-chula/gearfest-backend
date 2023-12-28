package controllers

import (
	"net/http"

	"github.com/esc-chula/gearfest-backend/src/domains"
	"github.com/esc-chula/gearfest-backend/src/usecases"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	UserUsecases usecases.UserUsecases
}

func NewUserController(repository usecases.UserRepository) *UserController {
	return &UserController{
		UserUsecases: usecases.UserUsecases{
			UserRepository: repository,
		},
	}
}

func handleErrorResponse(ctx *gin.Context, statusCode int, errorMessage string) {
	ctx.AbortWithStatusJSON(statusCode, gin.H{
		"error": errorMessage,
	})
}

func respondWithData(ctx *gin.Context, statusCode int, data gin.H) {
	ctx.JSON(statusCode, gin.H{
		"data": data,
	})
}

func (controller *UserController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := controller.UserUsecases.Get(id)
	switch err {
	case nil:
		respondWithData(ctx, http.StatusOK, gin.H{"user": user})
	case gorm.ErrRecordNotFound:
		handleErrorResponse(ctx, http.StatusNotFound, "User not found.")
	default:
		handleErrorResponse(ctx, http.StatusInternalServerError, "Internal server error.")
	}
}

func (controller *UserController) PostCheckin(ctx *gin.Context) {

	//convert request into obj
	var CheckinDTO domains.CreateCheckinDTO
	err := ctx.ShouldBindJSON(&CheckinDTO)
	if err != nil {
		handleErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format.")
		return
	}
	//post the obj to db using userId,LocationId (checkInId auto gen)
	newCheckin, err := controller.UserUsecases.Post(CheckinDTO)
	switch err {
	case nil:
		respondWithData(ctx, http.StatusCreated, gin.H{"checkin": newCheckin})
	case gorm.ErrForeignKeyViolated:
		handleErrorResponse(ctx, http.StatusNotFound, "User not found.")
	case gorm.ErrDuplicatedKey:
		handleErrorResponse(ctx, http.StatusConflict, "User is already checked in.")
	default:
		handleErrorResponse(ctx, http.StatusInternalServerError, "Internal server error.")	
	}
}

func (controller *UserController) PatchUserName(ctx *gin.Context) {

	id := ctx.Param("id")
	//convert request into obj
	var requestDTO domains.CreateUserNameDTO
	err := ctx.ShouldBindJSON(&requestDTO)
	if err != nil {
		handleErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format.")
		return
	}
	//patch user in db using id,DTO
	patchedUser, err := controller.UserUsecases.PatchUserName(id, requestDTO)
	switch err {
	case nil:
		respondWithData(ctx, http.StatusOK, gin.H{"user": patchedUser})
	case gorm.ErrRecordNotFound:
		handleErrorResponse(ctx, http.StatusNotFound, "User not found.")
	default:
		handleErrorResponse(ctx, http.StatusInternalServerError, "Internal server error.")
	}

}

func (controller *UserController) PatchUserComplete(ctx *gin.Context) {

	id := ctx.Param("id")
	isUserCompleted, err := controller.UserUsecases.IsUserCompleted(id)
	switch err {
	case nil:
		if isUserCompleted {
			handleErrorResponse(ctx, http.StatusForbidden, "User has already completed.")
			return
		}
	case gorm.ErrRecordNotFound:
		handleErrorResponse(ctx, http.StatusNotFound, "User not found.")
	default:
		handleErrorResponse(ctx, http.StatusInternalServerError, "Internal server error.")
	}
	//convert request into obj
	var requestDTO domains.CreateUserCompletedDTO
	err = ctx.ShouldBindJSON(&requestDTO)
	if err != nil {
		handleErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format.")
		return
	}
	//patch user in db using id,DTO
	patchedUser, err := controller.UserUsecases.PatchUserComplete(id, requestDTO)
	if err != nil {
		handleErrorResponse(ctx, http.StatusInternalServerError, "Internal server error.")
		return
	}
	respondWithData(ctx, http.StatusOK, gin.H{"user": patchedUser})
}
