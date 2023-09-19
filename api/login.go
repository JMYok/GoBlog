package api

import (
	"GoBlog/common"
	"GoBlog/context"
	"GoBlog/service"
)

func (*APIHandler) Login(ctx *context.MsContext) {
	userName := ctx.GetJson("username").(string)
	passwd := ctx.GetJson("passwd").(string)
	loginRes, err := service.Login(userName, passwd)
	if err != nil {
		common.Error(ctx.W, err)
	}
	common.Success(ctx.W, *loginRes)
}
