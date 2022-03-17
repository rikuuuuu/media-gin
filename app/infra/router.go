package infra

import (
	"media-gin/app/interfaces/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	userController := controllers.NewUserController(NewFirestoreHandler())
	articleController := controllers.NewArticleController(NewFirestoreHandler())

	// user
	router.GET("/users", func(c *gin.Context) { userController.Index(c) })
	router.GET("/user/:id", func(c *gin.Context) { userController.Show(c) })
	router.POST("/user/create", func(c *gin.Context) { userController.Create(c) })
	router.POST("/user/update", func(c *gin.Context) { userController.Update(c) })
	router.POST("/user/delete/:id", func(c *gin.Context) { userController.Delete(c) })

	// article
	router.GET("/articles", func(c *gin.Context) { articleController.Index(c) })
	router.GET("/article/:id", func(c *gin.Context) { articleController.Show(c) })
	router.POST("/article/create", func(c *gin.Context) { articleController.Create(c) })
	router.POST("/article/update", func(c *gin.Context) { articleController.Update(c) })
	router.POST("/article/delete/:id", func(c *gin.Context) { articleController.Delete(c) })

	Router = router
}
