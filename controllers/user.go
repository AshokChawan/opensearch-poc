package controllers

import (
	"github.com/dashboard_poc.com/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UserController struct{}

// var userModel = new(models.User)

func (u UserController) Retrieve(c *gin.Context) {
	userModel := &models.User{
		ID:       "123",
		Name:     "test",
		BirthDay: "01-01.2020",
		Gender:   "M",
		PhotoURL: "www.test.com",
		Time:     time.Now().UnixNano(),
	}
	c.JSON(http.StatusOK, gin.H{"message": "User found!", "user": userModel})
}
