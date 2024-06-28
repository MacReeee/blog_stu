package routers

import (
	"goblog/controller"
	"net/http"
)

func Routers()  {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/login", controller.LoginHtml)
	http.HandleFunc("/api/v1/login", controller.Login)
	http.HandleFunc("/writing", controller.Writing)
	http.HandleFunc("/api/v1/qiniu/token",controller.QiniuToken)
	http.HandleFunc("/api/v1/post", controller.AddOrUpdate)
	http.HandleFunc("/api/v1/post/", controller.GetPost)
	http.HandleFunc("/p/",controller.PostDetail)
	http.HandleFunc("/c/",controller.HTML.Category)
	http.HandleFunc("/pigeonhole",controller.HTML.Pigeonhole)
	http.HandleFunc("/api/v1/post/search",controller.API.PostSearch)
	http.Handle("/resource/",http.StripPrefix("/resource/",http.FileServer(http.Dir("public/resource/"))))
}
