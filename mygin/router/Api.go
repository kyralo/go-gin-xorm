package router

import (
	"github.com/gin-gonic/gin"
	"mygin/internal/web"
	"net/http"
)

var db = make(map[string]string)

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		web.GetPing(c)
	})

	r.GET("/user", func(c *gin.Context) {
		web.GetAll(c)
	})

	r.GET("/user/:id", func(c *gin.Context) {
		web.GetById(c)
	})

	// Get user value
	//r.GET("/user/:name", func(c *gin.Context) {
	//	user := c.Params.ByName("name")
	//	value, ok := db[user]
	//	if ok {
	//		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	//	} else {
	//		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	//	}
	//})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))



	//权限认证
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}
