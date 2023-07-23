package models

import "GoBlog/config"

type HomeResponse struct {
	config.Viewer
	Categories []Category
	Posts      []PostMore
	Total      int   //总数据数
	Page       int   //当前页数
	Pages      []int //总页数
	PageEnd    bool  //是否到最后一页
}
