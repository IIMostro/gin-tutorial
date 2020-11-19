package main

import (
	"github.com/gin-gonic/gin"
	"ilmostro.org/gin-tutorial/repository"
	cRouter "ilmostro.org/gin-tutorial/router"
)
import "ilmostro.org/gin-tutorial/configuration"

func main() {
	properties := configuration.GetProperties()
	gin.SetMode(properties.Server.Mode)
	repository.Setup()
	router := gin.Default()
	cRouter.InitRouter(router)
	_ = router.Run(properties.Server.Port)
}
