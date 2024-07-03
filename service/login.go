package service

import (
	"errors"
	"goblog/dao"
	"goblog/models"
	"goblog/utils"
)

func Login(userName, password string) (*models.LoginResponse, error) {
	password = utils.Md5Crypt(password, "yny")
	user, _ := dao.GetUser(userName, password)
	// fmt.Println(user)
	if user == nil {
		return nil, errors.New("用户名或密码错误")
	}
	uid := user.Uid
	//生成token
	token, err := utils.Award(&uid)
	if err != nil {
		return nil, errors.New("token生成失败")
	}
	loginResponse := &models.LoginResponse{
		Token: token,
		UserInfo: models.UserInfo{
			Uid:      user.Uid,
			UserName: user.UserName,
			Avatar:   user.Avatar,
		},
	}
	return loginResponse, nil
}
