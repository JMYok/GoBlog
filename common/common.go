package common

import (
	"GoBlog/config"
	"GoBlog/models/template"
	"sync"
)

var Template models.HtmlTemplate

func LoadTemplate() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		//耗时
		var err error
		Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
		if err != nil {
			panic(err)
		}
	}()
	wg.Wait()
}
