package product

import "gorm.io/gorm"

type Product struct {
}

type ProductStore struct {
	db *gorm.DB
}

func NewProductStore(db *gorm.DB) ProductStore {
	return ProductStore{db: db}
}

// TODO database operations, models definition
