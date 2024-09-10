package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/singhpranshu/streak-assignment/handlers"
)

func main() {
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/find-pairs", services.FindPairHandler).Methods(http.MethodPost)

    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8000", r))
}