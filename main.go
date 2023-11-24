package main

import (
	"net/http"

	"github.com/go-study1/gee-web/gee"
)

func main() {
	r := gee.New()
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
	r.Run(":9999")
}
