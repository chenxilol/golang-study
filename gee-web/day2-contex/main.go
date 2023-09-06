package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.Post("/g", func(c *gee.Context) {
		c.Json(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})

	})
	r.Run(":9999")
}
