package router

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "mygin/docs" //引入docs文件 必须
	"mygin/internal/web"
	mid "mygin/middleware"
)

func SetupRouter() *gin.Engine {
	// Disable Console Color gin.DisableConsoleColor()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Redirect(302,"/swagger/index.html")
	})

	// ginSwagger middleware
	r.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "NAME_OF_ENV_VARIABLE"))

	url := r.Group("/api/v2")

	// 不经过权限认证的接口用 r 调用
	url.GET("/login",web.UserLogin())
	url.GET("/user", web.GetAll())





	// 要经过权限认证的接口用 authorized 调用

	// 用户接口鉴权
	userAuthorized := url.Group("/")
	userAuthorized.Use(mid.JWTAuth()).Use(mid.UserAuthentic())


	userAuthorized.GET("/user/:id", web.GetById())


	// 管理员接口鉴权
	adminAuthorized := url.Group("/")
	adminAuthorized.Use(mid.JWTAuth()).Use(mid.AdminAuthentic())

	// ....


	return r
}
