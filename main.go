package main

import (
	"GoBlog/common"
	"GoBlog/server"
)

// TODO: 1.错误日志优化

func init() {
	//模板加载
	common.LoadTemplate()
}

func main() {
	server.APP.Start()
}
