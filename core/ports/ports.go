package ports

import (
	"github.com/horThana/Backend/core/domain"
)

// สำหรับ second port
type ProductRepository interface {
	SaveProduct(product domain.Product) error
	GetProductByID(id string) (domain.Product, error)
	GetAllProduct() ([]domain.Product, error)
	DeleteProduct(id string) error

}


type UsersRepository interface {
	SaveUser(user domain.User) error
	GetUserByID(id string) (domain.User, error)
	GetAllUser() ([]domain.User, error)
	DeleteUser(id string) error
}

type PayMent_Repository interface {
	SavePayMentSubscription(payMentSubscription domain.PayMentSubscription) error
	GetPayMentSubscriptionByID(id string) (domain.PayMentSubscription, error)
	GetAllPayMentSubscription() ([]domain.PayMentSubscription, error)
	DeletePayMentSubscription(id string) error
}
