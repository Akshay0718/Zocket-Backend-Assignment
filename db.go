package db

import "fmt"

type Product struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Price int    `json:"price"`
}

type DB struct {
    products []Product
}

func NewDB() *DB {
    return &DB{products: []Product{}}
}

func (db *DB) CreateProduct(product *Product) error {
    product.ID = len(db.products) + 1
    db.products = append(db.products, *product)
    return nil
}

func (db *DB) GetProductByID(id int) (*Product, error) {
    for _, product := range db.products {
        if product.ID == id {
            return &product, nil
        }
    }
    return nil, fmt.Errorf("product not found")
}

func (db *DB) ListProducts(filters ProductFilters) ([]Product, error) {
    // For simplicity, returning all products without filter logic
    return db.products, nil
}

type ProductFilters struct {
    UserID   string
    PriceMin string
    PriceMax string
    Name     string
}
