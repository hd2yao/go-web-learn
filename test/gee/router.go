package gee

import "net/http"

type router struct {
	handler map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handler: make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	if _, ok := r.handler[key]; !ok {
		r.handler[key] = handler
	}
}

func (r *router) handle(c *Context) {
	key := c.Req.Method + "-" + c.Req.URL.Path
	if handler, ok := r.handler[key]; ok {
		handler(c)
	} else {
		c.W.WriteHeader(http.StatusNotFound)
	}
}
