package main

import (
	"log"
	"mini-gin/gin"
	"net/http"
	"time"
)

func onlyForV2() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.String(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "<h1>Hello gin</h1>")
		})

		v1.GET("/hello", func(c *gin.Context) {
			// expect /hello?name=ginktutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *gin.Context) {
			// expect /hello/ginktutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
