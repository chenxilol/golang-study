package gee

import (
	"net/http"
	"strings"
)

//	type Router struct {
//		handlers map[string]HandlerFunc
//	}
type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

//	func newRouter() *Router {
//		return &Router{handlers: make(map[string]HandlerFunc)}
//	}
func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

//	func (router Router) addRoute(method string, patten string, handler HandlerFunc) {
//		log.Printf("Router %4s - %s", method, patten)
//		key := method + "-" + patten
//		router.handlers[key] = handler
//
// }
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

//	func (r *Router) handle(c *Context) {
//		key := c.Method + "-" + c.Req.URL.Path
//		// 如果查到了就返回value，查不到就返回false
//		if handler, ok := r.handlers[key]; ok {
//			handler(c)
//		} else {
//			fmt.Fprintf(c.Writer, "404 NOT FOUND: %s \n", c.Req.URL)
//		}
//	}
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}
func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
