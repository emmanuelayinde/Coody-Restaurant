package routes

import (
	"github.com/emmanuelayinde/coody-restaurant/controllers"
	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.RouterGroup) {
	foodRouter := incomingRoutes.Group("/food")
	foodRouter.GET("/", controllers.GetUsers())
	foodRouter.GET("/:user-id", controllers.GetOneUser())
}
