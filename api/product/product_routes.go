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
	r.HandleFunc("/", controller.GetAllProducts).Methods(http.MethodGet)
	r.HandleFunc("/:id", controller.GetProductByID).Methods(http.MethodGet)
	r.HandleFunc("/", controller.CreateProduct).Methods(http.MethodPost)
	r.HandleFunc("/", controller.UpdateProduct).Methods(http.MethodPut)
	r.HandleFunc("/:id", controller.DeleteProduct).Methods(http.MethodDelete)

}
