package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type Product struct {
	ID          uint            `gorm:"column:id;primary_key" json:"id"`
	CategoryID  uint            `gorm:"column:category_id" json:"category_id"`
	Name        string          `gorm:"column:name;not null" json:"name"`
	Description string          `gorm:"column:description" json:"description"`
	BasePrice   decimal.Decimal `gorm:"column:base_price;type:decimal(10,2);not null" json:"base_price"`
	ImageURL    string          `gorm:"column:image_url;not null" json:"image_url"`
	IsActive    bool            `gorm:"column:is_active;default:true" json:"is_active"`
	CreatedAt   time.Time       `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time       `gorm:"column:updated_at" json:"updated_at"`

	Category Category         `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Variants []ProductVariant `gorm:"foreignKey:ProductID" json:"variants,omitempty"`
}

type ProductVariant struct {
	ID              uint                `gorm:"column:id;primary_key" json:"id"`
	ProductID       uint                `gorm:"column:product_id" json:"product_id"`
	SKU             string              `gorm:"column:sku;unique;not null" json:"sku"`
	Size            string              `gorm:"column:size;not null" json:"size"`
	Color           string              `gorm:"column:color;not null" json:"color"`
	StockQuantity   int                 `gorm:"column:stock_quantity;default:0" json:"stock_quantity"`
	PriceOverride   decimal.NullDecimal `gorm:"column:price_override;type:decimal(10,2)" json:"price_override"`
	DiscountedPrice decimal.NullDecimal `gorm:"column:discounted_price;type:decimal(10,2)" json:"discounted_price"`
	CreatedAt       time.Time           `gorm:"column:created_at" json:"created_at"`
}
