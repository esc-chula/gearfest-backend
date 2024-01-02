package server

import (
	"github.com/esc-chula/gearfest-backend/src/config"
	"github.com/esc-chula/gearfest-backend/src/controllers"
	"github.com/esc-chula/gearfest-backend/src/server/repositories"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func loadRoutes(db *gorm.DB, cfg config.GoogleConfig) *gin.Engine {
	g := gin.New()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	corsConfig.AllowCredentials = true
	corsConfig.AllowAllOrigins = true
	g.Use(cors.New(corsConfig))
	userRoutes := g.Group("/user")
	userRoutes.Use(Validation(cfg))
	loadUserRoutes(userRoutes, db)
	return g
}

func loadUserRoutes(g *gin.RouterGroup, db *gorm.DB) {
	UserRepository := repositories.NewUserRepository(db)
	UserController := controllers.NewUserController(UserRepository)
	g.GET("", UserController.GetUser)
	g.POST("/signin", UserController.SignIn)
	g.POST("/checkin", UserController.Checkin)
	g.PATCH("/complete", UserController.PatchUserCompleted)
	g.PATCH("/name", UserController.PatchUserName)
	g.PATCH("/reset", UserController.Reset)
}
