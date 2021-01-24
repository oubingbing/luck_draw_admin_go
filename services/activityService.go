package services

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"luck-admin/enums"
	model "luck-admin/models"
	"time"
)

func SaveActivity(db *gorm.DB,param *enums.ActivityCreateParam) (int64,*enums.ErrorInfo) {
	attachments,encodeErr := json.Marshal(param.Attachments)
	if encodeErr != nil {
		return 0,&enums.ErrorInfo{Code:enums.ACTIVITY_IMAGE_ENCODE_ERR,Err:enums.ActivityEncodeImageErr}
	}
	shareImage,encodeErr := json.Marshal(param.ShareImage)
	if encodeErr != nil {
		return 0,&enums.ErrorInfo{Code:enums.ACTIVITY_IMAGE_ENCODE_ERR,Err:enums.ActivityEncodeImageErr}
	}

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
	}

	var parseErr error
	activity.StartAt,parseErr = time.Parse("2006-01-02 15:04:05",param.StartAt)
	if parseErr != nil {
		fmt.Println(parseErr.Error())
		return 0,&enums.ErrorInfo{enums.StartDateErr,enums.ACTIVITY_START_DATE_ERR}
	}

	activity.EndAt,parseErr = time.Parse("2006-01-02 15:04:05",param.EndAt)
	if parseErr != nil {
		fmt.Println(parseErr.Error())
		return 0,&enums.ErrorInfo{enums.EndDateErr,enums.ACTIVITY_END_DATE_ERR}
	}

	activity.RunAt,parseErr = time.Parse("2006-01-02 15:04:05",param.RunAt)
	if parseErr != nil {
		fmt.Println(parseErr.Error())
		return 0,&enums.ErrorInfo{enums.RunDateErr,enums.ACTIVITY_RUN_DATE_ERR}
	}

	_,err := FirstGiftById(db,activity.GiftId)
	if err != nil {
		return 0,err
	}

	effect,saveErr := activity.Store(db)
	return effect,&enums.ErrorInfo{saveErr,enums.ACTIVITY_SAVE_ERR}
}

func ActivityPage(db *gorm.DB,page *model.PageParam) (model.AcPage,*enums.ErrorInfo) {
	activity := &model.Activity{}
	activities,err := activity.Page(db,page)
	if err != nil {
		return nil,err
	}

	return activities,nil
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
