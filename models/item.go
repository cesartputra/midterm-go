package models

import (
	"time"
)

type Item struct {
	ID          uint      `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Name        string    `json:"name" gorm:"type:varchar(255);not null"`
	Description string    `json:"description" gorm:"type:varchar(255);not null"`
	Quantity    int       `json:"quantity" gorm:"type:integer; not null"`
	OrderID     uint      `json:"orderID" gorm:"not null"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
