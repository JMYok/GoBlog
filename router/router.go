package router

import (
	"GoBlog/api"
	"GoBlog/context"
	"GoBlog/views"
	"net/http"
)

// Router 返回 1.页面 views  2.api 数据(json) 3.静态资源
func Router() {
	// 所有请求都交给实现 ServeHTTP 的 Context
	http.Handle("/", context.Context)

	//注册路由处理函数，插入前缀树
	// 1.页面
	context.Context.Handler("/", views.HTML.IndexPageNew)
	context.Context.Handler("/c/{id}", views.HTML.CategoryNew)
	context.Context.Handler("/p/{id}", views.HTML.PostDetail)
	context.Context.Handler("/login", views.HTML.LoginPageNew)
	context.Context.Handler("/writing", views.HTML.WritePage)
	context.Context.Handler("/pigeonhole", views.HTML.PigeonholePage)

	// 2. api
	context.Context.Handler("/api/v1/login", api.API.Login)

	// 文章管理
	context.Context.Handler("/api/v1/post", api.API.SaveAndUpdatePost)
	context.Context.Handler("/api/v1/post/{pid}", api.API.GetPost)
	context.Context.Handler("/api/v1/post/search", api.API.SearchPost)

	//3、静态资源映射
	http.Handle("/public/resource/", http.StripPrefix("/public/resource/", http.FileServer(http.Dir("public/resource/"))))
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
	http.Handle("/c/resource/", http.StripPrefix("/c/resource/", http.FileServer(http.Dir("public/resource/"))))
	http.Handle("/p/resource/", http.StripPrefix("/p/resource/", http.FileServer(http.Dir("public/resource/"))))
	http.Handle("/p/img/", http.StripPrefix("/p/img/", http.FileServer(http.Dir("public/resource/"))))
}
