package mock

import (
	"fmt"
	"mini-gin/gin"
	"net/http"
	"text/template"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func RenderTemplate() {
	r := gin.New()
	r.Use(gin.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "ginktutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "css.html", nil)
	})
	r.GET("/students", func(c *gin.Context) {
		c.HTML(http.StatusOK, "arr.html", gin.H{
			"title":  "gin",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *gin.Context) {
		c.HTML(http.StatusOK, "custom_func.html", gin.H{
			"title": "gin",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":5500")
}
