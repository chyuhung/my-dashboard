package middlewares

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var JwtKey = []byte("ddyyy12345")

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// SetToken 生成token
func SetToken(username string) (string, error) {
	expireTime := time.Now().Add(10 * time.Hour)
	SetClaims := MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ginblog",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	return reqClaim.SignedString(JwtKey)
}

// 验证token
func checkToken(tokenString string) *MyClaims {
	token, _ := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if token != nil {
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return claims
		}
	}
	return nil
}

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"message": "token 不存在",
			})
			c.Abort() // 终止剩余/后续中间件，当前中间件剩余代码会继续执行
			return    //终止当前中间件
		}
		userToken := strings.SplitN(tokenHeader, " ", 2)
		if len(userToken) != 2 && userToken[0] != "Bearer" {
			c.JSON(http.StatusOK, gin.H{
				"message": "token 格式错误",
			})
			c.Abort()
			return
		}
		key := checkToken(userToken[1])
		if key == nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "token 错误",
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			c.JSON(http.StatusOK, gin.H{
				"message": "token 过期",
			})
			c.Abort()
			return
		}

		c.Set("username", key.Username)
		c.Next() // 暂停当前中间件代码，执行后续中间件和控制器
	}
}
