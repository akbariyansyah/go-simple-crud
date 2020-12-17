package main

import (
	"go-simple-crud/api/product"
	"go-simple-crud/config"
	"log"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	conf := config.StartConfig()
	db := config.NewDatabase(conf)

	product.NewProductRoutes(router, db)

	log.Println("server is starting")

	config.StartServer(router)
}
