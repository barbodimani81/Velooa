package models

import "time"

type User struct {
	ID           uint      `gorm:"column:id;primary_key" json:"id"`
	Email        string    `gorm:"column:email;unique;not null" json:"email"`
	PasswordHash string    `gorm:"column:password_hash;not null" json:"-"`
	Username     string    `gorm:"column:username;not null" json:"username"`
	PhoneNumber  string    `gorm:"column:phone_number;not null" json:"phone_number"`
	BirthDate    time.Time `gorm:"column:birth_date" json:"birth_date"`

	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`

	Addresses []Address `gorm:"foreignKey:UserID" json:"addresses"`
}

type Address struct {
	ID          uint   `gorm:"column:id;primary_key" json:"id"`
	UserID      *uint  `gorm:"column:user_id" json:"user_id"`
	Label       string `gorm:"column:label" json:"label"`
	Line1       string `gorm:"column:line1;not null" json:"line_1"`
	Line2       string `gorm:"column:line2" json:"line_2"`
	City        string `gorm:"column:city" json:"city"`
	PostalCode  string `gorm:"column:postal_code;not null" json:"postal_code"`
	CountryCode string `gorm:"column:country_code;not null" json:"country_code"`
	IsDefault   bool   `gorm:"column:is_default;default:false" json:"is_default"`

	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
