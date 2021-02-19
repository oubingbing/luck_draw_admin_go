package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"luck-admin/enums"
	model "luck-admin/models"
	"luck-admin/util"
	"math/rand"
	"sort"
	"time"
)

func SaveActivity(db *gorm.DB,param *enums.ActivityCreateParam) (int64,*enums.ErrorInfo) {
	attachments,encodeErr := json.Marshal(param.Attachments)
	if encodeErr != nil {
		return 0,&enums.ErrorInfo{Code:enums.ACTIVITY_IMAGE_ENCODE_ERR,Err:enums.ActivityEncodeImageErr}
	}

	var shareImage []byte
	if len(param.ShareImage) > 0 {
		shareImage,encodeErr = json.Marshal(param.ShareImage)
		if encodeErr != nil {
			return 0,&enums.ErrorInfo{Code:enums.ACTIVITY_IMAGE_ENCODE_ERR,Err:enums.ActivityEncodeImageErr}
		}
	}

	countActivity := &model.Activity{}
	count,countErr := countActivity.CountToday(db)
	if countErr != nil {
		return 0,&enums.ErrorInfo{Code:enums.ACTIVITY_COUNT_TODAY_ERR,Err:enums.ActivityAcountTodayErr}
	}
	number := time.Now().Format(enums.DAY_FORMAT)+fmt.Sprintf("_%v",(count+1))

	activity := &model.Activity{
		Name:param.Name,
		GiftId:param.GiftId,
		Type:param.Type,
		OpenAd:param.OpenAd,
		FromType:model.ACTIVITY_FROM_USER,
		LimitJoin:param.LimitJoin,
		JoinLimitNum:param.JoinLimitNum,
		ReceiveLimit:param.ReceiveLimit,
		Des:param.Des,
		Attachments:string(attachments),
		ShareTitle:param.ShareTitle,
		ShareImage:string(shareImage),
		Status:model.ACTIVITY_STATSUS_TO_RELE,
		Really:param.Really,
		DrawType:param.DrawType,
		BigPic:param.BigPic,
		Number:number,
		IsTop:param.IsTop,
	}

	_,err := FirstGiftById(db,activity.GiftId)
	if err != nil {
		return 0,err
	}

	//判断是否真送，假送就生成假送数据
	var fakerUserIndex []int
	redisClient := util.Redis()
	defer redisClient.Close()

	var fakerNUm float32
	if param.DrawType == model.ACTIVITY_DRAW_TYPE_RAND_all {
		fakerNUm = activity.JoinLimitNum * 0.16
	}else if param.DrawType == model.ACTIVITY_DRAW_TYPE_TIME {
		fakerNUm = 18
	}else{
		fakerNUm = activity.ReceiveLimit
	}

	if param.Really == model.ACTIVITY_REALLY_N {
		for {
			if fakerNUm <= 0 {
				break
			}
			//生成假用户数据
			rand.Seed(time.Now().UnixNano()+int64(fakerNUm))
			key := rand.Intn(int(activity.JoinLimitNum))
			if key == 0 {
				key = 1
			}

			//保证假人能排入队伍
			if key > 0 && key <= int(activity.JoinLimitNum) - 2 {
				fakerUserIndex = append(fakerUserIndex, key)
				fakerNUm --
			}
		}
	}

	effect,saveErr := activity.Store(db)

	ctx := context.Background()
	if len(fakerUserIndex)  > 0{
		cacheKey := fmt.Sprintf("%v:%v",model.FAKER_USER_KEY,activity.ID)
		sort.Ints(fakerUserIndex)
		fakerCacheStr,_ := json.Marshal(&fakerUserIndex)
		redisClient.Set(ctx,cacheKey,fakerCacheStr,0)
	}

	if param.IsTop == model.ACTIVITY_IS_TOP_Y {
		refreshErr := RefreshTopActivity(db)
		if refreshErr != nil {
			return effect,refreshErr
		}
	}

	return effect,&enums.ErrorInfo{saveErr,enums.ACTIVITY_SAVE_ERR}
}

func RefreshTopActivity(db *gorm.DB) *enums.ErrorInfo {
	redisClient := util.Redis()
	defer redisClient.Close()

	topActivity := &model.Activity{}
	tops,topErr := topActivity.Tops(db)
	if topErr != nil {
		return &enums.ErrorInfo{enums.ActivityQueryTopErr,enums.ACTIVITY_QUERY_TOP_ERR}
	}
	topsByte,encodeErr := json.Marshal(&tops)
	if encodeErr != nil {
		return &enums.ErrorInfo{enums.DecodeErr,enums.DECODE_ARR_ERR}
	}
	redisClient.Set(ctx,model.TOP_ACTIVITY,string(topsByte),0)

	return nil
}

func ActivityPage(db *gorm.DB,page *model.PageParam) (model.AcPage,*enums.ErrorInfo) {
	activity := &model.Activity{}
	activities,err := activity.Page(db,page)
	if err != nil {
		return nil,err
	}

	domain,_ 	:= util.GetCosIni("cos_domain")
	for index,_:= range activities {
		 ActivityFormat(domain,&activities[index])
	}

	return activities,nil
}

func ActivityFormat(domain string,activity *model.ActivityPageFormat)  {
	activity.AttachmentsStr,_ = AppendDomain(domain,activity.Attachments)
	activity.ShareImageStr,_  = AppendDomain(domain,activity.ShareImage)

	switch int(activity.Status) {
		case model.ACTIVITY_STATSUS_TO_RELE:
			activity.StatusStr = "待上架"
			break
		case model.ACTIVITY_STATSUS_RUNNING:
			activity.StatusStr = "进行中"
			break
		case model.ACTIVITY_STATSUS_DOWN:
			activity.StatusStr = "已下架"
			break
		case model.ACTIVITY_STATSUS_FINISH:
			activity.StatusStr = "已结束"
			break
		default:
			activity.StatusStr = "未知状态"
	}

	switch int(activity.Type) {
		case model.ACTIVITY_TYPE_RED_PAK:
			activity.TypeStr = "红包"
			break
		case model.ACTIVITY_TYPE_GOODS:
			activity.TypeStr = "礼品"
			break
		case model.ACTIVITY_TYPE_PHONE_BILL:
			activity.TypeStr = "话费"
			break
		default:
			activity.TypeStr = "未知"
	}

	switch int(activity.DrawType) {
	case model.ACTIVITY_DRAW_TYPE_AVERAGE:
		activity.DrawTypeStr = "平均"
		break
	case model.ACTIVITY_DRAW_TYPE_RAND:
		activity.DrawTypeStr = "拼手气"
		break
	case model.ACTIVITY_DRAW_TYPE_RAND_all:
		activity.DrawTypeStr = "拼手气人人有份"
		break
	default:
		activity.TypeStr = "未知"
	}

	if int(activity.OpenAd) == model.ACTIVITY_OP_AD_Y {
		activity.OpenAdStr = "已开启广告"
	}else{
		activity.OpenAdStr = "广告关闭"
	}

	if int(activity.Really) == model.ACTIVITY_REALLY_N {
		activity.ReallyStr = "假送"
	}else{
		activity.ReallyStr = "真送"
	}

	if int(activity.BigPic) == model.ACTIVITY_BIG_PIC_BIG {
		activity.BigPicStr = "大图"
	}else{
		activity.BigPicStr = "小图"
	}
}

func ActivityDetail(db *gorm.DB,id string) (*enums.ActivityDetailFormat,*enums.ErrorInfo) {
	activity := &model.Activity{}
	detail,acNotFound,err := activity.Detail(db,id)
	if err != nil {
		return nil,&enums.ErrorInfo{err,enums.ACTIVITY_DETAIL_QUERY_ERR}
	}

	if acNotFound {
		return nil,&enums.ErrorInfo{enums.ActivityDetailNotFound,enums.ACTIVITY_DETAIL_NOT_FOUND}
	}

	gift := &model.Gift{}
	giftDetail,notFound,err := gift.First(db,detail.GiftId)
	if err != nil {
		return nil,&enums.ErrorInfo{err,enums.GIFT_GET_DETAIL_ERR}
	}
	if notFound {
		return nil,&enums.ErrorInfo{enums.GiftNotFound,enums.GIFT_NOT_FOUND}
	}
	detail.Gift = giftDetail

	return detail,nil
}

func ActivityDelete(db *gorm.DB,id uint) *enums.ErrorInfo {
	activity := &model.Activity{}
	activity.ID = id
	err := activity.Delete(db)
	if err != nil {
		return &enums.ErrorInfo{enums.ActivityDeleteErr,enums.ACTIVITY_DELETE_ERR}
	}

	joinLog := &model.JoinLog{}
	err = joinLog.DeleteByAid(db,id)
	if err != nil {
		return &enums.ErrorInfo{enums.ActivityDeleteErr,enums.ACTIVITY_DELETE_ERR}
	}

	refreshErr := RefreshTopActivity(db)
	if refreshErr != nil {
		return refreshErr
	}

	return nil
}

func ActivityUpdateStatus(db *gorm.DB,id interface{},status interface{}) *enums.ErrorInfo {
	var err error
	activity := &model.Activity{}

	if status.(float64) == float64(model.ACTIVITY_STATSUS_RUNNING) {
		//上架
		err = activity.Up(db,id)
	}else if status.(float64) == float64(model.ACTIVITY_STATSUS_DOWN) {
		//下架
		err = activity.Down(db,id)
	}else{
		return &enums.ErrorInfo{enums.ActivityUpdateStatusErr,enums.ACTIVITY_UPDATE_STATUS_ERR}
	}

	if err != nil {
		return &enums.ErrorInfo{enums.ActivityUpdateStatusErr,enums.ACTIVITY_UPDATE_STATUS_BAD_ERR}
	}

	err = activity.FindById(db,id)
	if activity.IsTop == model.ACTIVITY_IS_TOP_Y {
		refreshErr := RefreshTopActivity(db)
		if refreshErr != nil {
			return refreshErr
		}
	}

	return nil
}

func StrToArr(str string) ([]string,*enums.ErrorInfo) {
	var sli []string
	err := json.Unmarshal([]byte(str),&sli)
	if err != nil {
		return nil,&enums.ErrorInfo{enums.DecodeErr,enums.DECODE_ARR_ERR}
	}

	return sli,nil
}

func AppendDomain(domain,str string) ([]string,*enums.ErrorInfo) {
	sli,err := StrToArr(str)
	if err != nil {
		return nil,err
	}

	for index,_ := range sli {
		sli[index] = domain+"/"+sli[index]
	}

	return sli,nil
}

func ActivityUpdate(db *gorm.DB,param *enums.ActivityUpdateParam) (*model.ActivityPageFormat,*enums.ErrorInfo) {
	activity := &model.Activity{}
	err := activity.FindById(db,param.Id)
	if err != nil {
		return nil,&enums.ErrorInfo{enums.ActivityDetailNotFound,enums.ACTIVITY_NOT_FOUND}
	}

	_,findGiftErr := FirstGiftById(db,activity.GiftId)
	if findGiftErr != nil {
		return nil,&enums.ErrorInfo{Code:enums.GIFT_GET_DETAIL_ERR,Err:enums.GiftNotFound}
	}

	if activity.IsTop == model.ACTIVITY_IS_TOP_Y {
		refreshErr := RefreshTopActivity(db)
		if refreshErr != nil {
			return nil,refreshErr
		}
	}

	/*attachments,encodeErr := json.Marshal(param.Attachments)
	if encodeErr != nil {
		return nil,&enums.ErrorInfo{Code:enums.ACTIVITY_IMAGE_ENCODE_ERR,Err:enums.ActivityEncodeImageErr}
	}*/

	/*var shareImage []byte
	if len(param.ShareImage) > 0 {
		shareImage,encodeErr = json.Marshal(param.ShareImage)
		if encodeErr != nil {
			return nil,&enums.ErrorInfo{Code:enums.ACTIVITY_IMAGE_ENCODE_ERR,Err:enums.ActivityEncodeImageErr}
		}
	}*/

	activity.Name = param.Name
	activity.GiftId = param.GiftId
	activity.Type = param.Type
	activity.OpenAd = param.OpenAd
	activity.FromType = model.ACTIVITY_FROM_USER
	activity.LimitJoin = param.LimitJoin
	activity.JoinLimitNum = param.JoinLimitNum
	activity.ReceiveLimit = param.ReceiveLimit
	activity.Des = param.Des
	activity.Attachments = param.Attachments
	activity.ShareTitle = param.ShareTitle
	activity.ShareImage = param.ShareImage
	activity.Status = model.ACTIVITY_STATSUS_TO_RELE
	activity.Really = param.Really
	activity.DrawType = param.DrawType
	activity.BigPic = param.BigPic
	saveErr := db.Save(activity).Error
	if saveErr != nil {
		return nil,&enums.ErrorInfo{Code:enums.ACTIVITY_UPDATE_ERR,Err:enums.ActivityUpdateStatusErr}
	}

	activityFormat := &model.ActivityPageFormat{}
	activityFormat.ID = activity.ID
	activityFormat.Name = activity.Name
	activityFormat.GiftId = activity.GiftId
	activityFormat.Type = activity.Type
	activityFormat.OpenAd = activity.OpenAd
	activityFormat.FromType = int32(activity.FromType)
	activityFormat.LimitJoin = activity.LimitJoin
	activityFormat.JoinLimitNum = activity.JoinLimitNum
	activityFormat.ReceiveLimit = activity.ReceiveLimit
	activityFormat.Des = activity.Des
	activityFormat.Attachments = activity.Attachments
	activityFormat.ShareTitle = activity.ShareTitle
	activityFormat.ShareImage = activity.ShareImage
	activityFormat.Status = activity.Status
	activityFormat.Really = activity.Really
	activityFormat.DrawType = activity.DrawType
	activityFormat.BigPic = activity.BigPic

	return activityFormat,nil
}