package services

import (
	"midterm/models"

	"gorm.io/gorm"
)

func CreateItem(db *gorm.DB, item *models.Item) error {
	return db.Create(&item).Error
}

func UpdateItem(db *gorm.DB, id uint, item *models.Item) error {
	return db.Model(&models.Item{}).Where("id = ?", id).Updates(item).Error
}

func GetItemById(db *gorm.DB, id uint) (*models.Item, error) {
	var item models.Item
	err := db.First(&item, id).Error
	return &item, err
}

func DeleteItemById(db *gorm.DB, id uint) error {
	return db.Delete(&models.Item{}, id).Error
}
