package routes

import (
	"github.com/emmanuelayinde/coody-restaurant/controllers"
	"github.com/gin-gonic/gin"
)

func RestaurantRoutes(incomingRoutes *gin.RouterGroup) {
	restaurantRouter := incomingRoutes.Group("/restaurant")
	restaurantRouter.GET("/", controllers.GetUsers())
	restaurantRouter.GET("/:restaurant-id", controllers.GetOneUser())
}
