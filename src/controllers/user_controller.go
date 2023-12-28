package controllers

import (
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

func (controller *UserController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := controller.UserUsecases.Get(id)
	switch err {
	case nil:
		ctx.JSON(200, gin.H{
			"data": gin.H{
				"user": user,
			},
		})
	case gorm.ErrRecordNotFound:
		ctx.AbortWithStatusJSON(404, gin.H{
			"error": "User not found.",
		})
	default:
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": "Internal server error.",
		})
	}
}

func (controller *UserController) PostCheckin(ctx *gin.Context) {

	//convert request into obj
	var CheckinDTO domains.CreateCheckinDTO
	err := ctx.ShouldBindJSON(&CheckinDTO)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Invalid JSON format.",
		})
		return
	}
	//post the obj to db using userId,LocationId (checkInId auto gen)
	newCheckin, err := controller.UserUsecases.Post(CheckinDTO)
	switch err {
	case nil:
		ctx.JSON(201, gin.H{
			"data": gin.H{
				"checkin": newCheckin,
			},
		})
	case gorm.ErrForeignKeyViolated:
		ctx.AbortWithStatusJSON(404, gin.H{
			"error": "User not found.",
		})
	case gorm.ErrDuplicatedKey:
		ctx.AbortWithStatusJSON(409, gin.H{
			"error": "User is already checked in.",
		})
	default:
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": "Internal server error.",
		})
	}
}

func (controller *UserController) PatchUserName(ctx *gin.Context) {

	id := ctx.Param("id")
	//convert request into obj
	var requestDTO domains.CreateUserNameDTO
	err := ctx.ShouldBindJSON(&requestDTO)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Invalid JSON format.",
		})
		return
	}
	//patch user in db using id,DTO
	patchedUser, err := controller.UserUsecases.PatchUserName(id, requestDTO)
	switch err {
	case nil:
		ctx.JSON(200, gin.H{
			"data": gin.H{
				"user": patchedUser,
			},
		})
	case gorm.ErrRecordNotFound:
		ctx.AbortWithStatusJSON(404, gin.H{
			"error": "User not found.",
		})
	default:
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": "Internal server error.",
		})
	}

}

func (controller *UserController) PatchUserComplete(ctx *gin.Context) {

	id := ctx.Param("id")
	isUserCompleted, err := controller.UserUsecases.IsUserCompleted(id)
	switch err {
	case nil:
		if isUserCompleted {
			ctx.AbortWithStatusJSON(403, gin.H{
				"error": "User has already completed.",
			})
			return
		}
	case gorm.ErrRecordNotFound:
		ctx.AbortWithStatusJSON(404, gin.H{
			"error": "User not found.",
		})
	default:
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": "Internal server error.",
		})
	}
	//convert request into obj
	var requestDTO domains.CreateUserCompletedDTO
	err = ctx.ShouldBindJSON(&requestDTO)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Invalid JSON format.",
		})
		return
	}
	//patch user in db using id,DTO
	patchedUser, err := controller.UserUsecases.PatchUserComplete(id, requestDTO)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": "Internal server error.",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": gin.H{
			"user": patchedUser,
		},
	})
}
