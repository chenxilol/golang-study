package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

const (
	Success = 0
	Fault   = 7
)

func Reust(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func OkWithData(data interface{}, msg string, c *gin.Context) {
	Reust(Success, data, "成功", c)
}
func Faults(data interface{}, c *gin.Context) {
	Reust(http.StatusBadRequest, data, "失败", c)
}
func Fault1(data interface{}, c *gin.Context) {
	Reust(http.StatusBadRequest, data, "失败", c)
}
