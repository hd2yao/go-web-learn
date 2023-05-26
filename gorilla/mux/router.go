package main

import "github.com/gorilla/mux"

func RegisterRouters(r *mux.Router) {
    indexRouter := r.PathPrefix("/index").Subrouter()
    indexRouter.Handle("/", &helloHandler{})

    userRouter := r.PathPrefix("/user").Subrouter()
    userRouter.HandleFunc("/names/{name}/countries/{country}", ShowVisitorInfo)
}
