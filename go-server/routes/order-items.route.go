package routes

import (
	"github.com/emmanuelayinde/coody-restaurant/controllers"
	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.RouterGroup) {
	orderItemRouter := incomingRoutes.Group("/order-items")
	orderItemRouter.GET("/", controllers.GetUsers())
	orderItemRouter.GET("/:user-id", controllers.GetOneUser())
}
