package main

import (
	"GoBlog/config"
	"GoBlog/models"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type IndexData struct {
	Title       string
	Description string
}

func getNextName(navigation []string, index int) string {
	return navigation[index+1]
}

func isODD(num int) bool {
	return num%2 == 0
}

func date(layout string) string {
	return time.Now().Format(layout)
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	var indexData IndexData
	indexData.Title = "go博客"
	indexData.Description = "go博客练手项目"
	t := template.New("index.html")
	//拿到当前路径
	path := config.Cfg.System.CurrentDir

	prefixPath := path + "/template"

	header := prefixPath + "/layout/header.html"
	footer := prefixPath + "/layout/footer.html"
	pagination := prefixPath + "/layout/pagination.html"
	personal := prefixPath + "/layout/personal.html"
	postList := prefixPath + "/layout/post-list.html"
	home := prefixPath + "/home.html"
	index := prefixPath + "/index.html"

	//映射页面上的方法
	t.Funcs(template.FuncMap{"isODD": isODD, "getNextName": getNextName, "date": date})

	t, err := t.ParseFiles(header, footer, pagination, personal, postList, home, index)
	if err != nil {
		log.Println("解析模板出错：", err)
	}

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

	err = t.Execute(w, hr)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", indexPage)

	//静态资源映射
	//TODO index.html 中 js路径为/resource 为什么默认public resource
	http.Handle("/public/resource/", http.StripPrefix("/public/resource/", http.FileServer(http.Dir("public/resource/"))))
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
