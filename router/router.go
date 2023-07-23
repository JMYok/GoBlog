package router

import (
	"GoBlog/api"
	"GoBlog/views"
	"net/http"
)

// Router 返回 1.页面 views  2.api 数据(json) 3.静态资源
func Router() {
	//1. 页面
	http.HandleFunc("/", views.HTML.IndexPage)
	http.HandleFunc("/c/", views.HTML.CategoryPage)

	//2. api
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)

	//3、静态资源映射
	//TODO index.html 中 js路径为/resource 为什么默认public/resource
	http.Handle("/public/resource/", http.StripPrefix("/public/resource/", http.FileServer(http.Dir("public/resource/"))))
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
