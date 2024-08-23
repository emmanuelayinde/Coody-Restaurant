package routes

import (
	"github.com/emmanuelayinde/coody-restaurant/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.RouterGroup) {
	authRouter := incomingRoutes.Group("/auth")
	authRouter.POST("/sign-up", controllers.SignUp())
	authRouter.POST("/sign-in", controllers.SignIn())
	authRouter.POST("/forget-password", controllers.ForgetPassword())
	authRouter.POST("/reset-password", controllers.ResetPassword())
}
