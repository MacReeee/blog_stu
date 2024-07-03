package views

import (
	"errors"
	"goblog/common"
	"goblog/service"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Detail(w http.ResponseWriter, r *http.Request) {
	detail := common.Template.Detail
	path := r.URL.Path
	pidStr := strings.TrimPrefix(path, "/p/")
	pidStr = strings.TrimSuffix(pidStr, ".html")
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		detail.WriteData(w, errors.New("不识别此请求路径,请联系管理员: "+string(err.Error())))
		return
	}
	postResponse, err := service.GetPostDetail(pid)
	if err != nil {
		detail.WriteData(w, errors.New("系统错误,请联系管理员: "+string(err.Error())))
		return
	}
	detail.WriteData(w, postResponse)
}
