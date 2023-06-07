package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

func ParseJsonRequestHandler(w http.ResponseWriter, r *http.Request) {
	var p Person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Person: %+v", p)
}
