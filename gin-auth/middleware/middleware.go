package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	var token string = "1234567"
	return func(context *gin.Context) {
		access_token := context.Request.Header.Get("access_token")
		if token != access_token {
			context.JSON(http.StatusBadRequest, gin.H{
				"msg": "验证失败",
			})
			//立即停止当前请求的执行流程，并阻止后续中间件或路由处理器的进一步处理。当调用  Abort()  函数时，剩余的中间件和路由处理器将不会执行，响应将立即返回给客户端。
			context.Abort()
		} else if token == "" {
			context.JSON(http.StatusBadRequest, gin.H{
				"msg": "验证失败",
			})
		}
		//当调用  Next()  函数时，当前中间件或路由处理器的执行将被暂停，请求将被传递给下一个中间件或路由处理器进行处理。这个函数通常用于在中间件中执行一些操作后，将请求继续传递给下一个处理器，以便实现链式处理请求的目的。
		context.Next()
	}

}
