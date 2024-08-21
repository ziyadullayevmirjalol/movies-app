package main

import (
	"Bookstore/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	fmt.Println("Server is listening on port :8000")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
