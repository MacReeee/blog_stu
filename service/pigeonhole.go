package service

import (
	"goblog/config"
	"goblog/dao"
	"goblog/models"
)

func FindPostPigeonhole() *models.Pigeonhole {
	categorys, _ := dao.GetAllCategory()
	posts, _ := dao.GetPostAll()
	pigeonholeMap := make(map[string][]models.Post)
	for _, post := range posts {
		at := post.CreateAt
		month := at.Format("2006-01")
		pigeonholeMap[month] = append(pigeonholeMap[month], post)
	}
	return &models.Pigeonhole{
		Viewer:       config.Cfg.Viewer,
		SystemConfig: config.Cfg.System,
		Categorys:    categorys,
		Lines:        pigeonholeMap,
	}
}
