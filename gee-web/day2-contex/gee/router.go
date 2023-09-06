package gee

import (
	"fmt"
	"log"
)

type Router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *Router {
	return &Router{handlers: make(map[string]HandlerFunc)}
}
func (router Router) addRoute(method string, patten string, handler HandlerFunc) {
	log.Printf("Router %4s - %s", method, patten)
	key := method + "-" + patten
	router.handlers[key] = handler

}
func (r *Router) handle(c *Context) {
	key := c.Method + "-" + c.Req.URL.Path
	// 如果查到了就返回value，查不到就返回false
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		fmt.Fprintf(c.Writer, "404 NOT FOUND: %s \n", c.Req.URL)
	}
}
