package models

type Product struct {
	ID     int           `json:"id" gorm:"primary_key:auto_increment"`
	Name   string        `json:"name" form:"name" gorm:"type: varchar(255)"`
	Price  int           `json:"price" form:"price" gorm:"type: int"`
	Image  string        `json:"image" form:"image" gorm:"type: varchar(255)"`
	Qty    int           `json:"-" form:"qty"`
	UserID int           `json:"-" form:"user_id"`
	User   UsersResponse `json:"-"`
}

type ProductResponse struct {
	ID     int           `json:"id"`
	Name   string        `json:"name"`
	Price  int           `json:"price"`
	Image  string        `json:"image"`
	Qty    int           `json:"qty"`
	UserID int           `json:"-"`
	User   UsersResponse `json:"-"`
}

type ProductOrderResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
}

func (ProductResponse) TableName() string {
	return "products"
}

// type ProductUserResponse struct {
// 	ID     int    `json:"id"`
// 	Name   string `json:"name"`
// 	Price  int    `json:"price"`
// 	Image  string `json:"image"`
// 	Qty    int    `json:"qty"`
// 	UserID int    `json:"-"`
// }

// func (ProductUserResponse) TableName() string {
// 	return "products"
// }
