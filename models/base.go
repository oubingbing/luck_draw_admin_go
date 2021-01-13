package models

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/jinzhu/gorm"
	"luck-admin/enums"
	"luck-admin/util"
	"strings"
)

var (
	orm *gorm.DB
	err error
)

type PageParam struct {
	PageNum 		int 		`form:"page_num" json:"page_num" binding:"required"`
	PageSize 		int  		`form:"page_size" json:"page_size" binding:"required"`
	OrderBY 		string 	 	`form:"order_by" json:"order_by" binding:"required"`
	Sort			string 		`form:"sort" json:"sort" binding:"required"` 							//分享图片
}

func Init(c db.Connection) {
	orm, err = gorm.Open("mysql", c.GetDB("default"))

	if err != nil {
		panic("initialize orm failed")
	}
}

func GetMysqlConfig() (string) {
	cfg := config.ReadFromJson("./config.json")
	databaseConfig := cfg.Databases.GetDefault()

	var builder strings.Builder
	builder.WriteString(databaseConfig.User)
	builder.WriteString(":")
	builder.WriteString(databaseConfig.Pwd)
	builder.WriteString("@(")
	builder.WriteString(databaseConfig.Host)
	builder.WriteString(":")
	builder.WriteString(databaseConfig.Port)
	builder.WriteString(")/")
	builder.WriteString(databaseConfig.Name)
	builder.WriteString("?charset=utf8")
	builder.WriteString("&parseTime=True&loc=Local")
	return builder.String()
}

func Connect() (*gorm.DB,*enums.ErrorInfo) {
	errorInfo := &enums.ErrorInfo{}
	orm, err = gorm.Open("mysql", GetMysqlConfig())
	if err != nil {
		util.Info(fmt.Sprintf("连接数据库错误：%v\n",err.Error()))
		errorInfo.Err = enums.ConnectErr
		errorInfo.Code = enums.DB_CONNECT_ERR
		return nil,errorInfo
	}

	orm.LogMode(true)

	return  orm,nil
}

/**
 * 通用分页
 */
func Page(db *gorm.DB,table string,page *PageParam) *gorm.DB {
	return db.Table(table).Limit(page.PageSize).Offset((page.PageNum-1)*page.PageSize)
}
