package dao

import (
	"goblog/models"
	"log"
)

func CountGetAllPostByCategoryID(cid int) int {
	rows := DB.QueryRow("select count(1) from blog_post where category_id = ? ", cid)
	count := 0
	err := rows.Scan(&count)
	if err != nil {
		log.Println("CountGetAllPost 查询出错", err)
		return 0
	}
	return count
}

func CountGetAllPost() int {
	rows := DB.QueryRow("select count(1) from blog_post")
	count := 0
	err := rows.Scan(&count)
	if err != nil {
		log.Println("CountGetAllPost 查询出错", err)
		return 0
	}
	return count
}

func GetPostPage(page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_post limit ?,?", page, pageSize)
	if err != nil {
		log.Println("GetPostPage 查询出错", err)
		return nil, err
	}

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			log.Println("GetPostPage 取值出错", err)
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostPageBycategoryID(cid, page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_post where category_id = ? limit ?,?", cid, page, pageSize)
	if err != nil {
		log.Println("GetPostPage 查询出错", err)
		return nil, err
	}

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			log.Println("GetPostPage 取值出错", err)
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
