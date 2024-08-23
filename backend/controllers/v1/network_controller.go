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
			"err":     err,
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

func CheckNetworkHandler(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "查询参数不能为空",
		})
		return
	}
	nets, err := services.GetNetworks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取网络列表失败",
			"err":     err,
		})
		return
	}
	for _, net := range nets {
		if net.Name == query {
			c.JSON(http.StatusOK, gin.H{
				"message": "网络可用",
				"data":    net,
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "网络不可用",
	})
}
