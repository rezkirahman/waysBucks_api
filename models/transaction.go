package models

type Transaction struct {
	ID      int64        `json:"id" gorm:"primary_key:auto_increment"`
	UserID  int          `json:"user_id" gorm:"type:int"`
	User    UserResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status  string       `json:"status" gorm:"type:varchar(25)"`
	OrderID int          `json:"order_id" gorm:"type:int"`
	Order   []Order      `json:"topping" gorm:"many2many:transaction_order; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Total   int          `json:"total" gorm:"type:int"`
	// Phone   string             `json:"phone" form:"phone" gorm:"type:varchar(255)"`
	// Poscode string             `json:"poscode" form:"poscode" gorm:"type:varchar(255)"`
	// Address string             `json:"address" form:"address" gorm:"type:varchar(255)"`
}
