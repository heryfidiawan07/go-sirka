package main

import (
	"app/config"
	"app/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	// defer config.DB.Close()

	router := gin.Default()

	// Home
	router.GET("/", controller.Home)

	// User
	router.GET("/DisplayAllUsers", controller.DisplayAllUsers)
	router.POST("/DisplayUser", controller.DisplayUser)

	// router.Run("127.0.0.1:8000")
	// router.Run("go-sirka.herokuapp.com:8080")
	router.Run("https://go-sirka.herokuapp.com:8080/")
}
