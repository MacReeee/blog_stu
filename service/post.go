package service

import (
	"goblog/config"
	"goblog/dao"
	"goblog/models"
	"html/template"
)

func GetPostDetail(pid int) (*models.PostRes, error) {
	post, err := dao.GetPostById(pid)
	if err != nil {
		return nil, err
	}
	categoryName, _ := dao.GetCategoryNameById(post.CategoryId)
	userName, _ := dao.GetUserNameById(post.UserId)
	postRes := &models.PostRes{
		Viewer:       config.Cfg.Viewer,
		SystemConfig: config.Cfg.System,
		Article: models.PostMore{
			Pid:          post.Pid,
			Title:        post.Title,
			Slug:         post.Slug,
			Content:      template.HTML(post.Content),
			CategoryId:   post.CategoryId,
			CategoryName: categoryName,
			UserId:       post.UserId,
			UserName:     userName,
			ViewCount:    post.ViewCount,
			Type:         post.Type,
			CreateAt:     models.DateDay(post.CreateAt),
			UpdateAt:     models.DateDay(post.UpdateAt),
		},
	}
	return postRes, nil
}
