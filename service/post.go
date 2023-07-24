package service

import (
	"GoBlog/config"
	"GoBlog/dao"
	"GoBlog/models"
	models2 "GoBlog/models/template"
	"html/template"
	"log"
)

func GetPostDetail(pid int) (*models.PostRes, error) {
	post, err := dao.GetPostByPid(pid)

	if err != nil {
		return nil, err
	}

	categoryName, _ := dao.GetCategoryNameById(post.CategoryId)
	userName, _ := dao.GetUserNameById(post.UserId)

	var postMore = models.PostMore{
		Pid:          post.Pid,
		Title:        post.Title,
		Slug:         post.Slug,
		Content:      template.HTML(post.Content),
		CategoryId:   post.CategoryId,
		CategoryName: categoryName,
		UserId:       post.UserId,
		UserName:     userName,
		ViewCount:    post.ViewCount,
		Type:         post.Type,
		CreateAt:     models2.DateDay(post.CreateAt),
		UpdateAt:     models2.DateDay(post.UpdateAt),
	}

	var postRes = &models.PostRes{
		Viewer:       config.Cfg.Viewer,
		SystemConfig: config.Cfg.System,
		Article:      postMore,
	}

	return postRes, nil
}

func Writing() (wr models.WritingRes) {
	wr.Title = config.Cfg.Viewer.Title
	wr.CdnURL = config.Cfg.System.CdnURL
	categories, err := dao.GetAllCategory()
	if err != nil {
		log.Println(err)
		return
	}
	wr.Categories = categories
	return wr
}
