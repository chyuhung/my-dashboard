package v1

import (
	"net/http"

	"github.com/chyuhung/my-dashboard/services"
	"github.com/gin-gonic/gin"
)

func GetFlavorsHandler(c *gin.Context) {
	fvs, err := services.GetFlavors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取flavor失败",
			"err":     err,
		})
		return
	}
	var flavors []string
	for _, fv := range fvs {
		flavors = append(flavors, fv.Name)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "获取flavor成功",
		"data":    flavors,
	})
}

func CheckFlavorHandler(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "查询参数不能为空"})
		return
	}

	flavors, err := services.GetFlavors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "搜索失败",
			"error":   err,
		})
		return
	}
	for _, f := range flavors {
		if f.Name == query {
			c.JSON(http.StatusOK, gin.H{
				"message": "规格可用",
			})
			return
		}

	}
	c.JSON(http.StatusOK, gin.H{
		"message": "规格不可用",
	})
}
