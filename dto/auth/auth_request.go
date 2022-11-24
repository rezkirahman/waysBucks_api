package authdto

type RegisterRequest struct {
	Name     string `gorm:"type: varchar(255)" json:"name"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Role     string `gorm:"type: varchar(255)" json:"role"`
	Password string `gorm:"type: varchar(255)" json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}
