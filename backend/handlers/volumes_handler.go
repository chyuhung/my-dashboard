package handlers

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
)

func VolumesHandler(volumes []volumes.Volume) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 渲染HTML模板
		tmpl, err := template.ParseFiles("templates/volumes.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to load template")
			return
		}

		err = tmpl.Execute(c.Writer, volumes)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to render template")
			return
		}
	}
}
