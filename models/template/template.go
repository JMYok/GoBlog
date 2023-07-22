package models

import (
	"html/template"
	"io"
	"log"
	"time"
)

type TemplateBlog struct {
	*template.Template
}

type HtmlTemplate struct {
	Index      TemplateBlog
	Category   TemplateBlog
	Custom     TemplateBlog
	Detail     TemplateBlog
	Login      TemplateBlog
	Pigeonhole TemplateBlog
	Writing    TemplateBlog
}

func (t *TemplateBlog) WriteData(w io.Writer, data interface{}) {
	err := t.Execute(w, data)
	if err != nil {
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func (t *TemplateBlog) WriteError(w io.Writer, err error) {
	if err != nil {
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func InitTemplate(templateDir string) (HtmlTemplate, error) {
	tp, err := readTemplate(
		[]string{"index", "category", "custom", "detail", "login", "pigeonhole", "writing"},
		templateDir,
	)
	var htmlTemplate HtmlTemplate

	if err != nil {
		return htmlTemplate, err
	}

	htmlTemplate.Index = tp[0]
	htmlTemplate.Category = tp[1]
	htmlTemplate.Custom = tp[2]
	htmlTemplate.Detail = tp[3]
	htmlTemplate.Login = tp[4]
	htmlTemplate.Pigeonhole = tp[5]
	htmlTemplate.Writing = tp[6]

	return htmlTemplate, nil
}

func readTemplate(templates []string, templateDir string) ([]TemplateBlog, error) {
	var tbs []TemplateBlog

	for _, view := range templates {
		viewName := view + ".html"
		t := template.New(viewName)

		//页面共用部分
		header := templateDir + "layout/header.html"
		footer := templateDir + "layout/footer.html"
		pagination := templateDir + "layout/pagination.html"
		personal := templateDir + "layout/personal.html"
		postList := templateDir + "layout/post-list.html"
		home := templateDir + "home.html"

		//映射页面上的方法
		t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date, "dateDay": DateDay})

		t, err := t.ParseFiles(templateDir+viewName, header, footer, pagination, personal, postList, home)
		if err != nil {
			log.Println("解析模板出错：", err)
			return nil, err
		}
		tbs = append(tbs, TemplateBlog{t})
	}
	return tbs, nil
}

func DateDay(date time.Time) string {
	//格式化（go特色）
	return date.Format("2006-01-02 15:30:05")
}

func GetNextName(navigation []string, index int) string {
	return navigation[index+1]
}

func IsODD(num int) bool {
	return num%2 == 0
}

func Date(layout string) string {
	return time.Now().Format(layout)
}
