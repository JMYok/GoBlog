package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var indexData IndexData
	indexData.Title = "go博客"
	indexData.Desc = "go博客练手项目"
	msg, _ := json.Marshal(indexData)
	w.Write(msg)
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	var indexData IndexData
	indexData.Title = "go博客"
	indexData.Desc = "go博客练手项目"
	t := template.New("index.html")
	path, _ := os.Getwd()
	t, _ = t.ParseFiles(path + "/template/index.html")
	t.Execute(w, indexData)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	http.HandleFunc("/index.html", indexPage)
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
