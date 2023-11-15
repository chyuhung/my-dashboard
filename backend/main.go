package main

import (
	"github.com/chyuhung/my-dashboard/config"
	"github.com/chyuhung/my-dashboard/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config.Init()

	routes.Setup(router)

	router.Run(":8080")
}
