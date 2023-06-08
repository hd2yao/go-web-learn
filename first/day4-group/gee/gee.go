package gee

import (
    "log"
    "net/http"
)

type HandlerFunc func(*Context)

type (
    // RouterGroup 需要有访问 Router 的能力
    // 因此保存一个指针，指向 Engine，整个框架的所有资源都是由 Engine 统一协调的，那么就可以通过 Engine 间接地访问各种接口了
    RouterGroup struct {
        prefix      string
        middlewares []HandlerFunc // support middleware
        parent      *RouterGroup  // support nesting
        engine      *Engine       // all groups share a Engine instance
    }

    // Engine 将 Engine 作为最顶层的分组，即 Engine 拥有 RouterGroup 所有的能力
    // 因此，可将原先 Engine 的函数，都交给 RouterGroup 来实现
    Engine struct {
        *RouterGroup
        router *router
        groups []*RouterGroup // store all groups
    }
)

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    c := newContext(w, req)
    engine.router.handle(c)
}

// New is the constructor of gee.Engine
func New() *Engine {
    engine := &Engine{router: newRouter()}
    engine.RouterGroup = &RouterGroup{engine: engine}
    engine.groups = []*RouterGroup{engine.RouterGroup}
    return engine
}

// Group is defined to create a new RouterGroup
// remember all groups share the same Engine instance
func (group *RouterGroup) Group(prefix string) *RouterGroup {
    engine := group.engine
    newGroup := &RouterGroup{
        prefix: group.prefix + prefix,
        parent: group,
        engine: engine,
    }
    engine.groups = append(engine.groups, newGroup)
    return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
    pattern := group.prefix + comp
    log.Printf("Route %4s - %s", method, pattern)
    group.engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
    group.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
    group.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) error {
    return http.ListenAndServe(addr, engine)
}
