package services

import (
	"errors"

	"github.com/horThana/Backend/core/domain"
	"github.com/horThana/Backend/core/ports"
)

type PaymentService interface {
	CreatePaymentSubscription(payment domain.PayMentSubscription) error
	FindPaymentSubscriptionId(id string) (domain.PayMentSubscription, error)
	FindAllPaymentSubscriptions() ([]domain.PayMentSubscription, error)
	DeletePaymentSubscription(id string) error
}

type paymentServiceImpl struct {
	repo ports.PayMent_Repository
}

func NewPaymentService(repo ports.PayMent_Repository) PaymentService {
	return &paymentServiceImpl{repo: repo}
}

func(s *paymentServiceImpl) CreatePaymentSubscription(payment domain.PayMentSubscription) error {
	if payment.CardNumber == "" {
		return errors.New("card number is required")	
	}
	if err := s.repo.SavePayMentSubscription(payment); err != nil {
		return err
	}
	return nil
}

func(s *paymentServiceImpl) FindPaymentSubscriptionId(id string) (domain.PayMentSubscription, error) {
	payment, err := s.repo.GetPayMentSubscriptionByID(id)
	if err != nil {
		return domain.PayMentSubscription{}, err
	}
	return payment, nil
}

func(s *paymentServiceImpl) FindAllPaymentSubscriptions() ([]domain.PayMentSubscription, error) {
	payment, err := s.repo.GetAllPayMentSubscription()
	if err != nil {
		return []domain.PayMentSubscription{}, err
	}
	return payment, nil
}

func(s *paymentServiceImpl) DeletePaymentSubscription(id string) error {
	err := s.repo.DeletePayMentSubscription(id)
	if err != nil {
		return err
	}
	return nil
}

