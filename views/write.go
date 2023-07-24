package views

import (
	"GoBlog/common"
	"GoBlog/service"
	"net/http"
)

func (*HTMLApi) WritePage(w http.ResponseWriter, r *http.Request) {
	writingTemplate := common.Template.Writing
	wr := service.Writing()
	writingTemplate.WriteData(w, wr)
}
