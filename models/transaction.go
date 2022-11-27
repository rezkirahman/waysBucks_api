package models

type Transaction struct {
	ID     int             `json:"id" gorm:"primary_key:auto_increment"`
	UserID int             `json:"user_id" gorm:"type:int"`
	User   UserTransaction `json:"user"`
	Status string          `json:"status" gorm:"type:varchar(25)"`
	Order  []string        `json:"order" gorm:"type:varchar(25)"`
	Total  int             `json:"total"`
}
