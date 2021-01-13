package services

import (
	"github.com/jinzhu/gorm"
	"luck-admin/enums"
	"luck-admin/models"
	"luck-admin/util"
)

func UserPage(db *gorm.DB,page *models.PageParam) (*models.UserPage,*enums.ErrorInfo) {
	user := &models.User{}
	userList,err := user.Page(db,page)
	if err != nil {
		util.ErrDetail(enums.USER_PAGE_QUERY_ERR,enums.UserPageQueryErr.Error(),nil)
		return nil,&enums.ErrorInfo{Code:enums.USER_PAGE_QUERY_ERR,Err:enums.UserPageQueryErr}
	}

	return userList,nil
}

func FormatUserItem(item *enums.UserPage)  {
	if item.FromType == models.USER_FROM_MINI {
		item.FromTypeStr = "微信小程序"
	}else{
		item.FromTypeStr = "微信公众号"
	}

	switch item.Gender {
		case models.USER_GENDER_BOY:item.GenderStr = "男"
		case models.USER_GENDER_GIRL:item.GenderStr = "女"
		case models.USER_GENDER_OTHER:item.GenderStr = "其他"
		default:item.GenderStr = "未知"
	}

	if item.Status == models.USER_STATUS_WHITE {
		item.StatusStr = "正常"
	}else{
		item.StatusStr = "已被拉黑"
	}
}
