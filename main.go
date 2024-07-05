package main

import (
	"goblog/common"
	"goblog/router"
	"log"
	"net/http"
)

func init() {
	log.SetFlags(log.Ldate | log.Llongfile)
}

func main() {
	common.LoadTemplate()
	
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	//路由
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
