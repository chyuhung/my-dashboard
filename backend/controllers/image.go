package controllers

import (
	"net/http"

	"github.com/chyuhung/my-dashboard/services"
	"github.com/gin-gonic/gin"
)

func GetImage(c *gin.Context) {}
func ListImages(c *gin.Context) {
	images, err := services.ListImages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取镜像数据失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"data": images,
	})
}
