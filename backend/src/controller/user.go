package controller

import (
	"fmt"
	"net/http"
	"mongo"
)

func HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := Context{w, r}
	if r.Method == "GET" {
		users, err := mongo.GetUSers()
		if err != nil {
			ctx.Error("查询失败")
		} else {
			ctx.Success("查询成功", users)
		}
	} else {
		ctx.NotFound()
	}
}

func HandleAddUser(w http.ResponseWriter, r *http.Request) {
	ctx := Context{w, r}
	fmt.Println(r.Method)
	if r.Method == "POST" {
		r.ParseForm()
		account := r.Form.Get("account")
		password := r.Form.Get("password")

		fmt.Println("Account:", account)
		fmt.Println("Password:", password)

		user, err := mongo.AddUser(account, password)
		if err != nil {
			ctx.Error("添加失败！")
		} else {
			ctx.Success("添加成功", user)
		}
	} else {
		ctx.NotFound()
	}
}
