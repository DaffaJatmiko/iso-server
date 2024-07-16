package repository

import (
	"github.com/DaffaJatmiko/project-iso/internal/model"
	"gorm.io/gorm"
)

type GalleryRepository interface {
	Create(gallery *model.Gallery) error
	FindAll() ([]model.Gallery, error)
	FindByID(id uint) (*model.Gallery, error)
	Update(gallery *model.Gallery) error
	Delete(id uint) error
}

type galleryRepository struct {
	db *gorm.DB
}

func NewGalleryRepository(db *gorm.DB) GalleryRepository {
	return &galleryRepository{db}
}

func (r *galleryRepository) Create(gallery *model.Gallery) error {
	return r.db.Create(gallery).Error
}

func (r *galleryRepository) FindAll() ([]model.Gallery, error) {
	var galleries []model.Gallery
	err := r.db.Find(&galleries).Error
	return galleries, err
}

func (r *galleryRepository) FindByID(id uint) (*model.Gallery, error) {
	var gallery model.Gallery
	err := r.db.First(&gallery, id).Error
	if err != nil {
		return nil, err
	}
	return &gallery, nil
}

func (r *galleryRepository) Update(gallery *model.Gallery) error {
	return r.db.Save(gallery).Error
}

func (r *galleryRepository) Delete(id uint) error {
	return r.db.Delete(&model.Gallery{}, id).Error
}
