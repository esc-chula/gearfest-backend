package controller

import (
	"github.com/esc-chula/gearfest-backend/src/domain"
	"github.com/esc-chula/gearfest-backend/src/usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase usecase.UserUsecase
}

func NewUserController(repository usecase.UserRepository) *UserController {
	return &UserController{
		UserUsecase: usecase.UserUsecase{
			UserRepository: repository,
		},
	}
}

func (controller *UserController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := controller.UserUsecase.Get(id)
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
	var CheckinDTO domain.CreateCheckinDTO
	err := ctx.ShouldBindJSON(&CheckinDTO)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"Message": "Invalid JSON format",
		})
		return
	}
	//post the obj to db using userId,LocationId (checkInId auto gen)
	newCheckin, err := controller.UserUsecase.Post(CheckinDTO)

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"Message": "Internal server error",
		})

		return
	}
	ctx.JSON(201, newCheckin)
}
