package v1

import (
	"net/http"

	"github.com/chyuhung/my-dashboard/services"
	"github.com/gin-gonic/gin"
)

func GetNetworksHandler(c *gin.Context) {
	networks, err := services.GetNetworks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取network失败",
			"err":     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, networks)
}
