package main

import (
    "context"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
)

//type helloHandler struct{}
//
//func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//    fmt.Fprintf(w, "Hello World\n")
//}
//
//func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
//    fmt.Fprintf(w, "Welcome!\n")
//}
//
//func main() {
//    router := mux.NewRouter()
//    router.Handle("/", &helloHandler{})
//    router.HandleFunc("/welcome", WelcomeHandler)
//}

func main() {
    muxRouter := mux.NewRouter()
    RegisterRouters(muxRouter)
    server := &http.Server{
        Addr:    ":8080",
        Handler: muxRouter,
    }

    // 创建系统信号接收器
    done := make(chan os.Signal)
    signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-done

        if err := server.Shutdown(context.Background()); err != nil {
            log.Fatal("Shutdown server:", err)
        }
    }()

    log.Println("Starting HTTP server...")
    err := server.ListenAndServe()
    if err != nil {
        if err == http.ErrServerClosed {
            log.Print("Server closed under request")
        } else {
            log.Fatal("Server closed unexpected")
        }
    }
}
