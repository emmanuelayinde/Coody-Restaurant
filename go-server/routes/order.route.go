package routes

import (
	"github.com/emmanuelayinde/coody-restaurant/controllers"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.RouterGroup) {
	orderRouter := incomingRoutes.Group("/orders")
	orderRouter.GET("/", controllers.GetUsers())
	orderRouter.GET("/:order-id", controllers.GetOneUser())
}
