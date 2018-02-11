package main

import (
	"fmt"
	"log"
	"net/http"

	"controller"
	"mongo"
	"config"
)

func main() {
	// 链接数据库 "mongo/mongo.go"
	go mongo.LinkDb()

	// 处理请求
	http.HandleFunc("/addUser", controller.HandleAddUser)
	http.HandleFunc("/users", controller.HandleGetUsers)

	// 开启服务器
	sPort := ":" + config.Conf.Server.Port
	fmt.Printf("[Server is opened]: localhost%s \n", sPort)
	err := http.ListenAndServe(sPort, nil)
	
	if err != nil {
		log.Fatal("Server is open failed: ", err)
	}
}
