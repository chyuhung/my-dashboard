package v1

import (
	"net/http"

	"github.com/chyuhung/my-dashboard/services"
	"github.com/gin-gonic/gin"
)

func GetFlavorsHandler(c *gin.Context) {
	flavors, err := services.GetFlavors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取flavor失败",
			"err":     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, flavors)
}
