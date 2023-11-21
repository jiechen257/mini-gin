package main

import (
	"fmt"
	gin "mini-gin/gin"
	"time"
)

func ping(c *gin.Context) {
	fmt.Println("Response successful！", time.Now().Format("2006-01-02 15:04:05"))
	c.String("%s", "pong")
}

func main() {
	r := gin.New()
	r.AddRoute("GET", "/ping", ping)
	r.Run("127.0.0.1:9090")
}
