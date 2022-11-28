package models

type Order struct {
	ID        int             `json:"id" gorm:"primary_key:auto_increment"`
	UserID    int             `json:"user_id" gorm:"type:int"`
	User      UserTransaction `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Qty       int             `json:"qty" gorm:"type:int"`
	ProductID int             `json:"-"`
	Product   ProductOrder    `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	//ToppingID []int           `json:"-" gorm:"-"`
	Topping       []Topping `json:"topping" gorm:"many2many:order_topping; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Subtotal      int       `json:"subtotal"`
	//TransactionID int       `json:"transaction_id"`
	// Transaction   Transaction     `json:"-"`
}

type TransactionOrder struct {
	ID        int            `json:"id"`
	ProductID int            `json:"product_id"`
	Product   ProductOrder   `json:"product"`
	ToppingID int            `json:"topping_id"`
	Topping   []ToppingOrder `json:"toppings"`
	Qty       int            `json:"qty"`
	Subtotal  int            `json:"subtotal"`
}

func (TransactionOrder) TableName() string {
	return "orders"
}
