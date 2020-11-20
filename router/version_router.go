package router

import (
	"github.com/gin-gonic/gin"
	"ilmostro.org/gin-tutorial/configuration"
	"ilmostro.org/gin-tutorial/result"
)

var properties = configuration.Properties

func InitRouter(engine *gin.Engine) {

	v1 := engine.Group("/v1")
	v1.GET("/", func(context *gin.Context) {
		context.JSON(200, result.Success("Hello World!"))
	})
	v1.GET("/properties", func(context *gin.Context) {
		context.JSON(200, result.Success(properties))
	})
	v1.GET("/exception", func(context *gin.Context) {
		panic(result.InnerException)
	})
	UserInit(v1)
}
