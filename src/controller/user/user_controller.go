package controller

import (
	"github.com/esc-chula/gearfest-backend/src/interfaces"
	"github.com/esc-chula/gearfest-backend/src/usecase"
	"github.com/gin-gonic/gin"
	"gtihub.com/esc-chula/gearfest-backend/src/domain"
)

type UserController struct {
	UserUsecase usecase.UserUsecase
}

func NewUserController(sqlHandler interfaces.SqlHandler) *UserController {
	return &UserController{
		UserUsecase: usecase.UserUsecase{
			UserRepository: usecase.UserRepository{
				SqlHandler: sqlHandler,
			},
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
	id := ctx.Param("id")
	
	var newCheckin domain.Checkin
	err := ctx.BindJSON(&newCheckin)
	if err != nil {
		ctx.AbortWithStatusJSON(400,gin.H{
			"Message" : "Invalid JSON format"
		})
		return
	}




	
	



}