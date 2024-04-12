package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/project-GoLang/book-store/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	port := "localhost:9010"
	log.Printf("Server is starting and listening on port %s...", port)
	log.Fatal(http.ListenAndServe(port, r))
}
