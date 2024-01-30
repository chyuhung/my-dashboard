package controllers

import (
	"net/http"

	"github.com/chyuhung/my-dashboard/services"
	"github.com/gin-gonic/gin"
)

func GetVolume(c *gin.Context) {}
func ListVolumes(c *gin.Context) {
	volumes, err := services.ListVolumes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取卷数据失败",
		})
		return
	}

	// 返回实例数据
	c.JSON(http.StatusOK, gin.H{
		"data": volumes,
	})
}
