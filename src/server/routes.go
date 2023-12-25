package server

import (
	controller "github.com/esc-chula/gearfest-backend/src/controller/user"
	"github.com/esc-chula/gearfest-backend/src/server/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func loadRoutes(db *gorm.DB) *gin.Engine {
	g := gin.New()
	loadUserRoutes(g, db)
	return g
}

func loadUserRoutes(g *gin.Engine, db *gorm.DB) {
	UserRepository := repository.NewUserRepository(db)
	user_controller := controller.NewUserController(UserRepository)
	g.GET("/user/:id", user_controller.GetUser)
	//post checkin route
	g.POST("/user/checkin", user_controller.PostCheckin)
	g.PATCH("user/complete",user_controller.PatchUser)
	g.PATCH("user/name",user_controller.PatchUser)
}
