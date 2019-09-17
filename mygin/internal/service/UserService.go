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
		c.JSONP(http.StatusOK,users)
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
		c.JSONP(http.StatusOK,user)
	}else {
		c.JSON(http.StatusNotFound,gin.H{
			"message" : "404 not found",
		})
	}
}
