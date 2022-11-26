package repositories

import (
	"waysbucks/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	FindOrders() ([]models.Order, error)
	// GetOrder(ID int) (models.Order, error)
	// CreateOrder(user models.Order) (models.Order, error)
	// DeleteOrder(user models.Order) (models.Order, error)
}

func RepositoryOrder(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindOrders() ([]models.Order, error) {
	var users []models.Order
	err := r.db.Preload("Products").Preload("Toppings").Find(&users).Error

	return users, err
}
