package controllers

import (
	"net/http"

	"github.com/chyuhung/my-dashboard/services"
	"github.com/gin-gonic/gin"
)

func GetNetwork(c *gin.Context) {}
func ListNetworks(c *gin.Context) {
	networks, err := services.ListNetworks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取网络数据失败",
		})
		return
	}

	// 返回实例数据
	c.JSON(http.StatusOK, gin.H{
		"data": networks,
	})
}
