package main

import (
	"gin-redis-cluster/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	c := gin.Default()
	// gin.Recovery() 是
	c.Use(gin.Logger(), gin.Recovery())
	routers.InitRouters(c)

	c.Run(":8080")
}
