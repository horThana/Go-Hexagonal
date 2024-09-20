package repository

import (
	"github.com/horThana/Backend/core/domain"
	"github.com/horThana/Backend/core/ports"
	"gorm.io/gorm"
)

type GormPayment_Repository struct {
	db *gorm.DB
}

func NewGormPaymentSubscriptionRepository(db *gorm.DB) ports.PayMent_Repository {
	return &GormPayment_Repository{db: db}
}

func (r *GormPayment_Repository) SavePayMentSubscription(payment domain.PayMentSubscription) error {
	if result := r.db.Create(&payment); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormPayment_Repository) GetPayMentSubscriptionByID(id string) (domain.PayMentSubscription, error) {
	var payment domain.PayMentSubscription
	if result := r.db.First(&payment, id); result.Error != nil {
		return domain.PayMentSubscription{}, result.Error
	}
	return payment, nil
}

func (r *GormPayment_Repository) GetAllPayMentSubscription() ([]domain.PayMentSubscription, error) {
	var payments []domain.PayMentSubscription
	if result := r.db.Find(&payments); result.Error != nil {
		return []domain.PayMentSubscription{}, result.Error
	}
	return payments, nil
}

func (r *GormPayment_Repository) DeletePayMentSubscription(id string) error {
	if result := r.db.Delete(&domain.PayMentSubscription{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}