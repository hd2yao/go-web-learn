package gee

import "net/http"

type HandlerFunc func(*Context)

func New() *Engine {
    return &Engine{router: newRouter()}
}

type Engine struct {
    router *router
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
    engine.router.addRoute(method, pattern, handler)
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
    engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
    engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(address string) (err error) {
    return http.ListenAndServe(address, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    c := newContext(w, r)
    engine.router.handle(c)
}