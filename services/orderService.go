package services

import (
	"fmt"
	"midterm/models"

	"gorm.io/gorm"
)

func CreateOrder(db *gorm.DB, order *models.Order) error {
	return db.Create(&order).Error
}

func GetOrderstWithItems(db *gorm.DB) ([]models.Order, error) {
	var orders []models.Order

	if err := db.Preload("Items").Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func GetOrderByIdWithItems(db *gorm.DB, id uint) (*models.Order, error) {
	var order models.Order
	err := db.Preload("Items").First(&order, id).Error
	fmt.Println(&order)
	return &order, err
}

func UpdateOrder(db *gorm.DB, id uint, order *models.Order) error {
	return db.Model(&models.Order{}).Where("id = ?", id).Updates(order).Error
}

func GetOrderById(db *gorm.DB, id uint) (*models.Order, error) {
	var order models.Order
	err := db.First(&order, id).Error
	return &order, err
}

func DeleteOrderById(db *gorm.DB, id uint) error {
	return db.Delete(&models.Order{}, id).Error
}
