package router

import (
	"github.com/gin-gonic/gin"
	"mygin/internal/web"
	mid "mygin/middleware"
)

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()


	//不经过权限认证的接口用 r 调用
	r.GET("/user", web.GetAll())


	//要经过权限认证的接口用 authorized 调用
	authorized := r.Group("/")
	authorized.Use(mid.JWTAuth())

	authorized.GET("/user/:id", web.GetById())

	return r
}
