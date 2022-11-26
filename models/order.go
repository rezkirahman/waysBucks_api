package models

type Order struct {
	ID         int                    `json:"id" gorm:"primary_key:auto_increment"`
	ProductsID int                    `json:"product_id"`
	Product    ProductOrderResponse   `json:"product"`
	Toppings   []ToppingOrderResponse `json:"toppings" gorm:"many2many:order_toppings"`
	ToppingsID []int `json:"topping_id" form:"topping_id" gorm:"-"`
}

type OrderResponse struct {
	ID         int                    `json:"id"`
	ProductsID int                    `json:"-"`
	Products   ProductOrderResponse   `json:"product"`
	Toppings   []ToppingOrderResponse `json:"toppings"`
}

func (OrderResponse) TableName() string {
	return "orders"
}
