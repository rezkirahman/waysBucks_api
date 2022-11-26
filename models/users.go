package models

type User struct {
	ID       int                   `json:"id" gorm:"primary_key:auto_increment"`
	Name     string                `json:"name" gorm:"type: varchar(255)"`
	Email    string                `json:"email" gorm:"type: varchar(255)"`
	Role     string                `json:"role" gorm:"type: varchar(255)"`
	Password string                `json:"-" gorm:"type: varchar(255)"`
	Products []ProductResponse `json:"-"`
	Toppings []ToppingResponse `json:"-"`
}

type UsersResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

func (UsersResponse) TableName() string {
	return "users"
}
