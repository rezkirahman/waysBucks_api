package models

type Topping struct {
	ID     int           `json:"id" gorm:"primary_key:auto_increment"`
	Name   string        `json:"name" form:"name" gorm:"type: varchar(255)"`
	Price  int           `json:"price" form:"price" gorm:"type: int"`
	Image  string        `json:"image" form:"image" gorm:"type: varchar(255)"`
	Qty    int           `json:"qty" form:"qty"`
	UserID int           `json:"-" form:"user_id"`
	User   UsersResponse `json:"-"`
}

type ToppingResponse struct {
	ID     int           `json:"id"`
	Name   string        `json:"name"`
	Price  int           `json:"price"`
	Image  string        `json:"image"`
	Qty    int           `json:"qty"`
	UserID int           `json:"-"`
	User   UsersResponse `json:"-"`
}

type ToppingOrderResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
	UserID int           `json:"-"`
	User   UsersResponse `json:"-"`
}

func (ToppingResponse) TableName() string {
	return "toppings"
}

// type ToppingUserResponse struct {
// 	ID     int    `json:"id"`
// 	Name   string `json:"name"`
// 	Price  int    `json:"price"`
// 	Image  string `json:"image"`
// 	Qty    int    `json:"qty"`
// 	UserID int    `json:"-"`
// }

// func (ToppingUserResponse) TableName() string {
// 	return "toppings"
// }
