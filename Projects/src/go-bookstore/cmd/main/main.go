package main

import (
	"log"
	_ "log"
	"net/http"
	_ "net/http"

	"github.com/CrossStack-Q/Go_Backend_2024/tree/main/Projects/src/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8008", r))
}
