package models

import (
	"github.com/jinzhu/gorm"
	"luck-admin/enums"
	"time"
)

//奖品类型，1=红包，2=商品，3=话费
const (
	GIFT_TYPE_RED_PAK 		= 1
	GIFT_TYPE_GOODS   		= 2
	GIFT_TYPE_PHONE_BILL 	= 3
)

//奖品来源，1=平台，2=用户
const (
	GIFT_FROM_PLATFORM 		= 1
	GIFT_FROM_USER	    	= 2
)

//奖品状态，1=上架，2=下架，下架不可用
const (
	GIFT_STATUS_UP			= 1
	GIFT_STATUS_DOWN		= 2
)

type Gift struct {
	gorm.Model
	Name 		string 		`gorm:"column:name"`
	UserId 		int 		`gorm:"column:user_id"`
	Num 		float64 	`gorm:"column:num"`
	Type 		int8   		`gorm:"column:type"`			//奖品类型，1=红包，2=商品，3=话费
	FROM        int8   		`gorm:"column:from_type"`  		//奖品来源，1=平台，2=用户
	STATUS      int8   		`gorm:"column:status"`			//奖品状态，1=上架，2=下架，下架不可用
	Des    		string      `gorm:"column:des"`
	Attachments string  	`gorm:"column:attachments"`
}

//活动分页
type ActivityPageFormat struct {
	ID        		uint
	Name 			string
	GiftId 			int64
	Type 			int8   		 	//活动类型
	TypeStr 		string   		 //活动类型
	OpenAd 			int8   		 	//是否开启广告
	OpenAdStr 		string   		//是否开启广告
	FromType 		int32   		//发布活动的用户类型
	JoinNum 		int32 		   	//已参加人数
	JoinLimitNum 	float32 	 	//限制参加人数
	Status 			int8		 	//活动状态
	StatusStr 		string		 	//活动状态
	Gift			*Gift
	GiftName		string
	Attachments		string
	AttachmentsStr	[]string
	ShareImage		string
	ShareImageStr	[]string
	ShareTitle		string
	LimitJoin 		int32 	 	  	//是否限制参加人数
	ReceiveLimit 	float32 	 	//每人限领数量
	Des 			string
	DrawType 		int8
	DrawTypeStr 	string
	Really 			int8
	ReallyStr 		string
	Consume			float32
	BigPic 			int8
	BigPicStr 		string
}

type ActivityTop struct {
	ID        		uint
	Name 			string
	GiftId 			int64
	Type 			int8   		 	//活动类型
	FromType 		int32   		 //发布活动的用户类型
	JoinNum 		int32 		   	//已参加人数
	JoinLimitNum 	float32 	 	//限制参加人数
	Attachments 	string
	AttachmentsSli 	[]string
	Status 			int8		 	//活动状态
	Gift			*Gift
	New				int
	CreatedAt		*time.Time
	IsTop			int8
	Number 			string
}

type GiftEnable struct {
	ID        	uint
	Name 		string 		`gorm:"column:name"`
	STATUS      int8   		`gorm:"column:status"`			//奖品状态，1=上架，2=下架，下架不可用
}

type GiftPage []Gift

func (Gift) TableName() string  {
	return "gift"
}

func (gift *Gift)Store(db *gorm.DB) (int64,error) {
	result := db.Create(gift)
	return result.RowsAffected,result.Error
}

func (gift *Gift)First(db *gorm.DB,id int64) (*enums.GiftDetail,bool,error) {
	detail := &enums.GiftDetail{}
	notFound := db.Table(gift.TableName()).
		Select("name,user_id,num,type,des,attachments").
		Where("id = ?",id).
		First(detail).
		RecordNotFound()
	return detail,notFound,db.Error
}

func (gift *Gift)Page(db *gorm.DB,page *PageParam) (GiftPage,error) {
	var gifts GiftPage
	err :=  Page(db,gift.TableName(),page).
		Select("id,user_id,name,num,type,from_type,status,des,attachments,created_at").
		Order("id desc").
		Find(&gifts).Error

	return gifts,err
}

func (gift *Gift) FindEnable(db *gorm.DB) ([]*GiftEnable,error) {
	var data []*GiftEnable
	db.Table(gift.TableName()).
		Where("status = ?",GIFT_STATUS_UP).
		Select("id,name,status").
		Order("id desc").
		Find(&data)
	return data,db.Error
}
