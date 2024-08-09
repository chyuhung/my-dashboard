package v1

import (
	"net/http"

	"github.com/chyuhung/my-dashboard/services"
	"github.com/gin-gonic/gin"
)

func GetImagesHandler(c *gin.Context) {
	images, err := services.GetImages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取image失败",
			"err":     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, images)
}
