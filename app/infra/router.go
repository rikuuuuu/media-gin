package infra

import (
	"media-gin/app/interfaces/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	userController := controllers.NewUserController()

	router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	router.GET("/users", func(c *gin.Context) { userController.Index(c) })
	router.GET("/users/:id", func(c *gin.Context) { userController.Show(c) })

	Router = router
}
