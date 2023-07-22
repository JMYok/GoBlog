package views

import (
	"GoBlog/common"
	"GoBlog/service"
	"errors"
	"log"
	"net/http"
	"strconv"
)

// IndexPage 返回主页html逻辑
func (*HTMLApi) IndexPage(w http.ResponseWriter, r *http.Request) {
	indexTemplate := common.Template.Index

	//解析表单
	if err := r.ParseForm(); err != nil {
		log.Println("Index获取数据出错：", err)
		indexTemplate.WriteError(w, errors.New("系统错误,请联系站长！"))
		return
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	//每页显示的数量
	pageSize := 5
	hr, err := service.GetAllIndexInfo(page, pageSize)
	if err != nil {
		log.Println("index获取数据出错：", err)
		indexTemplate.WriteError(w, errors.New("系统错误,请联系站长！"))
	}
	indexTemplate.WriteData(w, hr)
}
