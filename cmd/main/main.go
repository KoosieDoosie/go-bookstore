package main

import (
	"github.com/KoosieDoosie/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r) // points to routes functionality r = routes
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
