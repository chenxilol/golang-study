package routers

import "github.com/gin-gonic/gin"

func InitRouters(c *gin.Engine) {
	api := c.Group("/api")
	initCourse(api)
	initLogin(api)
	initUser(api)
}
