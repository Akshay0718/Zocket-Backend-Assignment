package database

import "fmt"

type Product struct {
    ID    int    `json:"id"`
    itemName  string `json:"name"`
    Price int    `json:"price"`
}

type DB struct {
    items []Product
}

func NewDB() *DB {
    return &DB{items: []Product{}}
}

func (database *DB) CreateProduct(item *Product) error {
    item.ID = len(database.items) + 1
    database.items = append(database.items, *item)
    return nil
}

func (database *DB) GetProductByID(id int) (*Product, error) {
    for _, item := range database.items {
        if item.ID == id {
            return &item, nil
        }
    }
    return nil, fmt.Errorf("item not found")
}

func (database *DB) ListProducts(queryFilters ItemFilters) ([]Product, error) {
    // For simplicity, returning all items without filter logic
    return database.items, nil
}

type ItemFilters struct {
    userIdentifier   string
    minPrice string
    maxPrice string
    itemName     string
}
