package router

import (
	"github.com/gin-gonic/gin"
	"ilmostro.org/gin-tutorial/configuration"
	"ilmostro.org/gin-tutorial/repository"
)

func InitRouter(engine *gin.Engine) {

	v1 := engine.Group("/v1")

	v1.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	properties := configuration.GetProperties()

	v1.GET("/properties", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"properties": properties,
		})
	})

	person := repository.Student{Id: 1, Name: "ilmsotro", Age: 18}
	person1 := new(repository.Student)
	person1.Id = 1
	person1.Age = 1
	person1.Name = "ilmostro"
	v1.GET("/person", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"eat": person1.Eat(),
			"run": person.Run(),
		})
	})

	v1.GET("/db", func(context *gin.Context) {
		students := repository.GetAllUserFromDB()
		context.JSON(200, gin.H{
			"students": students,
		})
	})
}
