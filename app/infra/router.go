package infra

import (
	"media-gin/app/interfaces/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	userController := controllers.NewUserController(NewFirestoreHandler())

	router.GET("/users", func(c *gin.Context) { userController.Index(c) })
	router.GET("/users/:id", func(c *gin.Context) { userController.Show(c) })
	router.POST("/users/create", func(c *gin.Context) { userController.Create(c) })
	router.POST("/users/update", func(c *gin.Context) { userController.Update(c) })
	router.POST("/users/delete/:id", func(c *gin.Context) { userController.Delete(c) })

	Router = router
}
