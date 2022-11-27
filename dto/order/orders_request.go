package orderdto

type CreateOrderRequest struct{
	ID int `json:"id"`
	UserID    int   `json:"user_id"`
	ProductID int   `json:"product_id"`
	ToppingID []int `json:"topping_id"`
	Qty       int   `json:"qty"`
	Subtotal  int   `json:"subtotal"`
}

type UpdateOrderRequest struct {
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
	ToppingID int `json:"topping_id"`
}