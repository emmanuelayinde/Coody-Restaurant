package main

import (
	"github.com/emmanuelayinde/coody-restaurant/config"
	"github.com/emmanuelayinde/coody-restaurant/db"
	"github.com/emmanuelayinde/coody-restaurant/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	router := gin.New()
	v1Router := router.Group("/api/v1")
	routes.SetupRoutes(v1Router)
	db.OpenCollection(db.Client, "users")

	v1Router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})
	v1Router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "ROUTE_NOT_FOUND", "message": "Route not found"})
	})
	router.Run(":" + config.AppConfig.PORT)
}
