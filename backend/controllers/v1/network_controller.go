package v1

import (
	"net/http"

	"github.com/chyuhung/my-dashboard/services"
	"github.com/gin-gonic/gin"
)

func GetNetworksHandler(c *gin.Context) {
	nets, err := services.GetNetworks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取network失败",
			"err":     err.Error(),
		})
		return
	}
	var networks []string
	for _, net := range nets {
		networks = append(networks, net.Name)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "获取network成功",
		"data":    networks,
	})
}
