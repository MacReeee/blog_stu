package main

import (
	"encoding/json"
	"goblog/config"
	"goblog/models"
	"html/template"
	"log"
	"net/http"
	"time"
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

func isODD(num int) bool {
	return num%2 == 0
}

func getNextName(strs []string, index int) string {
	return strs[index+1]
}

func date(layout string) string {
	return time.Now().Format(layout)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t := template.New("index.html")
	//1. 拿到当前路径
	path := config.Cfg.System.CurrentDir
	//2.拿到各个页面需要的前端页面
	home := path + "/template/home.html"
	header := path + "/template/layout/header.html"
	footer := path + "/template/layout/footer.html"
	personal := path + "/template/layout/personal.html"
	post := path + "/template/layout/post-list.html"
	pagination := path + "/template/layout/pagination.html"
	t.Funcs(template.FuncMap{"isODD": isODD, "getNextName": getNextName, "date": date})
	t, err := t.ParseFiles(path+"/template/index.html", home, header, footer, personal, post, pagination)
	if err != nil {
		log.Println("错误:", err)
	}

	//假数据
	var categorys = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}
	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go博客",
			Content:      "内容",
			UserName:     "张三",
			ViewCount:    123,
			CreateAt:     "2022-02-20",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}
	var hr = &models.HomeResponse{
		Viewer:    config.Cfg.Viewer,
		Categorys: categorys,
		Posts:     posts,
		Total:     1,
		Page:      1,
		Pages:     []int{1},
		PageEnd:   true,
	}
	t.Execute(w, hr)
}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/", indexHandler)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
