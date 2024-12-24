package models

import "gorm.io/gorm"

type Country struct {
	gorm.Model
	Title  string
	Cities []City
}
