package v1

import (
	"net/http"

	"github.com/chyuhung/my-dashboard/services"
	"github.com/gin-gonic/gin"
)

func GetVolumeTypesHandler(c *gin.Context) {
	vts, err := services.GetVolumeTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取卷类型失败",
			"err":     err.Error(),
		})
		return
	}
	var volumeTypes []string
	for _, vt := range vts {
		volumeTypes = append(volumeTypes, vt.Name)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "获取卷类型成功",
		"data":    volumeTypes,
	})
}
