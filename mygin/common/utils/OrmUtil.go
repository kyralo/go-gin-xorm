package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

var engine *xorm.Engine

func GetMysqlEngine() *xorm.Engine{
	var err error
	engine, err = xorm.NewEngine("mysql", "root:991010@tcp(127.0.0.1:3306)/go?charset=utf8")

	if err != nil {
		log.Fatalf(" 数据库链接失败: ",err)
		return nil
	}

	return engine

}
