package dao

import (
	"goblog/models"
	"log"
)

func SavePost(post *models.Post) {
	res, err := DB.Exec("insert into blog_post(title,content,markdown,category_id,user_id,view_count,type,slug,create_at,update_at) "+
		"values(?,?,?,?,?,?,?,?,?,?)",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.UserId,
		post.ViewCount,
		post.Type,
		post.Slug,
		post.CreateAt,
		post.UpdateAt,
	)
	if err != nil {
		log.Println("SavePost 插入出错", err)
	}
	pid, _ := res.LastInsertId()
	post.Pid = int(pid)
}

func UpdatePost(post *models.Post) {
	_, err := DB.Exec("UPDATE blog_post SET title=?, content=?, markdown=?, category_id=?, type=?, slug=?, update_at=? "+
		"WHERE pid=?",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.Type,
		post.Slug,
		post.UpdateAt,
		post.Pid,
	)
	if err != nil {
		log.Println("UpdatePost 更新出错", err)
	}
}

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

func GetPostById(pid int) (*models.Post, error) {
	row := DB.QueryRow("select * from blog_post where pid = ?", pid)
	if row.Err() != nil {
		log.Println("GetPostPage 查询出错", row.Err())
		return nil, row.Err()
	}
	post := &models.Post{}
	err := row.Scan(
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
	return post, nil
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

func CountGetAllPostBySlug(slug string) int {
	rows := DB.QueryRow("select count(1) from blog_post where slug = ?", slug)
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

func GetPostAll() ([]models.Post, error) {
	rows, err := DB.Query("select * from blog_post")
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

func GetPostPageBySlug(slug string, page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_post where slug = ? limit ?,?", slug, page, pageSize)
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
