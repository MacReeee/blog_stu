package service

import (
	"goblog/dao"
	"goblog/models"
	"goblog/utils"
	"html/template"
)

func PostByMonth() *map[string][]models.Post  {
	posts := dao.GetPostAll()
	lines := make(map[string][]models.Post)
	for _,post :=range posts{
		month := post.CreateAt.Format("2006-01")
		lines[month] = append(lines[month],post)
	}
	return &lines
}
func PostPageByCategory(page int,pageSize int,categoryId int) ([]models.PostMore,int)  {
	posts := dao.GetPostPageCategory(page,pageSize,categoryId)
	total := dao.GetPostCountCategory(categoryId);
	var postMores []models.PostMore
	for _,post := range posts {
		var pm models.PostMore
		pm.Pid = post.Pid
		pm.ViewCount = post.ViewCount
		pm.CategoryId = post.CategoryId
		pm.CategoryName = dao.GetCategoryNameById(post.CategoryId);
		content := []rune(post.Content)
		if len(content) > 100 {
			content = []rune(post.Content)[:100]
		}
		pm.Content = template.HTML(content)
		pm.Title = post.Title
		pm.Slug = post.Slug
		pm.CreateAt = utils.Format(post.CreateAt)
		pm.UserName = dao.GetUserNameById(post.UserId)
		postMores = append(postMores,pm)
	}
	return postMores,total
}

func PostPageBySlug(page int,pageSize int,slug string) ([]models.PostMore,int) {
	posts := dao.GetPostPageBySlug(page,pageSize,slug)
	total := dao.GetPostCountBySlug(slug);
	var postMores []models.PostMore
	for _,post := range posts {
		var pm models.PostMore
		pm.Pid = post.Pid
		pm.ViewCount = post.ViewCount
		pm.CategoryId = post.CategoryId
		pm.CategoryName = dao.GetCategoryNameById(post.CategoryId);
		content := []rune(post.Content)
		if len(content) > 100 {
			content = []rune(post.Content)[:100]
		}
		pm.Content = template.HTML(content)
		pm.Title = post.Title
		pm.Slug = post.Slug
		pm.CreateAt = utils.Format(post.CreateAt)
		pm.UserName = dao.GetUserNameById(post.UserId)
		postMores = append(postMores,pm)
	}
	return postMores,total
}
func PostPage(page int,pageSize int) ([]models.PostMore,int)  {
	posts := dao.GetPostPage(page,pageSize)
	total := dao.GetPostCount();
	var postMores []models.PostMore
	for _,post := range posts {
		var pm models.PostMore
		pm.Pid = post.Pid
		pm.ViewCount = post.ViewCount
		pm.CategoryId = post.CategoryId
		pm.CategoryName = dao.GetCategoryNameById(post.CategoryId);
		content := []rune(post.Content)
		if len(content) > 100 {
			content = []rune(post.Content)[:100]
		}
		pm.Content = template.HTML(content)
		pm.Title = post.Title
		pm.Slug = post.Slug
		pm.CreateAt = utils.Format(post.CreateAt)
		pm.UserName = dao.GetUserNameById(post.UserId)
		postMores = append(postMores,pm)
	}
	return postMores,total
}
