package controller

import (
	// "fmt"
	"app/config"
	"app/models"
	"app/request"

	"github.com/gin-gonic/gin"
)

func DisplayAllUsers(c *gin.Context) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": "Data not found !"})
		return
	}

	c.JSON(200, gin.H{"status": true, "data": users, "message": nil})
}

func DisplayUser(c *gin.Context) {
	var valid request.DisplayUser
	if err := c.ShouldBind(&valid); err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.First(&user, "id = ?", valid.Id).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": "Data not found !"})
		return
	}

	c.JSON(200, gin.H{"status": true, "data": user, "message": nil})
}
