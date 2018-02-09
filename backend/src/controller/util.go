package controller

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Response struct {
	Data interface{}
	Msg string
	Code int
}

type Context struct {
	w http.ResponseWriter
	r *http.Request
}

func (ctx *Context) SetHeader(key string, val string, unique bool) {
	if unique {
		ctx.w.Header().Set(key, val)
	} else {
		ctx.w.Header().Add(key, val)
	}
}

func (ctx *Context) fmtJson(code int, msg string, data interface{}) []byte {
	res := Response {
		Data: data,
		Msg: msg,
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

func (ctx *Context) Error(msg string) {
	ctx.SetHeader("Content-Type", "application/json", false)
	ctx.w.WriteHeader(400)
	ctx.w.Write(ctx.fmtJson(400, msg, nil))
}

func (ctx *Context) Success(msg string, data interface{}) {
	ctx.SetHeader("Content-Type", "application/json", false)
	ctx.w.WriteHeader(200)
	ctx.w.Write(ctx.fmtJson(200, msg, data))
}

func (ctx *Context) NotFound() {
	ctx.SetHeader("Content-Type", "application/json", false)
	ctx.w.WriteHeader(404)
	ctx.w.Write(ctx.fmtJson(404, "方法不存在", nil))
}