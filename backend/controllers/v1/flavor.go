package controllers

import (
	"net/http"

	"github.com/chyuhung/my-dashboard/services"
	"github.com/gin-gonic/gin"
)

func GetFlavor(c *gin.Context) {
	// 获取flavor id
	id := c.Param("id")

	// 获取flavor数据
	data, err := services.GetFlavor(id)
	if err != nil {
		// 处理获取flavor数据失败的情况
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取flavor数据失败",
		})
		return
	}

	// 返回flavor数据
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
func ListFlavors(c *gin.Context) {
	// 获取规格数据
	data, err := services.GetFlavors()
	if err != nil {
		// 处理获取规格数据失败的情况
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取flavor数据失败",
		})
		return
	}

	// 返回规格数据
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
