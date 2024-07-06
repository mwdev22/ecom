package product

import (
	"fmt"

	"github.com/mwdev22/ecom/app/types"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          int     `gorm:"primaryKey;autoIncrement"`
	Name        string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	Image       string  `gorm:"not null"`
	Price       float64 `gorm:"not null"`
	Quantity    int     `gorm:"not null"`
}

type ProductStore struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) ProductStore {
	return ProductStore{db: db}
}

func (s *ProductStore) GetProducts() ([]Product, error) {
	var products []Product
	result := s.db.Find(&products)
	if result.Error != nil {
		return nil, fmt.Errorf("error getting products: %s", result.Error)
	}
	return products, nil
}

func (s *ProductStore) CreateProduct(payload *types.AddProductPayload) error {
	existingProduct, err := s.GetProductByName(payload.Name)
	if err == nil {
		return fmt.Errorf("product with name %s already exists", existingProduct.Name)
	}
	newProduct := Product{
		Name:        payload.Name,
		Image:       payload.Image,
		Description: payload.Description,
		Price:       payload.Price,
		Quantity:    payload.Quantity,
	}
	err = s.db.Create(&newProduct).Error
	if err != nil {
		return fmt.Errorf("couldnt create user: %s", err)
	}
	return nil
}

func (s *ProductStore) GetProductByName(name string) (*Product, error) {
	var product Product
	err := s.db.Where(&Product{Name: name}).First(&product)
	if err != nil {
		return nil, fmt.Errorf("product with name: %s not found", name)
	}
	return &product, nil
}

func (s *ProductStore) GetProductById(id int) (*Product, error) {
	var product Product
	err := s.db.Where(&Product{ID: id}).First(&product)
	if err != nil {
		return nil, fmt.Errorf("product with id: %v not found", id)
	}
	return &product, nil
}

func (s *ProductStore) UpdateProduct(product *Product) error {
	if err := s.db.Save(product).Error; err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}
	return nil
}

func (s *ProductStore) DeleteProduct(id int) error {
	err := s.db.Delete(&Product{}, id).Error
	if err != nil {
		return fmt.Errorf("couldnt delete the product: %s", err)
	}

	return nil
}
