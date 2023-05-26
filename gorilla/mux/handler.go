package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World\n")
}

func ShowVisitorInfo(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    name := vars["name"]
    country := vars["country"]
    fmt.Fprintf(w, "This guy named %s, was coming from %s .", name, country)
}
