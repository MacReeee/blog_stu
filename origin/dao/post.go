package dao

import "goblog/models"
func GetPostCount() int{
	row := DB.QueryRow("select count(1) from blog_post")
	var total int
	_ = row.Scan(&total)
	return total
}
func GetPostCountBySlug(slug string) int{
	row := DB.QueryRow("select count(1) from blog_post where slug=?",slug)
	var total int
	_ = row.Scan(&total)
	return total
}
func GetPostCountCategory(categoryId int) int{
	row := DB.QueryRow("select count(1) from blog_post where category_id=?",categoryId)
	var total int
	_ = row.Scan(&total)
	return total
}

func GetPostPageCategory(page int ,pageSize int,categoryId int) []models.Post  {
	page = (page-1) * pageSize
	ret,_ := DB.Query("select * from blog_post where category_id=? limit ?,?",categoryId,page,pageSize)
	var posts []models.Post
	for ret.Next() {
		var post models.Post
		_ = ret.Scan(
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
			&post.UpdateAt)
		posts = append(posts,post)
	}
	return posts
}

func GetPostPageBySlug(page int ,pageSize int,slug string) []models.Post  {
	page = (page-1) * pageSize
	ret,_ := DB.Query("select * from blog_post where slug = ? limit ?,?",slug,page,pageSize)
	var posts []models.Post
	for ret.Next() {
		var post models.Post
		_ = ret.Scan(
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
			&post.UpdateAt)
		posts = append(posts,post)
	}
	return posts
}

func GetPostPage(page int ,pageSize int) []models.Post  {
	page = (page-1) * pageSize
	ret,_ := DB.Query("select * from blog_post limit ?,?",page,pageSize)
	var posts []models.Post
	for ret.Next() {
		var post models.Post
		_ = ret.Scan(
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
			&post.UpdateAt)
		posts = append(posts,post)
	}
	return posts
}

func GetPostAll() []models.Post  {
	ret,_ := DB.Query("select * from blog_post")
	var posts []models.Post
	for ret.Next() {
		var post models.Post
		_ = ret.Scan(
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
			&post.UpdateAt)
		posts = append(posts,post)
	}
	return posts
}
func UpdatePost(post *models.Post) error  {
	_,err := DB.Exec("update  blog_post set title=?," +
		"content=?," +
		"markdown=?," +
		"category_id=?," +
		"type=?," +
		"slug=?," +
		"update_at=? where user_id=? and pid=?",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.Type,
		post.Slug,
		post.UpdateAt,
		post.UserId,
		post.Pid)
	if err != nil {
		return err
	}
	return nil
}
func GetPostById(id int) (*models.Post,error)  {
	row := DB.QueryRow("select * from blog_post where pid=?",id)
	if row.Err() != nil {
		return nil,row.Err()
	}
	post := new(models.Post)
	_ = row.Scan(
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
		&post.UpdateAt)
	return post,nil
}
func SavePost(post *models.Post) error  {
	ret,err := DB.Exec("insert into blog_post (title,content,markdown,category_id,user_id,view_count,type,slug,create_at,update_at) values(?,?,?,?,?,?,?,?,?,?)",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.UserId,
		post.ViewCount,
		post.Type,
		post.Slug,
		post.CreateAt,
		post.UpdateAt)
	if err != nil {
		return err
	}
	id, err := ret.LastInsertId()
	if err != nil {
		return err
	}
	post.Pid = int(id)
	return nil
}

func PostSearch(search string) []models.Post  {
	ret ,_ := DB.Query("select * from blog_post where title like ?","%"+search+"%")
	var posts []models.Post
	if ret == nil {
		return posts
	}
	for ret.Next() {
		var post models.Post
		_ = ret.Scan(
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
			&post.UpdateAt)
		posts = append(posts,post)
	}
	return posts
}