package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rupali59/bookstore-server/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	http.Handle("/", r)
	routes.RegisterBookStoreRoutes(r)
	port := ":8081"
	fmt.Printf("Starting server at port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
