package repositories

import (
  "waysbucks/models"

  "gorm.io/gorm"
)

type ToppingRepository interface {
  FindToppings() ([]models.Topping, error)
  GetTopping(ID int) (models.Topping, error)
  CreateTopping(topping models.Topping) (models.Topping, error)
}

func RepositoryTopping(db *gorm.DB) *repository {
  return &repository{db}
}

func (r *repository) FindToppings() ([]models.Topping, error) {
  var toppings []models.Topping
  err := r.db.Preload("User").Find(&toppings).Error // add this code

  return toppings, err
}

func (r *repository) GetTopping(ID int) (models.Topping, error) {
  var topping models.Topping
  // not yet using category relation, cause this step doesnt Belong to Many
  err := r.db.Debug().Preload("User").First(&topping, ID).Error

  return topping, err
}

func (r *repository) CreateTopping(topping models.Topping) (models.Topping, error) {
  err := r.db.Create(&topping).Error

  return topping, err
}