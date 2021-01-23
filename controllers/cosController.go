package controllers

import (
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/gin-gonic/gin"
	"luck-admin/enums"
	"luck-admin/services"
	"luck-admin/util"
)

func GetCosToken(ctx *gin.Context)  {
	token,err := services.CosToken()
	util.Info(token.SessionToken)
	if err != nil {
		util.ResponseJson(ctx,err.Code,err.Err.Error(),nil)
		return
	}

	domain,_ 	:= util.GetCosIni("cos_domain")
	bucket,_ 	:= util.GetCosIni("cos_bucket")
	region,_ 	:= util.GetCosIni("cos_region")
	mp := map[string]interface{}{
		"tmp_secret_id":token.TmpSecretId,
		"tmp_secret_key":token.TmpSecretKey,
		"token":token.SessionToken,
		"start_time":token.StartTime,
		"expired_ime":token.ExpiredTime,
		"domain":domain,
		"bucket":bucket,
		"region":region,
		"env":config.GetEnv(),
	}
	util.ResponseJson(ctx,enums.SUCCESS,"",mp)
	return
}
