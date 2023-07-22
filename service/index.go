package service

import (
	"GoBlog/config"
	"GoBlog/dao"
	"GoBlog/models"
	"log"
)

func GetAllIndexInfo() (*models.HomeResponse, error) {
	categories, err := dao.GetAllCategory()
	if err != nil {
		log.Println(err)
		return nil, err
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

	return hr, nil
}
