package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

// gin session key
const KEY = "Your Secret Key"

// 使用 Cookie 保存 session
func EnableCookieSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(KEY))
	return sessions.Sessions("SESSIONID", store)
}

// 使用 Redis 保存 session
func EnableRedisSession() gin.HandlerFunc {
	store, _ := redis.NewStore(10, "tcp", "redis.yy:6379", "87660543", []byte(KEY))
	return sessions.Sessions("SESSIONID", store)
}

// 使用 内存 保存 session
func EnableMemorySession() gin.HandlerFunc {
	store := memstore.NewStore([]byte(KEY))
	return sessions.Sessions("SESSIONID", store)
}
