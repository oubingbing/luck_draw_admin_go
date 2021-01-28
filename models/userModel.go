package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"luck-admin/enums"
)

const (
	USER_FROM_MINI			= 1		//小程序
	USER_FRMO_OFFICIAL		= 2		//公众号
)
const (
	USER_GENDER_BOY			= 1		//男
	USER_GENDER_GIRL		= 2		//女
	USER_GENDER_OTHER		= 3     //其他
)
const (
	USER_STATUS_WHITE		= 1		//正常
	USER_STATUS_BLACK		= 2		//拉黑
)

const (
	FAKER_N 						= 0		//真用户
	FAKER_Y 						= 1		//假用户
)

type UserPage []*enums.UserPage

type User struct {
	gorm.Model
	NickName 		string		`gorm:"column:nick_name"`		//昵称
	AvatarUrl		string		`gorm:"column:avatar_url"`		//头像
	Gender			int8		`gorm:"column:gender"`			//性别
	OpenId			string		`gorm:"column:open_id"`			//openid
	UnionId			string		`gorm:"column:union_id"`
	City			string		`gorm:"column:city"`
	Country			string		`gorm:"column:country"`
	Language		string		`gorm:"column:language"`
	Province		string		`gorm:"column:province"`
	FromType		int8		`gorm:"column:from_type"`		//用户来源,1=小程序，2=h5公众号
	Status			int8		`gorm:"column:status"`			//状态，1=正常，2=拉黑
	Phone			string		`gorm:"column:phone"`
	Faker			int8		`gorm:"column:faker"`
}

func (User) TableName() string  {
	return "wechat_user"
}

func (user *User)Page(db *gorm.DB,page *PageParam) (*UserPage,error) {
	var userPage UserPage
	err := Page(db,user.TableName(),page).
		Select("id,nick_name,avatar_url,gender,open_id,city,country,language,province,from_type,status").
		Order(fmt.Sprintf("'%v' '%v'",page.OrderBY,page.Sort)).
		Find(&userPage).Error
	return &userPage,err
}
