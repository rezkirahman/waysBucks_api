package models

type Order struct {
	ID       int                    `json:"id" gorm:"primary_key:auto_increment"`
	Product  ProductOrderResponse   `json:"product"  gorm:"polymorphic:User"`
	Toppings []ToppingOrderResponse `json:"toppings" gorm:"polymorphic:User"`
}

type OrderResponse struct {
	ID       int                    `json:"id"`
	Products ProductOrderResponse   `json:"product"`
	Toppings []ToppingOrderResponse `json:"toppings"`
}
