package views

import (
	"GoBlog/common"
	"GoBlog/context"
	"GoBlog/service"
)

func (*HTMLApi) WritePage(ctx *context.MsContext) {
	writingTemplate := common.Template.Writing

	wr := service.Writing()
	writingTemplate.WriteData(ctx.W, wr)
}
