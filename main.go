package main

import (
	"go-simple-crud/api/product"
	"go-simple-crud/config"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	
	conf := config.StartConfig()
	
	db := config.NewDatabase(conf)
	
	product.NewProductRoutes(router, db)

	config.StartServer(router)
}
