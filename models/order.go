package models

import (
	"time"
)

type Order struct {
	ID           uint      `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	CustomerName string    `json:"customerName" gorm:"type:varchar(255);not null"`
	OrderedAt    time.Time `json:"orderedAt" gorm:"autoCreateTime"`
	CreatedAt    time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	Items        []Item    `json:"items" gorm:"foreignKey:OrderID"`
}
