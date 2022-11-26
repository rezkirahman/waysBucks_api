package productsdto

type CreateProductRequest struct {
  Name       string `json:"name" form:"name" gorm:"type: varchar(255)"`
  Price      int    `json:"price" form:"price" gorm:"type: int"`
  Image      string `json:"image" form:"image" gorm:"type: varchar(255)"`
  Qty        int    `json:"qty" form:"qty" gorm:"type: int"`
  UserID     int    `json:"user_id" gorm:"type: int"`
}

type UpdateProductRequest struct {
  Name       string `json:"name" form:"name" gorm:"type: varchar(255)"`
  Price      int    `json:"price" form:"price" gorm:"type: int"`
  Image      string `json:"image" form:"image" gorm:"type: varchar(255)"`
  Qty        int    `json:"qty" form:"qty" gorm:"type: int"`
}
