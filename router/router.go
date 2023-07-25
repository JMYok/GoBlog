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
	http.HandleFunc("/p/", views.HTML.DetailPage)
	http.HandleFunc("/login", views.HTML.LoginPage)
	http.HandleFunc("/writing", views.HTML.WritePage)

	//2. api
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/post/", api.API.GetPost)
	http.HandleFunc("/api/v1/login", api.API.Login)

	//3、静态资源映射
	//TODO index.html 中 js路径为/resource 为什么默认public/resource
	http.Handle("/public/resource/", http.StripPrefix("/public/resource/", http.FileServer(http.Dir("public/resource/"))))
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
	http.Handle("/c/resource/", http.StripPrefix("/c/resource/", http.FileServer(http.Dir("public/resource/"))))
	http.Handle("/p/resource/", http.StripPrefix("/p/resource/", http.FileServer(http.Dir("public/resource/"))))
	http.Handle("/p/img/", http.StripPrefix("/p/img/", http.FileServer(http.Dir("public/resource/"))))
}
