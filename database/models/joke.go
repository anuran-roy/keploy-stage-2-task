package models

import (
	"gorm.io/gorm"
)

type Joke struct {
	gorm.Model
	Id       uint   // `json:"id"`
	Joke     string // `json:"joke"`
	Category string // `json:"category"`
}
