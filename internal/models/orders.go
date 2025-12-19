package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
	"time"
)

type Order struct {
	ID                 uint            `gorm:"column:id;primary_key"`
	UserID             uint            `gorm:"column:user_id"`
	TotalAmount        decimal.Decimal `gorm:"column:total_amount;type:decimal(10,2);not null" json:"total_amount"`
	TotalDiscountSaved decimal.Decimal `gorm:"column:total_discount_saved;type:decimal(10,2);default:0" json:"total_discount_saved"`
	Status             bool            `gorm:"column:status;not null;default:'PENDING'" json:"status"`
	ShippingAddress    datatypes.JSON  `gorm:"column:shipping_address_json;not null;type:jsonb" json:"shipping_address"`
	PaymentProviderID  string          `gorm:"column:payment_provider_id" json:"payment_provider_id"`
	CreatedAt          time.Time       `gorm:"column:created_at" json:"created_at"`

	Items []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
}

type OrderItem struct {
	ID               uint            `gorm:"column:id;primary_key" json:"id"`
	OrderID          uint            `gorm:"column:order_id;not null" json:"order_id"`
	ProductVariantID uint            `gorm:"column:product_variant_id" json:"product_variant_id"`
	Quantity         int             `gorm:"column:quantity;not null" json:"quantity"`
	PriceAtPurchase  decimal.Decimal `gorm:"column:price_at_purchase;type:decimal(10,2);not null" json:"price_at_purchase"`
}
