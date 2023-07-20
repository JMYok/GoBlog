package common

import (
	models "GoBlog/models/template"
)

var Template models.HtmlTemplate

func LoadTemplate() {
	Template = models.InitTemplate("template/")
}
