package routes

import (
	"github.com/emmanuelayinde/coody-restaurant/controllers"
	"github.com/gin-gonic/gin"
)

func TableRoutes(incomingRoutes *gin.RouterGroup) {
	tableRouter := incomingRoutes.Group("/table")
	tableRouter.GET("/", controllers.GetUsers())
	tableRouter.GET("/:user-id", controllers.GetOneUser())
}
