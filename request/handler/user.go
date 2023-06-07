package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func ShowVisitorInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	country := vars["country"]
	fmt.Fprintf(w, "This guy named %s, was coming from %s .", name, country)
}
