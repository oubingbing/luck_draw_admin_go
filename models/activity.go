package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"luck-admin/enums"
	"time"
)

//活动类型，1=红包，2=商品，3=话费
const (
	ACTIVITY_TYPE_RED_PAK 		= 1
	ACTIVITY_TYPE_GOODS   		= 2
	ACTIVITY_TYPE_PHONE_BILL   	= 3
)

//活动发布来源，1=平台，2=用户
const (
	ACTIVITY_FROM_PLATFORM 	= 1
	ACTIVITY_FROM_USER	    = 2
)

//是否限制参加人数，0=不限制，1=限制
const (
	ACTIVITY_LIMIT_JOIN_N = 0
	ACTIVITY_LIMIT_JOIN_Y = 1
)

//活动状态，1=待发布，2=进行中，3=已下架，4=已结束
const (
	ACTIVITY_STATSUS_TO_RELE 	= 1
	ACTIVITY_STATSUS_RUNNING 	= 2
	ACTIVITY_STATSUS_DOWN 		= 3
	ACTIVITY_STATSUS_FINISH 	= 4

)

//抽奖方式，1=平均，2=拼手气，20%中奖,4=定时
const (
	ACTIVITY_DRAW_TYPE_AVERAGE  = 1
	ACTIVITY_DRAW_TYPE_RAND     = 2
	ACTIVITY_DRAW_TYPE_RAND_all = 3
	ACTIVITY_DRAW_TYPE_TIME		= 4
)

//是否真的送奖品，0=否，1=是 really
const (
	ACTIVITY_REALLY_N 			= 0
	ACTIVITY_REALLY_Y 			= 1
)

//是否大图，1=大图，2=小图
const (
	ACTIVITY_BIG_PIC_BIG		= 1
	ACTIVITY_BIG_PIC_SMALL		= 2
)

//是否开启广告,是否开启广告，0=否，1=是
const (
	ACTIVITY_OP_AD_Y 			= 1
	ACTIVITY_OP_AD_N 			= 0
)

//是否置顶，0=否，1=是
const (
	ACTIVITY_IS_TOP_Y 			= 1
	ACTIVITY_IS_TOP_N 			= 0
)

//假用户redis key
const FAKER_USER_KEY = "luck_draw_faker_activity"
const TOP_ACTIVITY 	 = "luck_draw_top_activity"

type Activity struct {
	gorm.Model
	Name 			string 		`gorm:"column:name"`
	GiftId 			int64 		`gorm:"column:gift_id"`
	Type 			int8   		`gorm:"column:type"` 			//活动类型
	OpenAd 			int8   		`gorm:"column:open_ad"` 		//是否开启广告
	FromType 		int8   		`gorm:"column:from_type"` 		//发布活动的用户类型
	JoinNum 		int32 		`gorm:"column:join_num"`   		//已参加人数
	LimitJoin 		int32 	 	`gorm:"column:limit_join"`  	//是否限制参加人数
	JoinLimitNum 	float32 	`gorm:"column:join_limit_num"` 	//限制参加人数
	ReceiveLimit 	float32 	`gorm:"column:receive_limit"` 	//每人限领数量
	Des 			string      `gorm:"column:des"`
	Attachments 	string   	`gorm:"column:attachments"`
	Status 			int8		`gorm:"column:status"` 			//活动状态
	ShareTitle 		string    	`gorm:"column:share_title"` 	//分享标题
	ShareImage 		string    	`gorm:"column:share_image"` 	//分享图片
	DrawType 		int8    	`gorm:"column:draw_type"`		//抽奖方式
	Really 			int8    	`gorm:"column:really"`			//是否真抽
	Consume			float32		`gorm:"column:consume"`			//礼品消耗数
	BigPic 			int8    	`gorm:"column:big_pic"`			//是否大图
	IsTop 			int8    	`gorm:"column:is_top"`			//是否置顶
	Number			string		`gorm:"column:number"`			//活动编号
}

type AcPage []ActivityPageFormat
type AcTop []ActivityTop

var pageErr error = errors.New("查询出错")

func (Activity) TableName() string  {
	return "activity"
}

func (activity *Activity) CountToday(db *gorm.DB) (int64,error) {
	var num int64
	err := db.Table(activity.TableName()).
		//Set("gorm:query_option", "FOR UPDATE").
		Not("status", []int8{JOIN_LOG_STATUS_FAIL,JOIN_LOG_STATUS_QUEUE}).
		Where("created_at  >= ?",time.Now().Format(enums.DATE_ONLY_FORMAT)).
		Where("created_at  <= ?",time.Now().Format(enums.DATE_FORMAT)).
		Where("deleted_at is null").
		Count(&num).Error

	return num,err
}

func (activity *Activity)Store(db *gorm.DB) (int64,error) {
	createResult := db.Create(activity)
	return createResult.RowsAffected,createResult.Error
}

func (activity *Activity)Page(db *gorm.DB,page *PageParam) (AcPage,*enums.ErrorInfo) {
	var activities AcPage
	err :=  Page(db,activity.TableName(),page).
			Where("activity.deleted_at is null").
			Joins("left join gift on gift.id = activity.gift_id").
			Select("activity.id,activity.name,gift_id,activity.type,share_title,share_image,activity.from_type,join_num,join_limit_num,receive_limit,activity.des,activity.attachments,activity.status,gift.name as gift_name,big_pic,draw_type,really").
			Order("id desc").
			Find(&activities).Error
	if err != nil {
		fmt.Printf("数据错误：%v\n",err)
		return nil,&enums.ErrorInfo{pageErr,enums.ACTIVITY_PAGE_ERR}
	}

	return activities,nil
}

func (activity *Activity) Detail(db *gorm.DB,id string) (*enums.ActivityDetailFormat,bool,error,) {
	activityDetail := &enums.ActivityDetailFormat{}
	err := db.Table(activity.TableName()).
		Select("id,name,gift_id,type,from_type,join_num,limit_join,join_limit_num,des,attachments,share_title,share_image,created_at").
		Where("id = ?",id).
		Where("deleted_at is null").
		First(activityDetail).Error
	return activityDetail,db.RecordNotFound(),err
}

func (activity *Activity)LockById(db *gorm.DB,id string) (bool,error) {
	err := db.Table(activity.TableName()).
		Set("gorm:query_option", "FOR UPDATE").
		Where("id = ?",id).
		Where("deleted_at is null").
		First(activity).Error

	return db.RecordNotFound(),err
}

func (activity *Activity)Delete(db *gorm.DB) error {
	err := db.Delete(activity).Error
	return err
}

func (activity *Activity) Up(db *gorm.DB,id interface{}) error {
	update := map[string]interface{}{
		"status":ACTIVITY_STATSUS_RUNNING,
	}
	err := db.Table(activity.TableName()).
		Where("id = ?",id).
		Where("deleted_at is null").
		Update(update).Error
	return err
}

func (activity *Activity) Down(db *gorm.DB,id interface{}) error {
	update := map[string]interface{}{
		"status":ACTIVITY_STATSUS_DOWN,
	}
	err := db.Table(activity.TableName()).
		Where("id = ?",id).
		Where("deleted_at is null").
		Update(update).Error
	return err
}

func (activity *Activity) FindById(db *gorm.DB,id interface{}) error {
	err := db.Table(activity.TableName()).
		Where("id = ?",id).
		Where("deleted_at is null").
		First(activity).Error
	return err
}

func (activity *Activity) Tops(db *gorm.DB) (AcTop,error) {
	var activities AcTop
	err :=  db.Table(activity.TableName()).
		Select("id,name,is_top,number,gift_id,type,from_type,join_num,attachments,join_limit_num,status,created_at").
		Where("deleted_at is null").
		Where("is_top = ?",ACTIVITY_IS_TOP_Y).
		Where("status in (?)",[]int8{ACTIVITY_STATSUS_RUNNING}).
		Order("id desc").
		Find(&activities).Error
	return activities,err
}
