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
			"message": "搜索卷类型失败",
			"err":     err,
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

func GetVolumesHandler(c *gin.Context) {
	vls, err := services.GetVolumes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取卷列表失败",
			"err":     err,
		})
		return
	}
	var volumes []string
	for _, v := range vls {
		volumes = append(volumes, v.Name)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "获取卷列表成功",
		"data":    volumes,
	})
}

func CheckVolumeTypeHandler(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "查询参数不能为空",
		})
		return
	}
	vts, err := services.GetVolumeTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取卷类型列表失败",
			"err":     err,
		})
		return
	}
	for _, vt := range vts {
		if vt.Name == query {
			c.JSON(http.StatusOK, gin.H{
				"message": "卷类型可用",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "卷类型不可用",
	})
}
