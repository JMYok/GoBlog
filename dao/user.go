package dao

import "log"

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
