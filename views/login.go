package views

import (
	"GoBlog/common"
	"GoBlog/config"
	"GoBlog/context"
	"net/http"
)

func (*HTMLApi) LoginPageNew(ctx *context.MsContext) {
	loginTemplate := common.Template.Login

	loginTemplate.WriteData(ctx.W, config.Cfg.Viewer)
}

func (*HTMLApi) LoginPage(w http.ResponseWriter, r *http.Request) {
	loginTemplate := common.Template.Login

	loginTemplate.WriteData(w, config.Cfg.Viewer)
}
