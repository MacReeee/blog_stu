package views

import (
	"errors"
	"goblog/common"
	"goblog/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Category(w http.ResponseWriter, r *http.Request) {
	categoryTemplate := common.Template.Category
	path := r.URL.Path
	cidStr := strings.TrimPrefix(path, "/c/")
	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		categoryTemplate.WriteData(w, errors.New("不识别此请求路径,请联系管理员: "+string(err.Error())))
		return
	}

	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败: ", err)
		categoryTemplate.WriteData(w, errors.New("系统错误,请联系管理员: "+string(err.Error())))
		return
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	pageSize := 10
	categoryResponse, err := service.GetPostsByCategoryId(cid, page, pageSize)
	if err != nil {
		log.Println("分类页获取数据出错: ", err)
		categoryTemplate.WriteData(w, errors.New("系统错误,请联系管理员: "+string(err.Error())))
		return
	}
	categoryTemplate.WriteData(w, categoryResponse)
}
