package dao

import (
	"goblog/models"
	"log"
)

func GetCategorys() []models.Category{
	ret,err := DB.Query("select * from blog_category")
	if err != nil {
		log.Println(err)
		return nil
	}
	var cs  []models.Category
	for ret.Next() {
		var cat models.Category
		_ = ret.Scan(&cat.Cid,&cat.Name,&cat.CreateAt,&cat.UpdateAt)
		cs = append(cs,cat)
	}
	return cs
}
func GetCategoryNameById(id int) string {
	row := DB.QueryRow("select name from blog_category where cid=?",id)
	var name string
	_ = row.Scan(&name)
	return name
}