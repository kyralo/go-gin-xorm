package model


//@Param //id body string true "用户id" //pk键必须声明,否则CURD失效
//@Param //name body string true "用户名"
//@Param //age body int true "用户年龄"
//@Param //sex body string true "用户性别"
type User struct {
	Id string `json:"id" xorm:"id not null pk CHAR(32)"` //用户id
	WechatId string `json:"wechat_id" xorm:"wechat_id not null CHAR(32)"` //微信id
	Name string `json:"name" xorm:"name null CHAR(32)"` //用户名
	Age int8 `json:"age" xorm:"age null INTEGER(2)"` //用户年龄
	Sex string `json:"sex" xorm:"sex null CHAR(2)"` //用户性别
}
