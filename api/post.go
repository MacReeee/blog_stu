package api

import (
	"errors"
	"goblog/common"
	"goblog/config"
	"goblog/dao"
	"goblog/models"
	"goblog/service"
	"goblog/utils"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func (*Api) GetPost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pidStr := strings.TrimPrefix(path, "/api/v1/post/")
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		common.Error(w, errors.New("不识别此请求路径,请联系管理员: "+string(err.Error())))
		return
	}
	post, err := dao.GetPostById(pid)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, post)
}

func (*Api) SaveAndUpdatePost(w http.ResponseWriter, r *http.Request) {
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
			log.Println("type参数为空")
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
		params := common.GetRequestJsonParam(r)
		if params["type"] == nil {
			log.Println("type参数为空")
			params["type"] = 0.0
		}
		// fmt.Println("params的参数如下：")
		// for k, v := range params {
		// 	log.Println(k,v)
		// }
		cid := int(params["categoryId"].(float64))
		categoryId := cid
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pidFloat64 := params["pid"].(float64)
		pid := int(pidFloat64)
		post := &models.Post{
			Pid:        pid,
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
		service.UodatePost(post)
		common.Success(w, post)
	}
}

func (*Api) QiniuToken(w http.ResponseWriter, r *http.Request) {
	//自定义凭证有效期（示例2小时，Expires 单位为秒，为上传凭证的有效时间）
	bucket := "zhangsan"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200 //示例2小时有效期
	mac := qbox.NewMac(config.Cfg.System.QiniuAccessKey, config.Cfg.System.QiniuSecretKey)
	upToken := putPolicy.UploadToken(mac)
	common.Success(w, upToken)
}

func (*Api) SearchPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	condition := r.Form.Get("val")
	searchResp := service.SearchPost(condition)
	common.Success(w, searchResp)
}
