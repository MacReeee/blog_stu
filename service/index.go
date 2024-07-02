package service

import (
	"goblog/config"
	"goblog/dao"
	"goblog/models"
	"html/template"
)

func GetAllIndexInfo(page, pageSize int) (*models.HomeResponse, error) {
	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}
	posts, err := dao.GetPostPage(page, pageSize)
	var postsMore []models.PostMore
	for _, post := range posts {
		categoryName, _ := dao.GetCategoryNameById(post.CategoryId)
		userName, _ := dao.GetUserNameById(post.UserId)
		content := []rune(post.Content)
		if len(content) > 200 {
			content = content[:200]
		}
		var postMore = models.PostMore{
			Pid:          post.Pid,
			Title:        post.Title,
			Slug:         post.Slug,
			Content:      template.HTML(content),
			CategoryId:   post.CategoryId,
			CategoryName: categoryName,
			UserId:       post.UserId,
			UserName:     userName,
			ViewCount:    post.ViewCount,
			Type:         post.Type,
			CreateAt:     models.DateDay(post.CreateAt),
			UpdateAt:     models.DateDay(post.UpdateAt),
		}
		postsMore = append(postsMore, postMore)
	}
	total := dao.CountGetAllPost()
	pagesNum := (total-1)/10 + 1
	pages := make([]int, pagesNum)
	for i := 0; i < pagesNum; i++ {
		pages[i] = i+1
	}
	var hr = &models.HomeResponse{
		Viewer:    config.Cfg.Viewer,
		Categorys: categorys,
		Posts:     postsMore,
		Total:     total,
		Page:      page,
		Pages:     pages,
		PageEnd:   page != pagesNum,
	}
	return hr, nil
}
//