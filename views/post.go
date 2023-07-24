package views

import (
	"GoBlog/common"
	"GoBlog/service"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) DetailPage(w http.ResponseWriter, r *http.Request) {
	detailTemplate := common.Template.Detail

	pathStr := r.URL.Path
	pidStr, _ := strings.CutPrefix(pathStr, "/p/")
	pidStr = strings.TrimSuffix(pidStr, ".html")
	pid, _ := strconv.Atoi(pidStr)

	postRes, err := service.GetPostDetail(pid)

	if err != nil {
		log.Println("post detail获取数据出错：", err)
		detailTemplate.WriteError(w, errors.New("系统错误,请联系站长！"))
	}

	detailTemplate.WriteData(w, *postRes)
}
