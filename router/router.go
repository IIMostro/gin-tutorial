package router

import (
	"github.com/gin-gonic/gin"
)
import "ilmostro.org/gin-tutorial/configuration"
import "ilmostro.org/gin-tutorial/services"

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

	person := services.Student{Id: 1, Name: "ilmsotro", Age: 18}
	person1 := new(services.Student)
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
		connection := configuration.GetConnection()
		students := services.GetAllUserFromDB(connection)
		context.JSON(200, gin.H{
			"students": students,
		})
		connection.Close()
	})
	//for _, student := range students{
	//	log.Printf("get database user: id:%d, name:%s, age:%d", student.Id, student.Name, student.Age)
	//}
}
