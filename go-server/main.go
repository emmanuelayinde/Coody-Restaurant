package main

import (
	"github.com/emmanuelayinde/coody-restaurant/config"
	"github.com/emmanuelayinde/coody-restaurant/db"
	"github.com/emmanuelayinde/coody-restaurant/routes"
	"github.com/gin-gonic/gin"
)

// var userCollection *mongo.Collection = db.OpenColletion(db.Client, "users")

func main() {
	db.ConnectDB()
    defer db.DisconnectDB()
	config.LoadConfig()

	router := gin.New()
	v1Router := router.Group("/api/v1")
	routes.SetupRoutes(v1Router)
	router.Run(":" + config.AppConfig.Port)
}
