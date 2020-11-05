package main

import "github.com/gin-gonic/gin"
import "ilmostro.org/gin-tutorial/configuration"
import cRouter "ilmostro.org/gin-tutorial/router"

func main() {
	properties := configuration.GetProperties()
	gin.SetMode(properties.Server.Mode)
	router := gin.Default()
	cRouter.InitRouter(router)
	_ = router.Run(properties.Server.Port)
}
