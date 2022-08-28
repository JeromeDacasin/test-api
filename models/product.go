package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `gorm:"varchar(20);not null"`
	Price       int    `gorm:"int(20);not null"`
	Description string `gorm:"varchar(20);not null"`
	Code        string `gorm:"varchar(20);not null"`
}
