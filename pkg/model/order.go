package model

import (
	"gorm.io/gorm"
)

type AutoReload struct {
	gorm.Model
	Name   string
	Value  int
	Status string
}
