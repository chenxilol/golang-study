package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

//	Context type Context struct {
//		Writer     http.ResponseWriter
//		Req        *http.Request
//		Method     string
//		StatusCode int
//		Path       string
//	}
type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	Params map[string]string
	// response info
	StatusCode int
}

func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Method: req.Method,
		Path:   req.URL.Path,
	}
}
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query 这个查询的是URL后面跟着的参数 // expect /hello?name=chenXi
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status 封装状态码
func (c Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader 封装头部信息
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// 封装向客户端发送字符串内容
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content_Type", "text/plain")
	c.Status(code)
	_, err := c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
	if err != nil {
		return
	}
}
func (c *Context) Json(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	// json.NewEncoder(这里就是要把序列后的代码，写入到哪里)
	encoder := json.NewEncoder(c.Writer)
	// 这里调用Encode是你想序列化的东西
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	_, err := c.Writer.Write(data)
	if err != nil {
		return
	}
}
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content_Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}
