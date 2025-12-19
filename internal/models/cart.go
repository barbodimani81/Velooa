package models

import "time"

type Cart struct {
	ID        uint      `gorm:"column:id;primary_key" json:"id"`
	UserID    *uint     `gorm:"column:user_id" json:"user_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`

	Items []CartItem `gorm:"foreignKey:CartID" json:"items"`
}

type CartItem struct {
	ID               uint      `gorm:"column:id;primary_key" json:"id"`
	CartId           uint      `gorm:"column:cart_id;not null" json:"cart_id"`
	ProductVariantID uint      `gorm:"column:product_variant_id;not null" json:"product_variant_id"`
	Quantity         uint      `gorm:"column:quantity;not null" json:"quantity"`
	CreatedAt        time.Time `gorm:"column:created_at" json:"created_at"`

	Variant ProductVariant `gorm:"foreignKey:ProductVariantID" json:"variant"`
}
