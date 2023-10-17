package service

import (
	"GoBlog/dao"
	"GoBlog/models"
	"GoBlog/utils"
	"errors"
)

func Login(userName, passwd string) (*models.LoginRes, error) {
	user, err := dao.GetUser(userName, passwd)
	if user == nil || err != nil {
		return nil, errors.New("账号秘密不正确")
	}
	uid := user.Uid
	//生成token
	token, err := utils.Award(&uid)
	if err != nil {
		return nil, errors.New("token未能生成")
	}
	var userInfo models.UserInfo
	userInfo.Uid = uid
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar
	var lr = &models.LoginRes{
		Token:    token,
		UserInfo: userInfo,
	}
	return lr, nil
}
