package views

import (
	"errors"
	"goblog/common"
	"goblog/service"
	"log"
	"net/http"
	"time"

	"strconv"
)

func isODD(num int) bool {
	return num%2 == 0
}

func getNextName(strs []string, index int) string {
	return strs[index+1]
}

func date(layout string) string {
	return time.Now().Format(layout)
}

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	//数据库查询
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败: ", err)
		index.WriteData(w, errors.New("系统错误,请联系管理员: "+string(err.Error())))
		return
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	pageSize := 10
	hr, err := service.GetAllIndexInfo(page, pageSize)
	if err != nil {
		log.Println("首页获取数据出错: ", err)
		index.WriteData(w, errors.New("系统错误,请联系管理员: "+string(err.Error())))
	}
	index.WriteData(w, hr)
}
