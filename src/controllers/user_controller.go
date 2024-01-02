package controllers

import (
	"net/http"

	"github.com/esc-chula/gearfest-backend/src/domains"
	"github.com/esc-chula/gearfest-backend/src/usecases"
	"github.com/esc-chula/gearfest-backend/src/utils"
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

func (controller *UserController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := controller.UserUsecases.Get(id)
	switch err {
	case nil:
		utils.RespondWithData(ctx, http.StatusOK, gin.H{"user": user})
	case gorm.ErrRecordNotFound:
		utils.HandleErrorResponse(ctx, http.StatusNotFound, "User not found.")
	default:
		utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "Internal server error.")
	}
}

func (controller *UserController) SignIn(ctx *gin.Context) {
	var inputUser domains.CreateUserDTO
	err := ctx.ShouldBindJSON(&inputUser)
	if err != nil {
		utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format.")
		return
	}
	user, err := controller.UserUsecases.Get(inputUser.UserID)
	if err == gorm.ErrRecordNotFound {
		//Not found then create new user
		newUser, err := controller.UserUsecases.PostCreateUser(inputUser)
		if err != nil {
			utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "Internal server error.")
			return
		}
		utils.RespondWithData(ctx, http.StatusCreated, gin.H{"user": newUser})
	} else if err == nil {
		//Found user return user
		utils.RespondWithData(ctx, http.StatusOK, gin.H{"user": user})
	} else {
		utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "Internal server error.")
	}
}

func (controller *UserController) Checkin(ctx *gin.Context) {
	//convert request into obj
	var CheckinDTO domains.CreateCheckinDTO
	err := ctx.ShouldBindJSON(&CheckinDTO)
	if err != nil {
		utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format.")
		return
	}
	//post the obj to db using userId,LocationId (checkInId auto gen)
	newCheckin, err := controller.UserUsecases.Post(CheckinDTO)
	switch err {
	case nil:
		utils.RespondWithData(ctx, http.StatusCreated, gin.H{"checkin": newCheckin})
	case gorm.ErrForeignKeyViolated:
		utils.HandleErrorResponse(ctx, http.StatusNotFound, "User not found.")
	case gorm.ErrDuplicatedKey:
		utils.HandleErrorResponse(ctx, http.StatusConflict, "User is already checked in.")
	default:
		utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "Internal server error.")
	}
}

func (controller *UserController) PatchUserName(ctx *gin.Context) {
	id := ctx.Param("id")
	//convert request into obj
	var requestDTO domains.CreateUserNameDTO
	err := ctx.ShouldBindJSON(&requestDTO)
	if err != nil {
		utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format.")
		return
	}
	//patch user in db using id,DTO
	patchedUser, err := controller.UserUsecases.PatchUserName(id, requestDTO)
	switch err {
	case nil:
		utils.RespondWithData(ctx, http.StatusOK, gin.H{"user": patchedUser})
	case gorm.ErrRecordNotFound:
		utils.HandleErrorResponse(ctx, http.StatusNotFound, "User not found.")
	default:
		utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "Internal server error.")
	}
}

func (controller *UserController) PatchUserCompleted(ctx *gin.Context) {
	id := ctx.Param("id")
	isUserCompleted, err := controller.UserUsecases.IsUserCompleted(id)
	switch err {
	case nil:
		if isUserCompleted {
			utils.HandleErrorResponse(ctx, http.StatusForbidden, "User has already completed.")
			return
		}
	case gorm.ErrRecordNotFound:
		utils.HandleErrorResponse(ctx, http.StatusNotFound, "User not found.")
	default:
		utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "Internal server error.")
	}
	//convert request into obj
	var requestDTO domains.CreateUserCompletedDTO
	err = ctx.ShouldBindJSON(&requestDTO)
	if err != nil {
		utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format.")
		return
	}
	//patch user in db using id,DTO
	patchedUser, err := controller.UserUsecases.PatchUserComplete(id, requestDTO)
	if err != nil {
		utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "Internal server error.")
		return
	}
	utils.RespondWithData(ctx, http.StatusOK, gin.H{"user": patchedUser})
}

func (controller *UserController) Reset(ctx *gin.Context) {
	id := ctx.Param("id")
	patchedUser, err := controller.UserUsecases.ResetComplete(id)
	switch err {
	case nil:
		utils.RespondWithData(ctx, http.StatusOK, gin.H{"user": patchedUser})
	case gorm.ErrRecordNotFound:
		utils.HandleErrorResponse(ctx, http.StatusNotFound, "User not found.")
	default:
		utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "Internal server error.")
	}
}
