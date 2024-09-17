package repository

import (
	"github.com/horThana/Backend/core/domain"
	"github.com/horThana/Backend/core/ports"
	"gorm.io/gorm"
)

type GormProductRepository struct {
    db *gorm.DB
}

func NewGormProductRepository(db *gorm.DB) ports.ProductRepository {
    return &GormProductRepository{db: db}
}

func (r *GormProductRepository) SaveProduct(product domain.Product) error {
    if result := r.db.Create(&product); result.Error != nil {
        return result.Error
    }
    return nil
}

func (r *GormProductRepository) GetProductByID(id string) (domain.Product, error) {
    var product domain.Product
    if result := r.db.First(&product, id); result.Error != nil {
        return domain.Product{}, result.Error
    }
    return product, nil
}

func (r *GormProductRepository) GetAllProduct() ([]domain.Product, error) {
    var products []domain.Product
    if result := r.db.Find(&products); result.Error != nil {
        return []domain.Product{}, result.Error
    }
    return products, nil
}

func (r *GormProductRepository) DeleteProduct(id string) error {
    if result := r.db.Delete(&domain.Product{}, id); result.Error != nil {
        return result.Error
    }
    return nil
}
