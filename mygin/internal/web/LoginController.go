package web

import (
	"github.com/gin-gonic/gin"
	"mygin/internal/service"
)

func UserLogin() gin.HandlerFunc{
	return func(c *gin.Context) {
		service.UserLogin(c)
	}
}