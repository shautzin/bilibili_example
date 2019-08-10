package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func loginHandler(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	// login success
	if password != "" {
		session := sessions.Default(c)
		session.Set("username", username)
		session.Save()

		c.String(200, "登录成功")

	} else {
		c.String(401, "密码错误")
	}
}

func currentUserHandler(c *gin.Context) {
	session := sessions.Default(c)
	var username string
	sessionValue := session.Get("username")
	username = sessionValue.(string)
	c.String(200, username)
}
