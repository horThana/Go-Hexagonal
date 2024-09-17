package domain

type PayMentSubscription struct {
	ID             int     `json:"id"`
	SubscriptionID int     `json:"subscription_id"`
	Amount         float64 `json:"amount"`
	IsPaid         bool    `json:"is_paid"`
	IsCanceled     bool    `json:"is_canceled"`
}

type User struct {
	ID                   int                   `json:"id"`
	Username             string                `json:"username"`
	Password             string                `json:"password"`
	Email                string                `json:"email"`
	Products             []Product             `gorm:"foreignKey:ID" json:"products"`
	PayMentSubscriptions []PayMentSubscription `gorm:"foreignKey:ID" json:"payment_subscriptions"`
}

type Product struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}