package controllers

import (
	"net/http"

	"github.com/chyuhung/my-dashboard/services"
	"github.com/gin-gonic/gin"
)

func GetFlavor(c *gin.Context) {}
func ListFlavors(c *gin.Context) {
	// 获取实例数据
	data, err := services.GetFlavors()
	if err != nil {
		// 处理获取实例数据失败的情况
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取flavor数据失败",
		})
		return
	}

	// 返回实例数据
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
