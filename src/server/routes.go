package server

import (
	controller "github.com/esc-chula/gearfest-backend/src/controller/user"
	"github.com/esc-chula/gearfest-backend/src/interfaces"
	"github.com/gin-gonic/gin"
)

func loadRoutes(sqlHandler interfaces.SqlHandler) *gin.Engine {
	g := gin.New()

	//user routes (need to refactor soon)
	user_controller := controller.NewUserController(sqlHandler)
	g.GET("/user/:id", user_controller.GetUser)

	return g
}