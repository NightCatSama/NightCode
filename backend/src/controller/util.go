package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
}

type Context struct {
	w http.ResponseWriter
	r *http.Request
}

// 设置响应头
func (ctx *Context) SetHeader(key string, val string, unique bool) {
	if unique {
		ctx.w.Header().Set(key, val)
	} else {
		ctx.w.Header().Add(key, val)
	}
}

// 格式化成 JSON
func (ctx *Context) fmtJson(code int, msg string, data interface{}) []byte {
	res := Response{
		Data: data,
		Msg:  msg,
		Code: code,
	}
	result, err := json.Marshal(res)
	if err != nil {
		fmt.Println("数据格式化出错，error:", err)
		return []byte(`{ "Msg": "未知错误", "Code": 500 }`)
	} else {
		fmt.Println(string(result))
		return result
	}
}

// 返回 400 错误
func (ctx *Context) Error(msg string, err error) {
	ctx.SetHeader("Content-Type", "application/json", false)
	ctx.w.WriteHeader(400)
	ctx.w.Write(ctx.fmtJson(400, msg, err))
}

// 返回 200 成功
func (ctx *Context) Success(msg string, data interface{}) {
	ctx.SetHeader("Content-Type", "application/json", false)
	ctx.w.WriteHeader(200)
	ctx.w.Write(ctx.fmtJson(200, msg, data))
}

// 返回 404 错误
func (ctx *Context) NotFound() {
	ctx.SetHeader("Content-Type", "application/json", false)
	ctx.w.WriteHeader(404)
	ctx.w.Write(ctx.fmtJson(404, "方法不存在", nil))
}
