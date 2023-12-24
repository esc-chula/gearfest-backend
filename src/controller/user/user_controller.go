package controller

import (
	"github.com/esc-chula/gearfest-backend/src/domain"
	"github.com/esc-chula/gearfest-backend/src/interfaces"
	"github.com/esc-chula/gearfest-backend/src/usecase"
	"github.com/gin-gonic/gin"
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
	
	//convert request into obj
	var requestCheckin  domain.Checkin 
	err := ctx.BindJSON(&requestCheckin)
	if err != nil {
		ctx.AbortWithStatusJSON(400,gin.H{
			"Message" : "Invalid JSON format",
		})
		return
	}

	//set UserID base on URL parameter
	requestCheckin.UserID = id

	//post the obj to db using userId,LocationId (checkInId auto gen)
	newCheckin,err := controller.UserUsecase.Post(id,requestCheckin.LocationID)
	
	if err != nil {
		
		ctx.AbortWithStatusJSON(500,gin.H{
			"Message" : "Internal server error",
		})
		
		return
	}
	ctx.JSON(201,newCheckin)
}




	
	



