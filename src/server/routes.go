package server

import (
	"github.com/esc-chula/gearfest-backend/src/controllers"
	"github.com/esc-chula/gearfest-backend/src/server/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func loadRoutes(db *gorm.DB) *gin.Engine {
	g := gin.New()
	userRoutes := g.Group("/user")
	loadUserRoutes(userRoutes, db)
	return g
}

func loadUserRoutes(g *gin.RouterGroup, db *gorm.DB) {
	UserRepository := repositories.NewUserRepository(db)
	UserController := controllers.NewUserController(UserRepository)
	g.GET("/:id", UserController.GetUser)
	g.POST("/create", UserController.PostUser)
	g.POST("/checkin", UserController.PostCheckin)
	g.PATCH("/complete/:id", UserController.PatchUserComplete)
	g.PATCH("/name/:id", UserController.PatchUserName)
}
