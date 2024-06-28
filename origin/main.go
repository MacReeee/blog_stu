package main

import (
	"goblog/common"
	"goblog/routers"
	"log"
	"net/http"
)

func init()  {
	common.Load()
}
func main() {
	routers.Routers()
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	if err := server.ListenAndServe();err != nil{
		log.Println(err)
	}
}
