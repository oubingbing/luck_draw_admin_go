package controllers

import (
	"github.com/gin-gonic/gin"
	"luck-admin/enums"
	"luck-admin/services"
	"luck-admin/util"
)

func GetCosToken(ctx *gin.Context)  {
	token,err := services.CosToken()
	if err != nil {
		util.ResponseJson(ctx,err.Code,err.Err.Error(),nil)
		return
	}

	domain,_ 	:= util.GetCosIni("cos_domain")
	mp := map[string]string{"token":token,"domain":domain}
	util.ResponseJson(ctx,enums.SUCCESS,"",mp)
	return
}
