package controllers

import (
	"github.com/esc-chula/gearfest-backend/src/domains"
	"github.com/esc-chula/gearfest-backend/src/usecases"
	"github.com/gin-gonic/gin"
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
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"Message": "Bad request",
		})
		return
	}
	ctx.JSON(200, user)
}

func (controller *UserController) PostCheckin(ctx *gin.Context) {

	//convert request into obj
	var CheckinDTO domains.CreateCheckinDTO
	err := ctx.ShouldBindJSON(&CheckinDTO)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"Message": "Invalid JSON format",
		})
		return
	}
	//post the obj to db using userId,LocationId (checkInId auto gen)
	newCheckin, err := controller.UserUsecases.Post(CheckinDTO)

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"Message": "Internal server error",
		})

		return
	}
	ctx.JSON(201, newCheckin)
}

func (controller *UserController) PatchUserName(ctx *gin.Context) {

	id := ctx.Param("id")
	//convert request into obj
	var requestDTO domains.CreateUserNameDTO
	err := ctx.ShouldBindJSON(&requestDTO)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"Message": "Invalid JSON format",
		})
		return
	}
	//patch user in db using id,DTO
	patchedUser, err := controller.UserUsecases.PatchUserName(id, requestDTO)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"Message": "Internal server error",
		})
		return
	}
	ctx.JSON(200, patchedUser)

}

func (controller *UserController) PatchUserComplete(ctx *gin.Context) {

	id := ctx.Param("id")
	//convert request into obj
	var requestDTO domains.CreateUserCompletedDTO
	err := ctx.ShouldBindJSON(&requestDTO)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"Message": "Invalid JSON format",
		})
		return
	}
	//patch user in db using id,DTO
	patchedUser, err := controller.UserUsecases.PatchUserComplete(id, requestDTO)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"Message": "Internal server error",
		})
		return
	}
	ctx.JSON(200, patchedUser)

}