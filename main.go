package main

import (
	"goblog/common"
	"goblog/router"
	"log"
	"net/http"
)

func main() {
	common.LoadTemplate()
	log.SetFlags(log.Ldate | log.Lshortfile)
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	//路由
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
