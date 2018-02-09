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
	go mongo.LinkDb()

	http.HandleFunc("/addUser", controller.HandleAddUser)
	http.HandleFunc("/users", controller.HandleGetUsers)

	sPort := ":" + config.Conf.Server.Port
	fmt.Printf("[Server is opened]: localhost%s \n", sPort)
	err := http.ListenAndServe(sPort, nil)
	if err != nil {
		log.Fatal("Server is open failed: ", err)
	}
}
