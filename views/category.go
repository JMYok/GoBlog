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

func (*HTMLApi) CategoryNew(ctx *context.MsContext) {
	categoryTemplate := common.Template.Category
	cidStr := ctx.GetPathVariable("id")
	cid, _ := strconv.Atoi(cidStr)
	pageStr, _ := ctx.GetForm("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)
	//每页显示的数量
	pageSize := 10
	categoryResponse, err := service.GetCategoryById(cid, page, pageSize)
	if err != nil {
		categoryTemplate.WriteError(ctx.W, err)
		return
	}
	categoryTemplate.WriteData(ctx.W, categoryResponse)
}

func (*HTMLApi) CategoryPage(w http.ResponseWriter, r *http.Request) {
	categoryTemplate := common.Template.Category

	pathStr := r.URL.Path
	cidStr, found := strings.CutPrefix(pathStr, "/c/")
	cid, _ := strconv.Atoi(cidStr)
	if found == false {
		categoryTemplate.WriteError(w, errors.New("未找到分类编号！"))
		return
	}

	//得到url中page参数
	pageStr, _ := strings.CutPrefix(r.URL.RawQuery, "page=")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	//每页显示的数量
	pageSize := 10

	//根据cid分页查询到该分类下的博客
	categoryResponse, err := service.GetCategoryById(cid, page, pageSize)

	if err != nil {
		log.Println(err)
		panic("分页查询错误")
	}

	//通过template，将数据写入到页面上
	categoryTemplate.WriteData(w, categoryResponse)
}
