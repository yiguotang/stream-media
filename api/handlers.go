package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

// 用户注册
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "create user handler")
}

// 用户登录
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}
