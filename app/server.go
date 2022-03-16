package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	testEngine := engine.Group("/test")
	{
		v1 := testEngine.Group("/v1")
		{
			v1.GET("/test", func(ctx *gin.Context) {
				ctx.String(200, "Hello, World!")
			})
		}
	}

	err := engine.Run("127.0.0.1:3000")
	if err != nil {
		log.Fatal(err)
	}
}
