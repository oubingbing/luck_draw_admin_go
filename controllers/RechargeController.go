package controllers

import (
	"github.com/gin-gonic/gin"
	"luck-admin/services"
)

func Recharge(ctx *gin.Context)  {
	services.RechargeBill()
}
