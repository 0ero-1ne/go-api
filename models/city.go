package models

import "gorm.io/gorm"

type City struct {
	gorm.Model
	CountryID uint
	Country   Country
	Title     string
	Users     []User
}
