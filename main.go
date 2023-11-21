package main

import (
	gin "mini-gin/gin"
)

func ping(c *gin.Context) {
	c.String("%s", "pong")
}

func main() {
	r := gin.New()
	r.AddRoute("GET", "/ping", ping)
	r.Run("127.0.0.1:9090")
}
