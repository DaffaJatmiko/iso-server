package model

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	FileName    string `json:"file_name"`
	Link        string `json:"link"`
	Description string `json:"description"`
	ImagePath   string `json:"image_path"`
}
