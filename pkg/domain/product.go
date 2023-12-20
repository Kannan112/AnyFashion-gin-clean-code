package domain

import "time"

type Category struct {
	Id         uint   `gorm:"primaryKey;unique;not null"`
	Name       string `gorm:"unique;not null"`
	Created_at time.Time
	Updated_at time.Time
}

type Product struct {
	Id          uint   `gorm:"primaryKey;unique;not null"`
	ProductName string `gorm:"unique;not null"`
	Description string
	Brand       string
	CategoryID  uint
	Category    Category
	Created_at  time.Time
	Updated_at  time.Time
}

type ProductItem struct {
	ID          uint `gorm:"primaryKey;unique;not null" json:"id"`
	ProductID   uint `gorm:"not null" json:"product_id" validate:"required"`
	Product     Product
	SKU         string
	QntyInStock int     `gorm:"not null" json:"qnty_in_stock" validate:"required"`
	Gender      string  `gorm:"not null" json:"gender" validate:"required"`
	Model       string  `gorm:"not null" json:"model" validate:"required"`
	Size        int     `gorm:"not null" json:"size" validate:"required"`
	Color       string  `gorm:"not null" json:"color" validate:"required"`
	Material    string  `gorm:"not null" json:"material" validate:"required"`
	Price       float64 `gorm:"not null" json:"price" validate:"required"`
	Created_at  time.Time
	Updated_at  time.Time
}

type ProductItems struct {
	ID          uint `gorm:"primaryKey;unique;not null" json:"id"`
	ProductID   uint `gorm:"not null" json:"product_id" validate:"required"`
	SKU         string
	QntyInStock int     `gorm:"not null" json:"qnty_in_stock" validate:"required"`
	Gender      string  `gorm:"not null" json:"gender" validate:"required"`
	Model       string  `gorm:"not null" json:"model" validate:"required"`
	Size        int     `gorm:"not null" json:"size" validate:"required"`
	Color       string  `gorm:"not null" json:"color" validate:"required"`
	Material    string  `gorm:"not null" json:"material" validate:"required"`
	Price       float64 `gorm:"not null" json:"price" validate:"required"`
	OfferPrice  float32
	Discout     float32
	ImagesId    uint
	Created_at  time.Time
	Updated_at  time.Time
}

type OfferTable struct {
	ID          uint `gorm:"primaryKey;unique;not null"`
	ProductId   uint `gorm:"not null" json:"product_id" validate:"required"`
	Product     Product
	Discount    float32
	StartDate   string
	EndDate     string
	Discription string
}

type Images struct {
	Id            uint
	ProductItemID uint
	ProductItem   ProductItem
	FileName      string
}
