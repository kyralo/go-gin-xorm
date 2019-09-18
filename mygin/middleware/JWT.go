package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	JwtUtil "mygin/common/utils"
	"net/http"
	"time"
)

// JWT验证
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("claims")
		if err != nil {
			c.JSON(http.StatusUnauthorized,gin.H{
				"message" : "401 Unauthorized",
			})
			log.Println(401,JwtUtil.TokenLose,time.Now())

			c.Abort()//中断请求
			return
		}

		j := JwtUtil.NewJWT()
		// token := t.(string)
		// claims, err := j.ParseToken(token)
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == JwtUtil.TokenExpired {
				if token, err = j.RefreshToken(token); err == nil {
					c.JSON(http.StatusOK, gin.H{"error": 0, "message": "refresh token", "token": token})
					return
				}
			}

			c.JSON(http.StatusForbidden,gin.H{
				"message" : "403 Forbidden",
			})

			log.Println(403,JwtUtil.TokenInvalid,time.Now())
			c.Abort()//中断请求
			return
		}

		c.Set("claims", claims)
		c.Next()
		return
	}
}


