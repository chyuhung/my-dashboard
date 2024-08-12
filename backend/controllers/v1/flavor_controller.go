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
			"err":     err.Error(),
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
