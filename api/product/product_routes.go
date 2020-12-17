package product

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewProductRoutes(r *mux.Router, db *sql.DB) {

	controller := NewProductController(db)
	r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world")
	})
	r.HandleFunc("/product", controller.GetAllProducts).Methods(http.MethodGet)
	r.HandleFunc("/product/{id}", controller.GetProductByID).Methods(http.MethodGet)
	r.HandleFunc("/product", controller.CreateProduct).Methods(http.MethodPost)
	r.HandleFunc("/product", controller.UpdateProduct).Methods(http.MethodPut)
	r.HandleFunc("/product/{id}", controller.DeleteProduct).Methods(http.MethodDelete)

}
