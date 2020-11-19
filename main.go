package main

import (
	"github.com/gin-gonic/gin"
	cRouter "ilmostro.org/gin-tutorial/router"
)
import "ilmostro.org/gin-tutorial/configuration"

var properties = configuration.Properties

func main() {
	gin.SetMode(properties.Server.Mode)
	router := gin.Default()
	cRouter.InitRouter(router)
	_ = router.Run(properties.Server.Port)
}
