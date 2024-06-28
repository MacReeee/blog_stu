package dao

import (
	"errors"
	"goblog/models"
)

func Login(req *models.LoginReq) (*models.User,*models.DBError)  {
	row := DB.QueryRow("select * from blog_user where user_name=? and passwd=? limit 1", req.Name, req.Passwd)
	user := new(models.User)
	err := row.Scan(&user.Uid,&user.UserName,&user.Passwd,&user.Avatar,&user.CreateAt,&user.UpdateAt)
	if err != nil {
		return nil,&models.DBError{Err: errors.New("未查询到数据"), IsNilError: true}
	}
	return user,nil
	//
}
func GetUserNameById(id int) string  {
	row := DB.QueryRow("select user_name from blog_user where uid=?", id)
	var username string
	_ = row.Scan(&username)
	return username
}