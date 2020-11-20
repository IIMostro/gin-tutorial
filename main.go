package main

import (
	"github.com/gin-gonic/gin"
	cRouter "ilmostro.org/gin-tutorial/router"
	"log"
)
import "ilmostro.org/gin-tutorial/configuration"

var properties = configuration.Properties.Server

func main() {
	gin.SetMode(properties.Mode)
	router := gin.Default()
	cRouter.InitRouter(router)
	log.Printf("server start model: %s, port%s", properties.Mode, properties.Port)
	_ = router.Run(properties.Port)
}
