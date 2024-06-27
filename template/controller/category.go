package controller

import (
	"goblog/common"
	"goblog/config"
	"goblog/dao"
	"goblog/models"
	"goblog/service"
	"net/http"
	"strconv"
	"strings"
)

var HTML = &HTMLApi{}

type HTMLApi struct {
}
var API = &Api{}

type Api struct {
}

func (*HTMLApi) Category(w http.ResponseWriter, r *http.Request)  {
	path := r.URL.Path
	id := strings.TrimPrefix(path, "/c/")
	cId,_:= strconv.Atoi(id)
	_ = r.ParseForm()
	page := r.Form.Get("page")
	if page == ""{
		page = "1"
	}
	currentPage,_ := strconv.Atoi(page)
	cName := dao.GetCategoryNameById(cId)
	categorys := dao.GetCategorys()
	post,total := service.PostPageByCategory(currentPage,10,cId)
	pagesAll := ((total-1)/10) + 1
	pages := []int{}
	for i:=1;i<=pagesAll;i++ {
		pages = append(pages,i)
	}
	hd := models.HomeData{
		config.Cfg.Viewer,
		categorys,
		post,
		total,
		currentPage,
		pages,
		currentPage != pagesAll,
	}
	var categoryData = &models.CategoryData{
		hd,
		cName,
	}

	common.Template.Category.WriteData(w,categoryData)
}