package routes

import (
	"github.com/emmanuelayinde/coody-restaurant/controllers"
	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.RouterGroup) {
	menuRouter := incomingRoutes.Group("/menu")
	menuRouter.GET("/", controllers.GetUsers())
	menuRouter.GET("/:menu-id", controllers.GetOneUser())
}
