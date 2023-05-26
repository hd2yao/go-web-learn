package gee

import "net/http"

type HandlerFunc func(*Context)

type Context struct {
	W   http.ResponseWriter
	Req *http.Request
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		W:   w,
		Req: req,
	}
}
