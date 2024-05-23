package main

import (
    "fmt"
    "gee"
    "net/http"
)

func main() {
    router := gee.New()

    router.GET("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
    })
    router.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
        for k, v := range r.Header {
            fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
        }
    })

    router.Run(":9999")
}
