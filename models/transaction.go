package models

type Transaction struct {
	ID      int     `json:"id" gorm:"primary_key:auto_increment"`
	Name    string  `json:"name" form:"name" gorm:"type:varchar(255)"`
	Email   string  `json:"email" form:"email" gorm:"type:varchar(255)"`
	Phone   string  `json:"phone" form:"phone" gorm:"type:varchar(255)"`
	Poscode string  `json:"poscode" form:"poscode" gorm:"type:varchar(255)"`
	Address string  `json:"address" form:"address" gorm:"type:varchar(255)"`
	Status  string  `json:"status" gorm:"type:varchar(25)"`
	OrderID []int   `json:"-" form:"order_id"`
	Order   []Order `json:"order"`
	Total   int     `json:"total"`
}
