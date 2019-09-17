package web

import (
	"github.com/gin-gonic/gin"
	"mygin/internal/service"
)

func GetAll() gin.HandlerFunc{
	return func(c *gin.Context) {
		service.ListUsers(c)
	}
}

func GetById() gin.HandlerFunc{
	return func(c *gin.Context) {
		service.SelectUserById(c)
	}
}

