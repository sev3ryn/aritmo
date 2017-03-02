package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("../frontend/js/")))
	http.ListenAndServe(":3000", r)
}
