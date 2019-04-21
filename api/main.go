package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	// 用户注册
	router.POST("/user", CreateUser)

	// 用户登录
	router.POST("/user/:user_name", Login)

	return router
}

func main() {
	r := RegisterHandlers()
	// 服务启动，监听8000端口
	// listen -> RegisterHandlers -> handlers
	http.ListenAndServe(":8000", r)
}
