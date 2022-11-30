package repositories

import (
	"waysbucks/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	FindOrders() ([]models.Order, error)
	GetOrder(ID int) (models.Order, error)
	CreateOrder(order models.Order) (models.Order, error)
	UpdateOrder(order models.Order) (models.Order, error)
	DeleteOrder(order models.Order) (models.Order, error)
	ProductOrder(ProductID int) (models.Product, error)
	ToppingOrder(ToppingID []int) ([]models.Topping, error)
}

func RepositoryOrder(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindOrders() ([]models.Order, error) {
	var order []models.Order
	err := r.db.Preload("User").Preload("Product").Preload("Topping").Find(&order).Error

	return order, err
}

func (r *repository) GetOrder(ID int) (models.Order, error) {
	var order models.Order
	err := r.db.Preload("User").Preload("Product").Preload("Topping").First(&order, ID).Error

	return order, err
}

func (r *repository) CreateOrder(order models.Order) (models.Order, error) {
	err := r.db.Preload("User").Preload("Product").Preload("Topping").Create(&order).Error

	return order, err
}

func (r *repository) UpdateOrder(order models.Order) (models.Order, error) {
	err := r.db.Preload("User").Preload("Product").Preload("Topping").Save(&order).Error

	return order, err
}

func (r *repository) DeleteOrder(order models.Order) (models.Order, error) {
	err := r.db.Preload("User").Preload("Product").Preload("Topping").Delete(&order).Error

	return order, err
}

func (r *repository) ProductOrder(ProductID int) (models.Product, error) {
	var product models.Product
	err := r.db.First(&product, ProductID).Error

	return product, err
}

func (r *repository) ToppingOrder(ID []int) ([]models.Topping, error){
	var topping []models.Topping
	err := r.db.Find(&topping, ID).Error
	return topping, err
}
