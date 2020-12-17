package product

import (
	"database/sql"
	"go-simple-crud/utils"
	"log"
	"net/http"
)

type ProductController struct {
	productUsecase IProductUsecase
}

func NewProductController(db *sql.DB) *ProductController {
	return &ProductController{newProductUsecase(db)}
}
func (c ProductController) GetAllProducts(w http.ResponseWriter, r *http.Request) {

}
func (c ProductController) GetProductByID(w http.ResponseWriter, r *http.Request) {

	id := utils.DecodeParams(r, "id")
	product, err := c.productUsecase.getProductByID(id)
	if err != nil {
		log.Println("Error")
		utils.SendErrorResponse("Internal server error", http.StatusInternalServerError, w)
		return
	}
	utils.SendSuccessResponse("ok", http.StatusOK, product, w)

}
func (c ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	
}
func (c ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {

}
func (c ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {

}
