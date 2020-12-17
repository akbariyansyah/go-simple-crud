package product

import (
	"database/sql"
	"go-simple-crud/model"
	"log"
)

type IProductRepo interface {
	getAllProducts(keyword string) (*[]model.Product, error)
	getProductByID(id string) (*model.Product, error)
	createProduct(product *model.Product) error
	updateProduct(product *model.Product) error
	deleteProduct(id string) error
}

type Repository struct {
	db *sql.DB
}

func (r Repository) getAllProducts(keyword string) (*[]model.Product, error) {
	var products []model.Product

	rows, err := r.db.Query("select * from m_product")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for rows.Next() {
		product := new(model.Product)
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity, &product.CreateAt, &product.UpdateAt)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		products = append(products, *product)
	}
	return &products, nil
}

func (r Repository) getProductByID(id string) (*model.Product, error) {
	p := new(model.Product)
	tx, err := r.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	stmt, err := tx.Prepare("select * from m_product where id = ?")
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return nil, err
	}
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price, &p.Quantity, &p.CreateAt, &p.UpdateAt)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return nil, err
	}
	tx.Commit()
	return p, nil
}

func (r Repository) createProduct(product *model.Product) error {
	tx, err := r.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	stmt, err := tx.Prepare("insert into m_product(id,name,price,quantity,created_at,updated_at) values(?,?,?,?,current_timestamp,current_timestamp)")
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}
	_, err = stmt.Exec(product.ID, product.Name, product.Price, product.Quantity)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}
	tx.Commit()
	return nil
}

func (r Repository) updateProduct(product *model.Product) error {
	tx, err := r.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	stmt, err := tx.Prepare("update m_product set name=?,price=?,quantity=?,updated_at=current_timestamp where id =?")
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}
	_, err = stmt.Exec(product.Name, product.Price, product.Quantity, product.ID)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}
	tx.Commit()
	return nil
}

func (r Repository) deleteProduct(id string) error {
	tx, err := r.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	stmt, err := tx.Prepare("delete from m_product where id = ?")
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}
	tx.Commit()
	return nil
}

func (r Repository) searchProduct(keyword string) (*model.Product, error) {
	panic("implement me")
}

func newProductRepository(db *sql.DB) IProductRepo {
	return &Repository{db: db}
}
