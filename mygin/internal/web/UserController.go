package web

import (
	"github.com/gin-gonic/gin"
	"mygin/internal/service"
)

// @Tags user
// @Summary 获取所有用户
// @Accept  json
// @Produce  json
// @Success 200 {array} model.User
// @Failure 404 {string} string "not found"
// @Failure 500 {string} string "server error"
// @Router /user [get]
func GetAll() gin.HandlerFunc{
	return func(c *gin.Context) {
		service.ListUsers(c)
	}
}

// @Tags user
// @Summary 获取用户信息通过 id
// @Accept  json
// @Accept  json
// @Param id query string true "用户id"
// @Success 200 {object} model.User
// @Failure 404 {string} string "not found"
// @Failure 500 {string} string "server error"
// @Router /user/{id} [get]
func GetById() gin.HandlerFunc{
	return func(c *gin.Context) {
		service.SelectUserById(c)
	}
}

