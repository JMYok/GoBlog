package views

import (
	"GoBlog/common"
	"GoBlog/service"
	"net/http"
)

func (*HTMLApi) PigeonholePage(w http.ResponseWriter, r *http.Request) {
	pigeonholeTemplate := common.Template.Pigeonhole
	pigeonholeRes := service.FindPostPigeonhole()
	pigeonholeTemplate.WriteData(w, pigeonholeRes)
}
