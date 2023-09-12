package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/hello/:name/a", func(c *gee.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at12 %s\n", c.Param("name"), c.Path)
	})
	r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.Run(":9999")
}
