package service

import (
	"errors"
	"fmt"
	"goblog/dao"
	"goblog/models"
)

func Login(userName, password string) (*models.LoginResponse, error) {
	user, _ := dao.GetUser(userName, password)
	fmt.Println(user)
	if user == nil {
		return nil, errors.New("用户名或密码错误")
	}
	loginResponse := &models.LoginResponse{}
	return loginResponse, nil
}
