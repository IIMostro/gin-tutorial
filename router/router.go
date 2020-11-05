package router

import "github.com/gin-gonic/gin"
import "ilmostro.org/gin-tutorial/configuration"
import "ilmostro.org/gin-tutorial/services"

func InitRouter(engine *gin.Engine) {

	engine.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	properties := configuration.GetProperties()

	engine.GET("/properties", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"properties": properties,
		})
	})

	person := services.Student{Name: "ilmsotro", Age: 18}
	engine.GET("/person", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"eat": person.Eat(),
			"run": person.Run(),
		})
	})
}
