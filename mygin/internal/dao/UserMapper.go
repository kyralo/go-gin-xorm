package dao

import (
	"github.com/go-xorm/xorm"
	"log"
	"mygin/common/utils"
	"mygin/internal/model"
)

var engine *xorm.Engine

func init() {
	engine = utils.GetMysqlEngine()
}

//查询所有用户
func ListUsers() []model.User{
	var users []model.User
	e := engine.Find(&users)
	if e != nil {
		log.Println("查询失败",e)
	}
	return users
}

func SelectUserById(id string) model.User{
	var user model.User
	_,e := engine.ID(id).Get(&user)
	if e != nil {
		log.Println("查询失败",e)
	}
	return user
}

func SelectUserByWeChatId(WechatId string) model.User {
	var user model.User

	_, e := engine.Where("wechat_id = ?", WechatId).Get(&user)

	if e != nil {
		log.Println("查询失败",e)
	}

	return user
}

func Insert(user model.User) int64 {
	i, e := engine.Insert(user)

	if e != nil {
		log.Println("新增失败",e)
	}

	return i
}

