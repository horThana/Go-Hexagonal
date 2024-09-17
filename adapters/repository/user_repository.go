package repository

import (
	"github.com/horThana/Backend/core/domain"
	"github.com/horThana/Backend/core/ports"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) ports.UsersRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) SaveUser(user domain.User) error {
	if result := r.db.Create(&user); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormUserRepository) GetUserByID(id string) (domain.User, error) {
	var user domain.User
	if result := r.db.First(&user, id); result.Error != nil {
		return domain.User{}, result.Error
	}
	return user, nil
}

func (r *GormUserRepository) GetAllUser() ([]domain.User, error) {
	var users []domain.User
	if result := r.db.Find(&users); result.Error != nil {
		return []domain.User{}, result.Error
	}
	return users, nil
}

func (r *GormUserRepository) DeleteUser(id string) error {
	if result := r.db.Delete(&domain.User{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}