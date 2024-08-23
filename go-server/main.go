package main

import (
	"github.com/emmanuelayinde/coody-restaurant/config"
	"github.com/emmanuelayinde/coody-restaurant/db"
	"github.com/emmanuelayinde/coody-restaurant/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	v1Router := router.Group("/api/v1")
	routes.SetupRoutes(v1Router)
	db.OpenCollection(db.Client, "users")
	router.Run(":" + config.AppConfig.PORT)
}
