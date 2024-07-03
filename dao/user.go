package dao

import (
	"goblog/models"
	"log"
)

func GetUserNameById(uid int) (string, error) {
	var name string
	err := DB.QueryRow("select user_name from blog_user where uid = ?", uid).Scan(&name)
	if err != nil {
		log.Println("GetUserNameById 查询出错", err)
		return "", err
	}
	return name, nil
}

func GetUser(userName, password string) (*models.User, error) {
	user  := &models.User{}
	row := DB.QueryRow(
		"select * from blog_user where user_name = ? and passwd = ? limit 1",
		userName,
		password,
	)
	if row.Err() != nil {
		log.Println("GetUser 查询出错", row.Err())
		return nil, row.Err()
	}
	err := row.Scan(&user.Uid, &user.UserName, &user.Passwd, &user.Avatar, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		log.Println("GetUserNameById 查询出错", err)
		return nil, err
	}
	return user, nil
}
