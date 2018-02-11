package controller

import (
	"fmt"
	"mongo/proxy"
	"net/http"
)

// 获取用户列表
func HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := Context{w, r}
	if r.Method == "GET" {
		users, err := proxy.GetUSers()
		if err != nil {
			ctx.Error("查询失败", err)
		} else {
			ctx.Success("查询成功", users)
		}
	} else {
		ctx.NotFound()
	}
}

// 添加用户
func HandleAddUser(w http.ResponseWriter, r *http.Request) {
	ctx := Context{w, r}
	if r.Method == "POST" {
		r.ParseForm()
		account := r.Form.Get("account")
		password := r.Form.Get("password")

		fmt.Println("Account:", account)
		fmt.Println("Password:", password)

		user, err := proxy.AddUser(account, password)
		if err != nil {
			ctx.Error("添加失败", err)
		} else {
			ctx.Success("添加成功", user)
		}
	} else {
		ctx.NotFound()
	}
}
