package main

import (
	"fmt"
	gin "mini-gin/gin"
)

func ping(c *gin.Context) {
	fmt.Println(1)
	c.String("%s", "ping")
}

func main() {
	r := gin.New()
	r.AddRoute("GET", "/ping", ping)
	r.Run("127.0.0.1:9090")
}
