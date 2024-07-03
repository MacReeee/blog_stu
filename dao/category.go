package dao

import (
	"goblog/models"
	"log"
)

func GetCategoryNameById(cid int) (string, error) {
	var name string
	err := DB.QueryRow("select name from blog_category where cid = ?", cid).Scan(&name)
	if err != nil {
		log.Println("GetCategoryNameById 查询出错", err)
		return "", err
	}
	return name, nil
}

func GetUserNameById(uid int) (string, error) {
	var name string
	err := DB.QueryRow("select user_name from blog_user where uid = ?", uid).Scan(&name)
	if err != nil {
		log.Println("GetUserNameById 查询出错", err)
		return "", err
	}
	return name, nil
}
func GetAllCategory() ([]models.Category, error) {
	rows, err := DB.Query("select * from blog_category")
	if err != nil {
		log.Println("GetAllCategory 查询出错", err)
		return nil, err
	}
	var categorys []models.Category
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.Cid, &category.Name, &category.CreateAt, &category.UpdateAt)
		if err != nil {
			log.Println("GetAllCategory 取值出错", err)
			return nil, err
		}
		categorys = append(categorys, category)
	}
	return categorys, nil
}
