package product

import (
	"database/sql"
	"go-simple-crud/model"
)

type IProductUsecase interface {
	getAllProducts(keyword string) (*[]model.Product, error)
	getProductByID(id string) (*model.Product, error)
	createProduct(product *model.Product) error
	updateProduct(product *model.Product) error
	deleteProduct(id string) error
}
type ProductUsecase struct {
	productRepo IProductRepo
}

func newProductUsecase(db *sql.DB) IProductUsecase {
	return &ProductUsecase{productRepo: newProductRepository(db)}
}
func (p *ProductUsecase) getAllProducts(keyword string) (*[]model.Product, error) {
	products, err := p.productRepo.getAllProducts(keyword)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *ProductUsecase) getProductByID(id string) (*model.Product, error) {
	product, err := p.productRepo.getProductByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductUsecase) createProduct(product *model.Product) error {
	err := p.productRepo.createProduct(product)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductUsecase) updateProduct(product *model.Product) error {
	err := p.productRepo.updateProduct(product)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductUsecase) deleteProduct(id string) error {
	err := p.productRepo.deleteProduct(id)
	if err != nil {
		return err
	}
	return nil
}
