package service

import (
	"GoBlog/config"
	"GoBlog/dao"
	"GoBlog/models"
)

func FindPostPigeonhole() models.PigeonholeRes {
	//查询所有的文章 进行月份的整理
	//查询所有的分类
	posts, _ := dao.GetPostAll()
	pigeonholeMap := make(map[string][]models.Post)
	for _, post := range posts {
		at := post.CreateAt
		month := at.Format("2006-01")
		pigeonholeMap[month] = append(pigeonholeMap[month], post)
	}
	categories, _ := dao.GetAllCategory()
	return models.PigeonholeRes{
		Viewer:       config.Cfg.Viewer,
		SystemConfig: config.Cfg.System,
		Categories:   categories,
		Lines:        pigeonholeMap,
	}
}
