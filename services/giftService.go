package services

import (
	"github.com/jinzhu/gorm"
	"luck-admin/enums"
	"luck-admin/models"
)

/**
 * 新建礼品
 */
func SaveGift(db *gorm.DB,userId int,giftParam *enums.GiftParam) (int64,*enums.ErrorInfo) {
	gift := &models.Gift{
		Name:giftParam.Name,
		Num:giftParam.Num,
		UserId:userId,
		Type:giftParam.Type,
		FROM:giftParam.FROM,
		STATUS:giftParam.STATUS,
		Des:giftParam.Des,
		Attachments:giftParam.Attachments,
	}

	effect,err := gift.Store(db)
	if err != nil {
		return effect,&enums.ErrorInfo{enums.GiftSaveErr,enums.GIFT_SAVE_ERR}
	}

	return effect,nil
}
