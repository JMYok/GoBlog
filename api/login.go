package api

import (
	"GoBlog/common"
	"GoBlog/context"
	"GoBlog/service"
	"GoBlog/utils"
)

func (*APIHandler) Login(ctx *context.MsContext) {
	params := common.GetRequestJsonParam(ctx.Request)
	username := params["username"].(string)
	passwd := utils.Md5Crypt(params["passwd"].(string), "jmycool")
	loginRes, err := service.Login(username, passwd)
	if err != nil {
		common.Error(ctx.W, err)
	}
	common.Success(ctx.W, *loginRes)
}
