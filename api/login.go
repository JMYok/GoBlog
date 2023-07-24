package api

import (
	"GoBlog/common"
	"GoBlog/service"
	"net/http"
)

func (*APIHandler) Login(w http.ResponseWriter, r *http.Request) {
	params := common.GetRequestJsonParam(r)
	userName := params["username"].(string)
	passwd := params["passwd"].(string)
	loginRes, err := service.Login(userName, passwd)
	if err != nil {
		common.Error(w, err)
	}
	common.Success(w, loginRes)
}
