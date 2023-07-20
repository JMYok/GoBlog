package views

import (
	"GoBlog/common"
	"GoBlog/config"
	"GoBlog/models"
	"net/http"
)

// IndexPage 返回主页html逻辑
func (*HTMLApi) IndexPage(w http.ResponseWriter, r *http.Request) {
	//页面上涉及到的所有数据
	var categories = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}

	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go blog",
			Content:      "哈哈哈哈哈哈",
			UserName:     "bob",
			ViewCount:    123,
			CreateAt:     "2023-7-19",
			CategoryName: "Go",
			Type:         0,
		},
	}

	var hr = &models.HomeResponse{
		Viewer:     config.Cfg.Viewer,
		Categories: categories,
		Posts:      posts}

	indexTemplate := common.Template.Index
	indexTemplate.WriteData(w, hr)
}
