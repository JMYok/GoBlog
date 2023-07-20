package common

import (
	"GoBlog/config"
	models "GoBlog/models/template"
)

var Template models.HtmlTemplate

func LoadTemplate() {
	Template = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
}
