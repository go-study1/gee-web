package main

import (
	"net/http"

	"github.com/go-study1/gee-web/gee"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.Html(http.StatusOK, `<h1>index</h1>`)
	})

	r.GET("/hello", func(c *gee.Context) {
		path := c.Path
		c.String(http.StatusOK, `<h1>path =%s</h1>`, path)
	})
	r.POST("/hello", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": "肖",
			"password": "123456",
		})
	})
	r.GET("/:lang/test", func(c *gee.Context) {
		obj := c.Param("lang")
		c.String(http.StatusOK, `<h1>obj =%s</h1>`, obj)
	})
	r.GET("/:lang/get/:id", func(c *gee.Context) {
		obj := c.Param("lang")
		id := c.Param("id")
		c.String(http.StatusOK, `<h1>obj =%s</h1><h1>obj =%s</h1>`, obj, id)

	})
	r.Run(":9999")
}
