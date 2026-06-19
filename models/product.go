package models

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	CategoryID  uint      `json:"category_id" binding:"required"`
	Category    *Category `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"category,omitempty"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name" binding:"required"`
	Price       float64   `gorm:"type:decimal(10,2);not null" json:"price" binding:"required,gt=0"`
	Stock       int       `gorm:"not null;default:0" json:"stock"`
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
