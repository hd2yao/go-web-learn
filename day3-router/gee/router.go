package gee

import (
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node // roots[method]
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// Only one * is allowed
// 因为 * 表示后续的全部子路径，因此不管后续是什么，都可以直接匹配成功
/*
	例如：/p/*filepath 可对应
			/p/go1/go2
			/p/go2
			/p/go2/go/other
*/
func parsePattern(pattern string) []string {
	// /hd/3/s*d/*s/anx
	vs := strings.Split(pattern, "/")
	// vs = [ , hd, 3, s*d, *s, anx]
	// vs[0] = "", len(vs) = 6

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" { // 筛去第一个 ""
			parts = append(parts, item)
			if item[0] == '*' { // 只保留第一个且首字符为 * 的 part
				break
			}
		}
	}
	return parts // [hd, 3, s*d, *s]
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok { // 新建 method 分支
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method] // 切换到对应的 method 分支根节点上

	if !ok {
		return nil, nil // 不存在该 method 分支
	}

	n := root.search(searchParts, 0) // 返回查找结果的节点

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index] // 将 /:lang 改为 /go
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/") // 将 /*xx 改为 /xx/xx/...
				// 退出循环
				break
			}
		}
		return n, params
	}

	return nil, nil
}

// 获取 method 分支上的全部路由(有handler)
func (r *router) getRoutes(method string) []*node {
	root, ok := r.roots[method]
	if !ok {
		return nil
	}
	nodes := make([]*node, 0)
	root.travel(&nodes)
	return nodes
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
