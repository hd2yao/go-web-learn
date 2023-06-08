package gee

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type H map[string]interface{}

// 这里和 http ServerHTTP 函数原始的两个参数对应 req 和 writer。
// req 是结构体，用指针可以节省内存，Writer 是一个接口类型，不能用指针。

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
    // middleware
    handlers []HandlerFunc
    index    int // index 记录当前执行到第几个中间件
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
    return &Context{
        Path:   req.URL.Path,
        Method: req.Method,
        Req:    req,
        Writer: w,
        index:  -1,
    }
}

func (c *Context) Next() {
    c.index++
    s := len(c.handlers)
    // 在 Next() 中加入循环的作用：
    // 不是所有的 handler 都会调用 Next()
    // 调用 Next()，一般用于在请求前后各实现一些行为
    // 如果中间件只作用于请求前，可以省略调用 Next()，算是一种兼容性比较好的写法
    for ; c.index < s; c.index++ {
        c.handlers[c.index](c)
    }
}

func (c *Context) Fail(code int, err string) {
    c.index = len(c.handlers)
    c.JSON(code, H{"message": err})
}

func (c *Context) Param(key string) string {
    value, _ := c.Params[key]
    return value
}

func (c *Context) PostForm(key string) string {
    return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
    return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
    c.StatusCode = code
    c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
    c.Writer.Header().Set(key, value)
}

func (c *Context) String(code int, format string, values ...interface{}) {
    c.SetHeader("Content-Type", "text/plain")
    c.Status(code)
    c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
    c.SetHeader("Content-Type", "application/json")
    c.Status(code)
    encoder := json.NewEncoder(c.Writer)
    if err := encoder.Encode(obj); err != nil {
        http.Error(c.Writer, err.Error(), 500)
    }
}

func (c *Context) Data(code int, data []byte) {
    c.Status(code)
    c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
    c.SetHeader("Content-Type", "text/html")
    c.Status(code)
    c.Writer.Write([]byte(html))
}
