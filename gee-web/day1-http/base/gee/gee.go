package gee

import (
	"fmt"
	"net/http"
)

// Engine 在 Go 语言中，实现了接口方法的 struct 都可以强制转换为接口类型
// Engine 实现了Handler接口
type Engine struct {
	router map[string]HandlerFunc
}

// HandlerFunc 这是提供给框架用户的，用户写完demo然后传递给服务端
type HandlerFunc func(http.ResponseWriter, *http.Request)

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	key := request.Method + "-" + request.URL.Path
	// 如果查到了就返回value，查不到就返回false
	if handler, ok := engine.router[key]; ok {
		handler(writer, request)
	} else {
		fmt.Fprintf(writer, "404 NOT FOUND: %s \n", request.URL)
	}
}

// New 实例化Engine
func New() *Engine {
	return &Engine{
		router: make(map[string]HandlerFunc),
	}
}

// 添加路由表
func (engine *Engine) addRoute(method string, patten string, handler HandlerFunc) {
	key := method + "-" + patten
	engine.router[key] = handler
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
