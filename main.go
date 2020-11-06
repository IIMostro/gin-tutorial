package main

import (
	"github.com/gin-gonic/gin"
	"ilmostro.org/gin-tutorial/repository"
)
import "ilmostro.org/gin-tutorial/configuration"
import cRouter "ilmostro.org/gin-tutorial/router"

func main() {
	properties := configuration.GetProperties()
	gin.SetMode(properties.Server.Mode)
	router := gin.Default()
	cRouter.InitRouter(router)
	repository.Setup()
	_ = router.Run(properties.Server.Port)
}
