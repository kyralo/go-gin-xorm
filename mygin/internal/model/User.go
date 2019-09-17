package model

type User struct {
	Id int64 `json:"id" xorm:"not null pk INTEGER(11)"`
	Name string `json:"name" xorm:"null CHAR(32)"`
	Age int8 `json:"age" xorm:"null INTEGER(2)"`
	Sex string `json:"sex" xorm:"null CHAR(2)"`
}
