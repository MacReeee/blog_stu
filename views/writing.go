package views

import (
	"goblog/common"
	"goblog/service"
	"net/http"
)

func (*HTMLApi) Writing(w http.ResponseWriter, r *http.Request) {
	writing := common.Template.Writing
	wr, _ := service.Writing()
	writing.WriteData(w, wr)
}
