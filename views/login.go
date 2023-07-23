package views

import (
	"GoBlog/common"
	"GoBlog/config"
	"net/http"
)

func (*HTMLApi) LoginPage(w http.ResponseWriter, r *http.Request) {
	loginTemplate := common.Template.Login

	loginTemplate.WriteData(w, config.Cfg.Viewer)
}
