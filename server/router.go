package server

import (
	"github.com/dashboard_poc.com/controllers"
	// "github.com/dashboard_poc.com/middlewares"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)
	user := new(controllers.UserController)
	upload := new(controllers.UploadController)
	router.GET("/health", health.Status)
	router.GET("/user/id/:id", user.Retrieve)
	router.POST("/upload", upload.UploadData)
	// router.Use(middlewares.AuthMiddleware())
	return router

}
