package service

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"log"
	util "mygin/common/utils"
	"mygin/internal/dao"
	"mygin/internal/model"
	"net/http"
	"reflect"
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

	id := c.Params.ByName("id")
	user := dao.SelectUserById(id)

	//值判空
	if !reflect.DeepEqual(user,model.User{}) {
		c.JSONP(http.StatusOK,user)
	}else {
		c.JSON(http.StatusNotFound,gin.H{
			"message" : "404 not found",
		})
	}
}

func UserLogin(c *gin.Context) {
	code := c.Params.ByName("code")
	bodyStr := util.GetBodyStr(code)
	if bodyStr == nil {
		log.Println("微信openid请求获取失败")
		c.JSONP(http.StatusInternalServerError,"微信openid请求获取失败")
		c.Abort()
	}
	var openId string
	err := json.Unmarshal(bodyStr, openId)

	if err != nil {
		log.Println(err)
		c.JSONP(http.StatusInternalServerError,"服务器错误")
		c.Abort()
	}



	token := c.Request.Header.Get("token")
	if token == "" {
		//签发token
		Jwt := util.NewJWT()

		var claims = util.CustomClaims{}
		claims.Subject = "ROLE_USER"
		claims.Id = openId//用户wechatId

		token, _ := Jwt.CreateToken(claims)


		user := dao.SelectUserByWeChatId(openId)

		if !reflect.DeepEqual(user,model.User{}) {
			c.JSONP(http.StatusOK,gin.H{
				"token": token,
				"user": user,
			})
		}else {
			c.JSONP(http.StatusForbidden,"403 资源禁止访问")
		}
	}else {
		id := uuid.NewV4().Bytes()
		replace := bytes.Replace(id, []byte("-"), []byte(""), 0)

		user := model.User{}
		user.Id = string(replace)
		user.WechatId = openId

		insert := dao.Insert(user)

		if insert > 0 {
				c.JSONP(http.StatusOK,gin.H{"token": token})
		}else {
			c.JSONP(http.StatusInternalServerError,"服务端错误")
		}
	}
}
