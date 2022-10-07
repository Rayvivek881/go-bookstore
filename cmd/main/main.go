package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rayvivek881/go-bookstore/pkg/routes"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
