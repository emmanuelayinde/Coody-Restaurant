package routes

import (
	"github.com/emmanuelayinde/coody-restaurant/controllers"
	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.RouterGroup) {
	invoiceRouter := incomingRoutes.Group("/invoice")
	invoiceRouter.GET("/", controllers.GetUsers())
	invoiceRouter.GET("/:user-id", controllers.GetOneUser())
}
