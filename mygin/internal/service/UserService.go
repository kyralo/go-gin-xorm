package service

import (
	"github.com/gin-gonic/gin"
	"mygin/internal/dao"
	"mygin/internal/model"
	"net/http"
	"reflect"
	"strconv"
)

//ListUsers
func ListUsers(c *gin.Context)  {
	users := dao.ListUsers()

	if users != nil {
		c.JSON(http.StatusOK,gin.H{
			"userList" : users,
		})
	}else {
		c.JSON(http.StatusNotFound,gin.H{
			"message" : "404 not found",
		})
	}
}

func SelectUserById(c *gin.Context){

	id, _ := strconv.Atoi(c.Params.ByName("id"))

	user := dao.SelectUserById(int64(id))

	//值判空
	if !reflect.DeepEqual(user,model.User{}) {
		c.JSON(http.StatusOK,gin.H{
			"user": user,
		})
	}else {
		c.JSON(http.StatusNotFound,gin.H{
			"message" : "404 not found",
		})
	}
}
