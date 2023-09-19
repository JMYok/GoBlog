package views

import (
	"GoBlog/common"
	"GoBlog/context"
	"GoBlog/service"
)

func (*HTMLApi) PigeonholePage(ctx *context.MsContext) {
	pigeonholeTemplate := common.Template.Pigeonhole
	pigeonholeRes := service.FindPostPigeonhole()
	pigeonholeTemplate.WriteData(ctx.W, pigeonholeRes)
}
