package main

import (
	"go-simple-crud/api/product"
	"go-simple-crud/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	conf := config.StartConfig()
	db := config.NewDatabase(conf)

	product.NewProductRoutes(router, db)

	log.Println("server is starting")

	log.Fatal(http.ListenAndServe(":9090", router))
}
