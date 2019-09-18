package middleware

import (
	"github.com/gin-gonic/gin"
	JWT "mygin/common/utils"
	"net/http"
)

/**
 * \* @author: WangChen
 * \* Date: 19-9-18
 * \* Time: 上午9:23
 */

func AdminAuthentic() gin.HandlerFunc{
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")

		if token == "" {
			c.JSONP(http.StatusUnauthorized,"401 未授权")
			c.Abort()
		}

		jwt := JWT.NewJWT()

		claims, _ := jwt.ParseToken(token)

		if claims.Subject == "ROLE_ADMIN" {
			c.Next()
		}else {
			c.JSONP(http.StatusForbidden,"403 禁止访问")
		}
	}
}