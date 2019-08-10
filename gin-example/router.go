package main

import (
	"github.com/gin-gonic/gin"
)

// 创建 gin 对象并管理 url 路由
func route() *gin.Engine {

	r := gin.Default()

	r.Use(EnableMemorySession())

	r.GET("/", indexHandler)

	r.GET("/save", saveHandler)
	r.GET("/get", getHandler)

	r.GET("/login", loginHandler)
	r.GET("/currentUser", currentUserHandler)

	return r
}
