package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type IndexData struct {
	Title       string
	Description string
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	var indexData IndexData
	indexData.Title = "go博客"
	indexData.Description = "go博客练手项目"
	t := template.New("index.html")
	path, _ := os.Getwd()

	prefixPath := path + "/template"

	header := prefixPath + "/layout/header.html"
	footer := prefixPath + "/layout/footer.html"
	pagination := prefixPath + "/layout/pagination.html"
	personal := prefixPath + "/layout/personal.html"
	postList := prefixPath + "/layout/post-list.html"
	home := prefixPath + "/header.html"
	index := prefixPath + "/index.html"

	t, _ = t.ParseFiles(header, footer, pagination, personal, postList, home, index)
	err := t.Execute(w, indexData)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", indexPage)
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
