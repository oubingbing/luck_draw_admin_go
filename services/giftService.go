package services

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"luck-admin/enums"
	"luck-admin/models"
	"luck-admin/util"
)

/**
 * 新建礼品
 */
func SaveGift(db *gorm.DB,userId int,giftParam *enums.GiftParam) (*models.Gift,*enums.ErrorInfo) {
	attachmentsStr,encodeErr := json.Marshal(giftParam.Attachments)
	if encodeErr != nil {
		return nil,&enums.ErrorInfo{enums.GiftAttachmentsEncodeErr,enums.GIFT_ATTACHMENTS_ENCODE_ERR}
	}

	gift := &models.Gift{
		Name:giftParam.Name,
		Num:giftParam.Num,
		UserId:userId,
		Type:giftParam.Type,
		FROM:giftParam.FROM,
		STATUS:giftParam.STATUS,
		Des:giftParam.Des,
		Attachments:string(attachmentsStr),
	}

	effect,err := gift.Store(db)
	if err != nil {
		return nil,&enums.ErrorInfo{enums.GiftSaveErr,enums.GIFT_SAVE_ERR}
	}

	if effect <= 0 {
		return nil,&enums.ErrorInfo{enums.GiftSaveErr,enums.GIFT_SAVE_ERR}
	}

	return gift,nil
}

func FormatGift(gift *models.Gift,cosDomain string) (*enums.GiftResponse,*enums.ErrorInfo) {
	var attachments []string
	err := json.Unmarshal([]byte(gift.Attachments),&attachments)
	if err != nil {
		util.ErrDetail(enums.GIFT_ATTACHMENTS_DECODE_ERR,err.Error(),nil)
		return nil,&enums.ErrorInfo{enums.GiftAttachmentsDecodeErr,enums.GIFT_ATTACHMENTS_DECODE_ERR}
	}

	for index,item := range attachments {
		attachments[index] = cosDomain + "/" + item
	}

	giftRes := &enums.GiftResponse{
		ID:          gift.ID,
		Name:        gift.Name,
		UserId:      gift.UserId,
		Num:         gift.Num,
		Type:        gift.Type,
		Des:         gift.Des,
		Attachments: attachments,
		CreatedAt: gift.CreatedAt,
	}

	return giftRes,nil
}

func PageGift(db *gorm.DB,page *models.PageParam) ([]*enums.GiftResponse,*enums.ErrorInfo) {
	gift := &models.Gift{}
	result,err := gift.Page(db,page)
	if err != nil {
		return nil,&enums.ErrorInfo{Code:enums.GIFT_PAGE_QUERY_FAIL,Err:enums.GiftPageQueryFail}
	}

	domain,_ 	:= util.GetCosIni("cos_domain")
	var giftRes []*enums.GiftResponse
	for _,item := range result {
		formatGift,formatErr := FormatGift(&item,domain)
		if formatErr != nil {
			return nil,formatErr
		}
		giftRes = append(giftRes,formatGift)
	}

	return giftRes,nil
}

func FirstGiftById(db *gorm.DB,id int64) (*enums.GiftDetail,*enums.ErrorInfo) {
	gift := &models.Gift{}
	detail,notFound,err := gift.First(db,id)
	if err != nil {
		return nil,&enums.ErrorInfo{err,enums.GIFT_FIRST_ERR}
	}

	if notFound {
		return nil,&enums.ErrorInfo{enums.GiftNotFound,enums.GIFT_NOT_FOUND}
	}

	return detail,nil
}

func FindGiftEnable(db *gorm.DB) ([]*models.GiftEnable,*enums.ErrorInfo) {
	gift := &models.Gift{}
	data,err := gift.FindEnable(db)
	if err != nil {
		return nil,&enums.ErrorInfo{Code:enums.GIFT_FIND_ENABLE_ERR,Err:enums.GiftFindEndableErr}
	}

	return data,nil
}
