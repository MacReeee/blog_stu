package views

import (
	"errors"
	"goblog/common"
	"goblog/service"
	"log"
	"net/http"
	"time"
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
	hr, err := service.GetAllIndexInfo()
	if err != nil {
		log.Println("首页获取数据出错: ", err)
		index.WriteData(w, errors.New("系统错误,请联系管理员: "+string(err.Error())))
	}
	index.WriteData(w, hr)
}
