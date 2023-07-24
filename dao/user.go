package dao

import (
	"GoBlog/models"
	"log"
)

func GetUserNameById(userId int) (string, error) {
	rows, err := DB.Query("select user_name from blog_user where uid=?", userId)
	if err != nil {
		log.Println("GetUserNameById出错:", err)
		return "", err
	}
	rows.Next()
	var userName string
	_ = rows.Scan(&userName)
	return userName, err
}

func GetUser(userName, passwd string) (*models.User, error) {
	row := DB.QueryRow("select * from blog_user where user_name=? and passwd=? limit 1", userName, passwd)
	if row.Err() != nil {
		log.Println("GetUserNameById出错:", row.Err())
		return nil, row.Err()
	}
	var user = &models.User{}
	err := row.Scan(&user.Uid, &user.UserName, &user.PassWd, &user.Avatar, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil
}
