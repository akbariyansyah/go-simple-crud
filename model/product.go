package model

import (
	"errors"
	"time"
)

// Product -> Product model
type Product struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Price    int       `json:"price"`
	Quantity int       `json:"quantity"`
	CreateAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"update_at"`
}

func (p *Product) Validate() error {
	if p.ID == "" {
		return errors.New("Product ID cannot be empty")
	} else if p.Name == "" {
		return errors.New("Product Name cannot be empty")
	} else if p.Price == 0 {
		return errors.New("Product Price cannot be empty")
	} else if p.Quantity == 0 {
		return errors.New("Product Quantity cannot be empty")
	}
	return nil
}
