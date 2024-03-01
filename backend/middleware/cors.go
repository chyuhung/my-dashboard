package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors 跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		cors.New(cors.Config{
			// AllowAllOrigins:        true,
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"*"},
			AllowHeaders: []string{"Origin"},
			// AllowCredentials: true,
			ExposeHeaders: []string{"Content-Length", "Authorization"},
			MaxAge:        12 * time.Hour,
		})
	}
}
