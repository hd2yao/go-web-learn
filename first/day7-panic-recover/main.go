package main

import (
    "fmt"
    "gee"
    "net/http"
)

func test_recover() {
    defer func() {

        if err := recover(); err != nil {
            fmt.Println("recover success")
        }
        fmt.Println("defer func")
    }()

    arr := []int{1, 2, 3}
    fmt.Println(arr[4])
    fmt.Println("after panic")
}

func main() {
    //test_recover()
    //fmt.Println("after recover")

    r := gee.Default()
    r.GET("/", func(c *gee.Context) {
        c.String(http.StatusOK, "Hello Geektutu\n")
    })
    // index out of range for testing Recovery()
    r.GET("/panic", func(c *gee.Context) {
        names := []string{"geektutu"}
        c.String(http.StatusOK, names[100])
    })

    r.Run(":9999")
}
