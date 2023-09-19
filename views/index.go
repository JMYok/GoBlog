package views

import (
	"GoBlog/common"
	"GoBlog/context"
	"GoBlog/service"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) IndexPageNew(ctx *context.MsContext) {
	indexTemplate := common.Template.Index

	//解析表单
	if err := ctx.Request.ParseForm(); err != nil {
		log.Println("Index获取数据出错：", err)
		indexTemplate.WriteError(ctx.W, errors.New("系统错误,请联系站长！"))
		return
	}

	//得到url中page参数
	pageStr := ctx.Request.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	//每页显示的数量
	pageSize := 10

	//取得URL slug 自定义路径
	path := ctx.Request.URL.Path
	slug := strings.TrimPrefix(path, "/")
	hr, err := service.GetAllIndexInfo(slug, page, pageSize)
	if err != nil {
		log.Println("index获取数据出错：", err)
		indexTemplate.WriteError(ctx.W, errors.New("系统错误,请联系站长！"))
	}
	indexTemplate.WriteData(ctx.W, hr)
}

// IndexPage 返回主页html逻辑
func (*HTMLApi) IndexPage(w http.ResponseWriter, r *http.Request) {
	indexTemplate := common.Template.Index

	//解析表单
	if err := r.ParseForm(); err != nil {
		log.Println("Index获取数据出错：", err)
		indexTemplate.WriteError(w, errors.New("系统错误,请联系站长！"))
		return
	}
	//得到url中page参数
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	//每页显示的数量
	pageSize := 10

	//取得URL slug 自定义路径
	path := r.URL.Path
	slug := strings.TrimPrefix(path, "/")
	hr, err := service.GetAllIndexInfo(slug, page, pageSize)
	if err != nil {
		log.Println("index获取数据出错：", err)
		indexTemplate.WriteError(w, errors.New("系统错误,请联系站长！"))
	}
	indexTemplate.WriteData(w, hr)
}
