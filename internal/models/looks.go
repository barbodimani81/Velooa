package models

import "time"

type Look struct {
	ID          uint      `gorm:"column:id;primary_key" json:"id"`
	Name        string    `gorm:"column:name;not null" json:"name"`
	Description string    `gorm:"column:description" json:"description"`
	ImageURL    string    `gorm:"column:image_url;not null" json:"image_url"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`

	Products []Product `gorm:"many2many:look_products;" json:"products"`
}
