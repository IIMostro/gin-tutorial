package router

import (
	"github.com/gin-gonic/gin"
	"ilmostro.org/gin-tutorial/repository"
	"ilmostro.org/gin-tutorial/result"
	"ilmostro.org/gin-tutorial/services"
	"log"
)

func UserInit(group *gin.RouterGroup) {

	userRouter := group.Group("/user")

	userRouter.GET("/", func(context *gin.Context) {
		students := services.NewRedisUserService().GetStudents()
		context.JSON(200, result.Success(students))
	})
	userRouter.GET("/:id", func(context *gin.Context) {
		id := context.Param("id")
		student := services.NewRedisUserService().GetStudent(id)
		context.JSON(200, result.Success(student))
	})

	userRouter.POST("/", func(context *gin.Context) {
		var student repository.Student
		err := context.BindJSON(&student)
		if err != nil {
			log.Printf("bind student error!, cause: %f", err)
			return
		}
		services.NewRedisUserService().Save(student)
	})
}
