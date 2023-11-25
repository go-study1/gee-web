package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-study1/gee-web/gee"
)

func onlyForV2() gee.HandleFunc {
	return func(c *gee.Context) {
		t := time.Now()
		c.Fail(500, "服务区内部错误")
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gee.New()
	r.LoadHTMLGlob("templates/*")
	r.GET("/index", func(c *gee.Context) {
		c.Html(http.StatusOK, `<h1>index</h1>`)
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/hello", func(c *gee.Context) {
			path := c.Path
			c.String(http.StatusOK, `<h1>path =%s</h1>`, path)
		})
		v1.POST("/hello", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"username": "肖",
				"password": "123456",
			})
		})
	}

	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s \n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}
	stu1 := &student{Name: "Geektutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	static := r.Group("/static")
	{
		static.GET("/students", func(c *gee.Context) {
			c.HtmlTmp(http.StatusOK, "arr.tmpl", gee.H{
				"title":  "学生",
				"stuArr": [2]*student{stu1, stu2},
			})
		})
	}
	static.Static("/", "./static")

	r.Run(":9999")
}

type student struct {
	Name string
	Age  int8
}
