package api

import (
	"GoBlog/common"
	"GoBlog/dao"
	"GoBlog/models"
	"GoBlog/service"
	"GoBlog/utils"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (*APIHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path, "/api/v1/post/")
	pid, err := strconv.Atoi(pIdStr)
	if err != nil {
		common.Error(w, errors.New("不识别此请求路径"))
		return
	}
	post, err := dao.GetPostByPid(pid)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, post)
}

func (*APIHandler) SaveAndUpdatePost(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	token := r.Header.Get("Authorization")
	_, claim, err := utils.ParseToken(token)
	uid := claim.Uid

	if err != nil {
		common.Error(w, errors.New("登录已过期"))
		return
	}

	switch method {
	case http.MethodPost:
		params := common.GetRequestJsonParam(r)
		cId, _ := strconv.Atoi(params["categoryId"].(string))
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := int(params["type"].(float64))
		post := &models.Post{
			Pid:        -1,
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: cId,
			UserId:     uid,
			Type:       postType,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		service.SavePost(post)
		common.Success(w, post)
	case http.MethodPut:
		params := common.GetRequestJsonParam(r)
		cId, _ := strconv.Atoi(params["categoryId"].(string))
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := int(params["type"].(float64))
		pid := int(params["pid"].(float64))
		post := &models.Post{
			Pid:        pid,
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: cId,
			UserId:     uid,
			Type:       postType,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		service.UpdatePost(post)
		common.Success(w, post)
	}
}
