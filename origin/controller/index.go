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

func Index(w http.ResponseWriter, r *http.Request) {
	indexTemplate := common.Template.Index
	if err := r.ParseForm();err != nil {
		indexTemplate.WriteError(w, err)
	}
	path := r.URL.Path
	page := r.Form.Get("page")
	if page == ""{
		page = "1"
	}
	categorys := dao.GetCategorys()
	currentPage,_ := strconv.Atoi(page)
	slug := strings.TrimPrefix(path, "/")
	var post []models.PostMore
	var total int
	if slug != ""{
		//请求的是自定义的路径 文章的slug
		//查询文章的时候 需要按照slug查询
		post,total = service.PostPageBySlug(currentPage,10,slug)
	}else{
		post,total = service.PostPage(currentPage,10)
	}

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
	indexTemplate.WriteData(w,hd)
}
