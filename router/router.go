package router

import (
	"GoBlog/context"
	"GoBlog/views"
	"net/http"
)

// Router 返回 1.页面 views  2.api 数据(json) 3.静态资源
func Router() {
	http.Handle("/", context.Context)

	context.Context.Handler("/c/{id}", views.HTML.CategoryNew)

	/*
		//1. 页面
		http.HandleFunc("/", views.HTML.IndexPage)
		http.HandleFunc("/c/", views.HTML.CategoryPage)
		http.HandleFunc("/p/", views.HTML.DetailPage)
		http.HandleFunc("/login", views.HTML.LoginPage)
		http.HandleFunc("/writing", views.HTML.WritePage)
		http.HandleFunc("/pigeonhole", views.HTML.PigeonholePage)

		//2. api
		http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
		http.HandleFunc("/api/v1/post/", api.API.GetPost)
		http.HandleFunc("/api/v1/post/search", api.API.SearchPost)
		http.HandleFunc("/api/v1/login", api.API.Login)
		//http.HandleFunc("/api/v1/qiniu/token", api.API.QiniuToken)
	*/

	//3、静态资源映射
	//TODO index.html 中 js路径为/resource 为什么默认public/resource
	http.Handle("/public/resource/", http.StripPrefix("/public/resource/", http.FileServer(http.Dir("public/resource/"))))
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
	http.Handle("/c/resource/", http.StripPrefix("/c/resource/", http.FileServer(http.Dir("public/resource/"))))
	http.Handle("/p/resource/", http.StripPrefix("/p/resource/", http.FileServer(http.Dir("public/resource/"))))
	http.Handle("/p/img/", http.StripPrefix("/p/img/", http.FileServer(http.Dir("public/resource/"))))
}
