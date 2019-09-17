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
	users := new([]model.User)
	e := engine.Find(users)
	if e != nil {
		log.Fatalf("查询失败",e)
	}
	return *users
}

func SelectUserById(id int64) model.User{
	user := new(model.User)
	_,e := engine.ID(id).Get(user)
	if e != nil {
		log.Fatalf("查询失败",e)
	}
	return *user
}

