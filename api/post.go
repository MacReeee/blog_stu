package api

import (
	"errors"
	"goblog/common"
	"goblog/models"
	"goblog/service"
	"goblog/utils"
	"net/http"
	"strconv"
	"time"
)

func (*Api) SaveAndUpdatePoset(w http.ResponseWriter, r *http.Request) {
	//获取用户ID，判断登录状态
	token := r.Header.Get("Authorization")
	_, claim, err := utils.ParseToken(token)
	if err != nil {
		common.Error(w, errors.New("请先登录"))
		return
	}

	method := r.Method
	switch method {
	case http.MethodPost: // post -> save
		params := common.GetRequestJsonParam(r)
		if params["type"] == nil {
			params["type"] = 0.0
		}
		// // 测试获取所有参数
		// fmt.Println("params的参数如下：")
		// for k, _ := range params {
		// 	log.Println(k)
		// }
		// log.Println(params["type"])
		cid := params["categoryId"].(string)
		categoryId, _ := strconv.Atoi(cid)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		post := &models.Post{
			Pid:        -1,
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: categoryId,
			UserId:     claim.Uid,
			ViewCount:  0,
			Type:       int(postType),
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		service.SavePost(post)
		common.Success(w, post)
	case http.MethodPut: // put -> update
	}
}
