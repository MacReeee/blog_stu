package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
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

func indexHtmlHandler(w http.ResponseWriter, r *http.Request) {
	t := template.New("index.html")
	//1. 拿到当前路径
	path, _ := os.Getwd()
	t, _ = t.ParseFiles(path + "/template/index.html")
	t.Execute(w, indexData{"测试", "现在是入门教程测试indexData"})
}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/index.html", indexHtmlHandler)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
