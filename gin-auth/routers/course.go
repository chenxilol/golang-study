package routers

import (
	"gin-redis-cluster/course"
	"gin-redis-cluster/middleware"
	"github.com/gin-gonic/gin"
)

func initCourse(c *gin.RouterGroup) {
	v1 := c.Group("/v1")
	v1.Use(middleware.Auth())
	{
		v1.GET("/course/:id", course.Get)
		v1.POST("/course", course.Update)
		v1.PUT("/course")

		c.GET("/", func(context *gin.Context) {
			context.JSON(0, gin.H{
				"test": "teset",
			})
		})
	}
}
