package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

type tomlConfig struct {
	Viewer Viewer
	System SystemConfig
}

type ServerTomlConfig struct {
	Server ServerConfig
}

type MysqlTomlConfig struct {
	Mysql MysqlConfig
}

type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Avatar      string
	UserName    string
	UserDesc    string
}
type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string
	QiniuAccessKey  string
	QiniuSecretKey  string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURL string
}

type MysqlConfig struct {
	UserName     string
	Password     string
	Host         string
	Port         string
	DatabaseName string
}

type ServerConfig struct {
	Ip   string
	Port string
}

var Cfg *tomlConfig

func init() {
	Cfg = new(tomlConfig)

	Cfg.System.AppName = "jmy-go-blog"
	Cfg.System.Version = 1.0
	currentDir, _ := os.Getwd()
	Cfg.System.CurrentDir = currentDir

	_, err := toml.DecodeFile("config/config.toml", &Cfg)
	if err != nil {
		panic(err)
	}
}
