package routes

import (
	"github.com/emmanuelayinde/coody-restaurant/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.RouterGroup) {
	userRouter := incomingRoutes.Group("/users")
	userRouter.GET("/", controllers.GetUsers())
	userRouter.GET("/:user-id", controllers.GetOneUser())
}
