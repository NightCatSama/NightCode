package main

import (
	"controller"
	"log"
	"mongo"
	"net/http"
)

func main() {
	go mongo.LinkDb()

	http.HandleFunc("/addUser", controller.HandleAddUser)
	http.HandleFunc("/users", controller.HandleGetUsers)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("Server is open failed: ", err)
	}
}
