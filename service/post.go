package service

import (
	"goblog/config"
	"goblog/dao"
	"goblog/models"
	"html/template"
	"log"
)

func SavePost(post *models.Post) {
	dao.SavePost(post)
}

func UodatePost(post *models.Post) {
	dao.UpdatePost(post)
}

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

func Writing() (*models.WritingRes, error) {
	categorys, err := dao.GetAllCategory()
	if err != nil {
		log.Println("Writing 获取分类出错", err)
		return nil, err
	}
	wr := &models.WritingRes{
		Title:     config.Cfg.Viewer.Title,
		CdnURL:    config.Cfg.System.CdnURL,
		Categorys: categorys,
	}
	return wr, nil
}

func SearchPost(condition string) []models.SearchResp {
	posts, err := dao.GetPostSearch(condition)
	if err != nil {
		log.Println("SearchPost 查询出错", err)
		return nil
	}
	var searchResps []models.SearchResp
	for _, post := range posts {
		searchResps = append(searchResps, models.SearchResp{
			Pid:   post.Pid,
			Title: post.Title,
		})
	}
	return searchResps
}
