package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

type Middleware func(handlerFunc http.HandlerFunc) http.HandlerFunc

// 中间件代码模版
func createNewMiddleware() Middleware {
    // 创建一个新的中间件
    middleware := func(next http.HandlerFunc) http.HandlerFunc {
        // 创建一个新的 handler 包裹 next
        handler := func(w http.ResponseWriter, r *http.Request) {
            // 中间件的处理逻辑

            // 调用下一个中间件或者最终的 handler 处理程序
            next(w, r)
        }

        // 返回新建的包装 handler
        return handler
    }

    // 返回新建的中间件
    return middleware
}

// Logging 记录每个 URL 请求的执行时长
func Logging() Middleware {
    // 创建中间件
    return func(handlerFunc http.HandlerFunc) http.HandlerFunc {
        // 创建一个新的 handler 包装 http.HandlerFunc
        return func(w http.ResponseWriter, r *http.Request) {
            // 中间件的处理逻辑
            start := time.Now()
            defer func() {
                log.Println(r.URL.Path, time.Since(start))
            }()

            // 调用下一个中间件或者最终的 handler 处理程序
            handlerFunc(w, r)
        }
    }
}

// Method 验证请求用的是否是指定的 HTTP method，不是则返回 400
func Method(m string) Middleware {
    return func(handlerFunc http.HandlerFunc) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
            if r.Method != m {
                http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
                return
            }
            handlerFunc(w, r)
        }
    }
}

// Chain 按照先后顺序和处理器本身链起来
func Chain(handlerFunc http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
    for _, m := range middlewares {
        handlerFunc = m(handlerFunc)
    }
    return handlerFunc
}

func Hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "hello world")
}

func main() {
    http.HandleFunc("/", Chain(Hello, Method("GET"), Logging()))
    http.ListenAndServe(":8080", nil)
}
