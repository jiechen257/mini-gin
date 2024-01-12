package test

import (
	"log"
	"mini-gin/gin"
	"net/http"
	"text/template"
	"time"
)

func onlyForV2() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func MockRouterGroup() {
	r := gin.New()
	r.Use(gin.Logger()) // global midlleware

	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "css.html", nil)
	})

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
