package model

import (
	"gorm.io/gorm"
)

type Gallery struct {
	gorm.Model
	ImagePath string `json:"image_path"`
}
