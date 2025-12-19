package models

type Category struct {
	ID       uint   `gorm:"column:id;primary_key"`
	Name     string `gorm:"column:name;not null"`
	Slug     string `gorm:"column:slug;not null;unique"`
	ParentID *uint  `gorm:"column:parent_id"`

	Children []Category `gorm:"foreignKey:ParentID" json:"children,omitempty"`
}
