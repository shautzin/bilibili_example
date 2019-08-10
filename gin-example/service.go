package main

import (
	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	c.JSON(200, &gin.H{
		"message": "ok",
	})
}

func saveHandler(c *gin.Context) {
	key := c.Query("key")
	value := c.Query("value")

	client := GetRedisClient()
	client.Set(key, value, 0)
	c.JSON(200, &gin.H{
		"success": true,
	})
}

func getHandler(c *gin.Context) {
	key := c.Query("key")
	result := GetRedisClient().Get(key)

	if result.Err() == nil {
		c.JSON(200, &gin.H{
			"success": true,
			"name":    result.String(),
		})
	} else {
		c.JSON(500, &gin.H{
			"success": false,
		})
	}
}
