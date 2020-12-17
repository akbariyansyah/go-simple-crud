package product

import (
	"database/sql"
	"encoding/json"
	"go-simple-crud/model"
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
	products, err := c.productUsecase.getAllProducts("")
	if err != nil {
		utils.SendErrorResponse("Internal server error", http.StatusInternalServerError, w)
		return
	}
	utils.SendSuccessResponse("ok", http.StatusOK, products, w)
}
func (c ProductController) GetProductByID(w http.ResponseWriter, r *http.Request) {

	id := utils.DecodeParams(r, "id")
	product, err := c.productUsecase.getProductByID(id)
	if err != nil {
		utils.SendErrorResponse("Internal server error", http.StatusInternalServerError, w)
		return
	}
	utils.SendSuccessResponse("ok", http.StatusOK, product, w)

}
func (c ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p model.Product

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Println(err)
		utils.SendErrorResponse("Bad request", http.StatusBadRequest, w)
		return
	}
	if err := p.Validate(); err != nil {
		log.Println(err)
		utils.SendErrorResponse(err.Error(), http.StatusBadRequest, w)
		return
	}
	err = c.productUsecase.createProduct(&p)
	if err != nil {
		log.Println(err)
		utils.SendErrorResponse("Internal server error", http.StatusInternalServerError, w)
		return
	}
	utils.SendSuccessResponse("Ok", http.StatusCreated, nil, w)

}
func (c ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var p model.Product

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Println(err)
		utils.SendErrorResponse("Bad request", http.StatusBadRequest, w)
		return
	}
	if err := p.Validate(); err != nil {
		log.Println(err)
		utils.SendErrorResponse(err.Error(), http.StatusBadRequest, w)
		return
	}
	err = c.productUsecase.updateProduct(&p)
	if err != nil {
		log.Println(err)
		utils.SendErrorResponse("Internal server error", http.StatusInternalServerError, w)
		return
	}
	utils.SendSuccessResponse("Ok", http.StatusCreated, nil, w)
}
func (c ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := utils.DecodeParams(r, "id")

	err := c.productUsecase.deleteProduct(id)
	if err != nil {
		log.Println(err)
		utils.SendErrorResponse("Internal server error", http.StatusInternalServerError, w)
		return
	}
	utils.SendSuccessResponse("Ok", http.StatusOK, nil, w)
}
