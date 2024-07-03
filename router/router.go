package router

import (
	"encoding/json"
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
	http.HandleFunc("/c/", views.HTML.Category)
	http.HandleFunc("/", views.HTML.Index)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
