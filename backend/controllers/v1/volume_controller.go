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

func GetVolumesHandler(c *gin.Context) {
	vl, err := services.GetVolumes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取卷列表失败",
			"err":     err.Error(),
		})
		return
	}
	var volumes []string
	for _, v := range vl {
		volumes = append(volumes, v.Name)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "获取卷列表成功",
		"data":    vl,
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
			"message": "获取卷类型失败",
			"err":     err.Error(),
		})
		return
	}
	for _, vt := range vts {
		if vt.Name == query {
			c.JSON(http.StatusOK, gin.H{
				"message": "类型可用",
			})
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "类型不可用",
	})

}
