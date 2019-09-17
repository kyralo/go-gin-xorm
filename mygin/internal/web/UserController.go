package web

import (
	"github.com/gin-gonic/gin"
	"mygin/internal/service"
)

func GetAll(c *gin.Context) {
	service.ListUsers(c)
}

func GetById(c *gin.Context){
	service.SelectUserById(c)
}

