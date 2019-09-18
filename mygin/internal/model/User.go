package model

type User struct {
	Id string `json:"id" xorm:"id not null pk CHAR(32)"`//pk键必须声明,否则CURD失效
	WechatId string `json:"wechat_id" xorm:"wechat_id not null CHAR(32)"`
	Name string `json:"name" xorm:"name null CHAR(32)"`
	Age int8 `json:"age" xorm:"age null INTEGER(2)"`
	Sex string `json:"sex" xorm:"sex null CHAR(2)"`
}
