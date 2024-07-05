package router

import (
	"encoding/json"
	"goblog/api"
	"goblog/views"
	"net/http"
)

type indexData struct {
	Title string
	Disc  string
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	indexData := indexData{"测试", "现在是入门教程测试indexData"}
	jsonStr, _ := json.Marshal(indexData)
	w.Write([]byte(jsonStr))
}

func Router() {
	http.HandleFunc("/login", views.HTML.Login)
	http.HandleFunc("/c/", views.HTML.Category)
	http.HandleFunc("/p/", views.HTML.Detail)
	http.HandleFunc("/writing", views.HTML.Writing)
	http.HandleFunc("/pigeonhole", views.HTML.Pigeonhole)
	http.HandleFunc("/", views.HTML.Index)
	http.HandleFunc("/api/v1/login", api.API.Login)
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/post/", api.API.GetPost)
	http.HandleFunc("/api/v1/post/search", api.API.SearchPost)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
