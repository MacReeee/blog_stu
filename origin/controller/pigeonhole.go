package controller

import (
	"goblog/common"
	"goblog/config"
	"goblog/dao"
	"goblog/models"
	"goblog/service"
	"net/http"
)

func (*HTMLApi) Pigeonhole(w http.ResponseWriter,r *http.Request)  {
	pigeonhole := common.Template.Pigeonhole
	var pd = &models.PigeonholeData{
		config.Cfg.Viewer,
		config.Cfg.System,
		dao.GetCategorys(),
		service.PostByMonth(),
		}
	pigeonhole.WriteData(w,pd)
}
