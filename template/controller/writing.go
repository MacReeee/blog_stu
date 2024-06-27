package controller

import (
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"goblog/common"
	"goblog/config"
	"goblog/dao"
	"net/http"
)

func Writing(w http.ResponseWriter, r *http.Request)  {
	writing := common.Template.Writing
	categorys := dao.GetCategorys()
	m := make(map[string]interface{})
	m["categorys"] = categorys
	m["CdnURL"] = config.Cfg.System.CdnURL
	m["Title"] = config.Cfg.Viewer.Title
	writing.WriteData(w,m)
}

func QiniuToken(w http.ResponseWriter,r *http.Request){
	//自定义凭证有效期（示例2小时，Expires 单位为秒，为上传凭证的有效时间）
	bucket := "mszlu"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200 //示例2小时有效期
	mac := qbox.NewMac(config.Cfg.System.QiniuAccessKey, config.Cfg.System.QiniuSecretKey)
	upToken := putPolicy.UploadToken(mac)
	common.ReturnSuccess(w,upToken)
}