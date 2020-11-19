package router

import (
	"github.com/gin-gonic/gin"
	"ilmostro.org/gin-tutorial/repository"
	"ilmostro.org/gin-tutorial/result"
	"ilmostro.org/gin-tutorial/services"
)

func UserInit(group *gin.RouterGroup) {

	userRouter := group.Group("/user")

	userRouter.GET("/", func(context *gin.Context) {
		students := services.GetStudents()
		context.JSON(200, result.Success(students))
	})
	userRouter.GET("/:id", func(context *gin.Context) {
		param := context.Param("id")
		student := repository.GetStudentById(param)
		context.JSON(200, result.Success(student))
	})
}
