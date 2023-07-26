package server

import (
	"GoBlog/config"
	"GoBlog/router"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"net/http"
)

var APP = &MyServer{}

type MyServer struct {
}

func (*MyServer) Start() {
	var serverConfig config.ServerTomlConfig
	_, err := toml.DecodeFile("config/server_config.toml", &serverConfig)

	if err != nil {
		panic(err)
	}

	server := http.Server{
		Addr: serverConfig.Server.Ip + ":" + serverConfig.Server.Port,
	}

	log.Printf("server running on %s:%s\n", serverConfig.Server.Ip, serverConfig.Server.Port)

	//路由
	router.Router()

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
