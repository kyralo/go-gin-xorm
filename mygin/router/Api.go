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
	r.GET("/",web.UserLogin())

	r.GET("/user", web.GetAll())
	//r.GET("/user/:id",web.GetById())


	//要经过权限认证的接口用 authorized 调用

	//用户接口鉴权
	userAuthorized := r.Group("/")
	userAuthorized.Use(mid.JWTAuth()).Use(mid.UserAuthentic())

	userAuthorized.GET("/user/:id", web.GetById())


	//管理员接口鉴权
	adminAuthorized := r.Group("/")
	adminAuthorized.Use(mid.JWTAuth()).Use(mid.AdminAuthentic())

	// ....


	return r
}
