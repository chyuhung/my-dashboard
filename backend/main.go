package backend

import (
	"github.com/chyuhung/my-dashboard/config"
	"github.com/chyuhung/my-dashboard/routes"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	config.Init()

	routes.Setup(router)

	router.Run(":8080")
}
