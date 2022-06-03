package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wralith/go-freecodecamp/3-book-management-system/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}

}
