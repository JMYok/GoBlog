package views

import (
	"GoBlog/common"
	"GoBlog/service"
	"errors"
	"log"
	"net/http"
)

// IndexPage 返回主页html逻辑
func (*HTMLApi) IndexPage(w http.ResponseWriter, r *http.Request) {
	indexTemplate := common.Template.Index
	hr, err := service.GetAllIndexInfo()
	if err != nil {
		log.Println("index获取数据出错：", err)
		indexTemplate.WriteError(w, errors.New("系统错误,请联系站长！"))
	}
	indexTemplate.WriteData(w, hr)
}
