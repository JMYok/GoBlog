package service

import (
	"GoBlog/config"
	"GoBlog/dao"
	"GoBlog/models"
	models2 "GoBlog/models/template"
	"html/template"
	"log"
)

// GetAllIndexInfo 获取主页所有数据
func GetAllIndexInfo(page, pageSize int) (*models.HomeResponse, error) {
	categories, err := dao.GetAllCategory()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	posts, err := dao.GetPostInfo(page, pageSize)
	var postMores []models.PostMore
	for _, post := range posts {
		categoryName, _ := dao.GetCategoryNameById(post.CategoryId)
		userName, _ := dao.GetUserNameById(post.UserId)

		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[:100]
		}

		postMore := models.PostMore{
			Pid:          post.Pid,
			Title:        post.Title,
			Slug:         post.Slug,
			Content:      template.HTML(content),
			CategoryId:   post.CategoryId,
			CategoryName: categoryName,
			UserId:       post.UserId,
			UserName:     userName,
			ViewCount:    post.ViewCount,
			Type:         post.Type,
			CreateAt:     models2.DateDay(post.CreateAt),
			UpdateAt:     models2.DateDay(post.UpdateAt),
		}
		postMores = append(postMores, postMore)
	}

	if err != nil {
		log.Println(err)
		return nil, err
	}

	var hr = &models.HomeResponse{
		Viewer:     config.Cfg.Viewer,
		Categories: categories,
		Posts:      postMores}

	return hr, nil
}
