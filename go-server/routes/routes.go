package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(v1Router *gin.RouterGroup) {
	v1Router.Use(gin.Logger())
	AuthRoutes(v1Router)
	// v1Router.Use(middlewares.AuthMiddleware())

	// RestaurantRoutes(v1Router)
	// UserRoutes(v1Router)
	// FoodRoutes(v1Router)
	// MenuRoutes(v1Router)
	// OrderRoutes(v1Router)
	// TableRoutes(v1Router)
	// OrderItemRoutes(v1Router)
	// InvoiceRoutes(v1Router)
}
