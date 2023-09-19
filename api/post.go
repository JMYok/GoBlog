package api

import (
	"GoBlog/common"
	"GoBlog/context"
	"GoBlog/dao"
	"GoBlog/models"
	"GoBlog/service"
	"GoBlog/utils"
	"errors"
	"net/http"
	"strconv"
	"time"
)

func (*APIHandler) SearchPost(ctx *context.MsContext) {
	_ = ctx.Request.ParseForm()
	condition, _ := ctx.GetForm("val")
	searchResp := service.SearchPost(condition)
	common.Success(ctx.W, searchResp)
}

func (*APIHandler) GetPost(ctx *context.MsContext) {
	w := ctx.W

	pIdStr := ctx.GetPathVariable("pid")
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

func (*APIHandler) SaveAndUpdatePost(ctx *context.MsContext) {
	r := ctx.Request
	w := ctx.W

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
		cId, _ := strconv.Atoi(ctx.GetJson("categoryId").(string))
		content := ctx.GetJson("content").(string)
		markdown := ctx.GetJson("markdown").(string)
		slug := ctx.GetJson("slug").(string)
		title := ctx.GetJson("title").(string)
		postType := int(ctx.GetJson("type").(float64))
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
		cId, _ := strconv.Atoi(ctx.GetJson("categoryId").(string))
		content := ctx.GetJson("content").(string)
		markdown := ctx.GetJson("markdown").(string)
		slug := ctx.GetJson("slug").(string)
		title := ctx.GetJson("title").(string)
		postType := int(ctx.GetJson("type").(float64))
		pid := int(ctx.GetJson("pid").(float64))
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
