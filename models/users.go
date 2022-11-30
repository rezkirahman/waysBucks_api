package models

type User struct {
	ID          int         `json:"id" gorm:"primary_key:auto_increment"`
	Name        string      `json:"name" gorm:"type: varchar(255)"`
	Email       string      `json:"email" gorm:"type: varchar(255)"`
	Password    string      `json:"-" gorm:"type: varchar(255)"`
}

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (UserResponse) TableName() string {
	return "users"
}
