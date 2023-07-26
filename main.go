package main

import (
	"GoBlog/common"
	"GoBlog/server"
)

func init() {
	//模板加载
	common.LoadTemplate()
}

func main() {
	server.APP.Start()
}
