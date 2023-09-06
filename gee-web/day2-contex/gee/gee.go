package gee

import (
	"net/http"
)

// Engine 在 Go 语言中，实现了接口方法的 struct 都可以强制转换为接口类型
// Engine 实现了Handler接口
type Engine struct {
	router *Router
}

// HandlerFunc 这是提供给框架用户的，用户写完demo然后传递给服务端
type HandlerFunc func(ctx *Context)

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	c := NewContext(writer, request)
	engine.router.handle(c)
}
func New() *Engine {
	return &Engine{router: newRouter()}
}

// 添加路由表
func (engine *Engine) addRoute(method string, patten string, handler HandlerFunc) {
	engine.router.addRoute(method, patten, handler)
}
func (engine *Engine) Get(patten string, handler HandlerFunc) {
	engine.addRoute("GET", patten, handler)
}
func (engine *Engine) Post(patten string, handler HandlerFunc) {
	engine.addRoute("POST", patten, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
