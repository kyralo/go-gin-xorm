package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mygin/internal/service"
	"net/http"
	"strconv"
)

func GetAll(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{"userList" : service.ListUsers()})
}

func GetById(c *gin.Context){
	fmt.Print(c.Params.ByName("id"))
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	c.JSON(http.StatusOK,gin.H{
		"user": service.SelectUserById(int64(id)),
	})
}

