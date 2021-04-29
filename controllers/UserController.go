package controllers

import (
	"bwastartup/config"
	"bwastartup/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context){

	var User []models.Users

	if err := config.DB.Find(&User).Error;err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":"Record Not Found"})
			return
	}
	c.JSON(http.StatusOK, User)
}