package controllers

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/template/types"
)

func GetDashBoard(ctx *context.Context) (types.Panel, error) {
	//user := auth.Auth(ctx)




	return types.Panel{
		Content:     "",
		Title:       "Dashboard",
		Description: "dashboard example",
	}, nil
}