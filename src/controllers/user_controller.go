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

// GetUser godoc
// @summary Get user
// @description Get user data domains.User.
// @tags user
// @id GetUser
// @security Bearer
// @produce json
// @Success 200 {object} utils.SuccessResponse
// @Failure      401  {object}  utils.ErrorResponse
// @Failure      403  {object}  utils.ErrorResponse
// @Failure      404  {object}  utils.ErrorResponse
// @Failure      500  {object}  utils.ErrorResponse
// @Router /user [get]
func (controller *UserController) GetUser(ctx *gin.Context) {
	id := ctx.GetString("user_id")
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

// SignIn godoc
// @summary Sign in user
// @description Return user data and create a user if the user is signing in for the first time.
// @tags user
// @id SignIn
// @security Bearer
// @produce json
// @Success 200 {object} utils.SuccessResponse
// @Success 201 {object} utils.SuccessResponse
// @Failure      401  {object}  utils.ErrorResponse
// @Failure      403  {object}  utils.ErrorResponse
// @Failure      500  {object}  utils.ErrorResponse
// @Router /user/signin [post]
func (controller *UserController) SignIn(ctx *gin.Context) {
	id := ctx.GetString("user_id")
	name := ctx.GetString("user_google_name")
	user, err := controller.UserUsecases.Get(id)
	if err == gorm.ErrRecordNotFound {
		//Not found then create new user
		newUser, err := controller.UserUsecases.PostCreateUser(id, name)
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

// Checkin godoc
// @summary Check in
// @description Check in user with location id.
// @tags user
// @id Checkin
// @security Bearer
// @accept json
// @produce json
// @param checkin body domains.CreateCheckinDTO true "Location id to be checked in."
// @Success 201 {object} utils.SuccessResponse
// @Failure      400  {object}  utils.ErrorResponse
// @Failure      401  {object}  utils.ErrorResponse
// @Failure      403  {object}  utils.ErrorResponse
// @Failure      404  {object}  utils.ErrorResponse
// @Failure      409  {object}  utils.ErrorResponse "User already checked in"
// @Failure      500  {object}  utils.ErrorResponse
// @Router /user/checkin [post]
func (controller *UserController) Checkin(ctx *gin.Context) {
	//convert request into obj
	id := ctx.GetString("user_id")
	var CheckinDTO domains.CreateCheckinDTO
	err := ctx.ShouldBindJSON(&CheckinDTO)
	if err != nil {
		utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format.")
		return
	}
	//post the obj to db using userId,LocationId (checkInId auto gen)
	newCheckin, err := controller.UserUsecases.Post(id, CheckinDTO)
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

// PatchUserName godoc
// @summary Patch user name
// @description Changed user name of the user.
// @tags user
// @id PatchUserName
// @security Bearer
// @accept json
// @produce json
// @param name body domains.CreateUserNameDTO true "Name to be changed."
// @Success 200 {object} utils.SuccessResponse
// @Failure      400  {object}  utils.ErrorResponse
// @Failure      401  {object}  utils.ErrorResponse
// @Failure      403  {object}  utils.ErrorResponse
// @Failure      404  {object}  utils.ErrorResponse
// @Failure      500  {object}  utils.ErrorResponse
// @Router /user/name [patch]
func (controller *UserController) PatchUserName(ctx *gin.Context) {
	id := ctx.GetString("user_id")
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

// PatchUserCompleted godoc
// @summary User completed
// @description Update is_user_completed to true and update cocktail_id.
// @tags user
// @id PatchUserCompleted
// @security Bearer
// @accept json
// @produce json
// @param cocktail_id body domains.CreateUserCompletedDTO true "Cocktail id of user."
// @Success 200 {object} utils.SuccessResponse
// @Failure      400  {object}  utils.ErrorResponse
// @Failure      401  {object}  utils.ErrorResponse
// @Failure      403  {object}  utils.ErrorResponse
// @Failure      404  {object}  utils.ErrorResponse
// @Failure      409  {object}  utils.ErrorResponse "User is already completed"
// @Failure      500  {object}  utils.ErrorResponse
// @Router /user/complete [patch]
func (controller *UserController) PatchUserCompleted(ctx *gin.Context) {
	id := ctx.GetString("user_id")
	isUserCompleted, err := controller.UserUsecases.IsUserCompleted(id)
	switch err {
	case nil:
		if isUserCompleted {
			utils.HandleErrorResponse(ctx, http.StatusConflict, "User has already completed.")
			return
		}
	case gorm.ErrRecordNotFound:
		utils.HandleErrorResponse(ctx, http.StatusNotFound, "User not found.")
		return
	default:
		utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "Internal server error.")
		return
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

// ResetComplete godoc
// @summary Reset complete
// @description Reset is_user_completed to false and reset cocktail_id to zero value.
// @tags user
// @id ResetComplete
// @security Bearer
// @produce json
// @Success 200 {object} utils.SuccessResponse
// @Failure      400  {object}  utils.ErrorResponse
// @Failure      401  {object}  utils.ErrorResponse
// @Failure      403  {object}  utils.ErrorResponse
// @Failure      404  {object}  utils.ErrorResponse
// @Failure      500  {object}  utils.ErrorResponse
// @Router /user/reset [patch]
func (controller *UserController) Reset(ctx *gin.Context) {
	id := ctx.GetString("user_id")
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
