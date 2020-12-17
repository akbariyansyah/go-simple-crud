package model

// Product -> Product model
type Product struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
	CreateAt string `json:"created_at"`
	UpdateAt string `json:"update_at"`
}
