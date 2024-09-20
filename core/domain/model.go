package domain

type User struct {
	ID                   int                   `json:"id"`
	Name                 string                `json:"name"`
	Username             string                `json:"username"`
	Password             string                `json:"password"`
	Email                string                `json:"email"`
	Products             []Product             `gorm:"foreignKey:ID" json:"products"`
	PayMentSubscriptions []PayMentSubscription `gorm:"foreignKey:UserID" json:"payment_subscriptions"`
}

type Product struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Image string  `json:"image"`
}

type PayMentSubscription struct {
	ID         uint   `json:"id"`
	UserID     uint   `json:"user_id"`
	CardNumber string `json:"card_number"`
	CardName   string `json:"card_name"`
	ExpireDate string `json:"expire_date"`
	Year       int    `json:"year"`
	CVV        int    `json:"cvv"`
}
