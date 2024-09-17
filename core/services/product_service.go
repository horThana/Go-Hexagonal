package services

import (
	"errors"

	"github.com/horThana/Backend/core/domain"
	"github.com/horThana/Backend/core/ports"
)

type ProductService interface {
    CreateProduct(product domain.Product) error
		FindProductId(id string) (domain.Product, error)
		FindAllProducts() ([]domain.Product, error)
		DeleteProduct(id string) error

}
// ProductServiceImpl is the implementation of ProductService.
type productServiceImpl struct {
    repo ports.ProductRepository
}

// NewProductService creates a new instance of ProductServiceImpl.
func NewProductService(repo ports.ProductRepository) ProductService {
    return &productServiceImpl{repo: repo}
}

// CreateProduct validates and saves a new product.
func (s *productServiceImpl) CreateProduct(product domain.Product) error {
		if product.Name == "" {
			return errors.New("name is required")	
		}
		if err := s.repo.SaveProduct(product); err != nil {
			return err
		}
		return nil
}

func(s *productServiceImpl) FindProductId(id string) (domain.Product, error) {
	product, err := s.repo.GetProductByID(id)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func(s *productServiceImpl) FindAllProducts() ([]domain.Product, error) {
	product, err := s.repo.GetAllProduct()
	if err != nil {
		return []domain.Product{}, err
	}
	return product, nil
}

func(s *productServiceImpl) DeleteProduct(id string) error {
	err := s.repo.DeleteProduct(id)
	if err != nil {
		return err
	}
	return nil
}