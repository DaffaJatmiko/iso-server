package service

import (
	"github.com/DaffaJatmiko/project-iso/internal/model"
	"github.com/DaffaJatmiko/project-iso/internal/repository"
	"github.com/DaffaJatmiko/project-iso/pkg/util"
	"mime/multipart"
	"path/filepath"
	"time"
)

type GalleryService interface {
	CreateGallery(imageFile *multipart.FileHeader) error
	GetGalleries() ([]model.Gallery, error)
	GetGalleryByID(galleryID uint) (*model.Gallery, error)
	UpdateGallery(galleryID uint, imageFile *multipart.FileHeader) error
	DeleteGallery(galleryID uint) error
}

type galleryService struct {
	repo repository.GalleryRepository
}

func NewGalleryService(repo repository.GalleryRepository) GalleryService {
	return &galleryService{repo}
}

func (s *galleryService) CreateGallery(imageFile *multipart.FileHeader) error {
	fileName := time.Now().Format("20060102150405") + "_" + imageFile.Filename
	filePath := filepath.Join("uploads", fileName)
	if err := util.SaveFile(imageFile, filePath); err != nil {
		return err
	}

	gallery := model.Gallery{
		ImagePath: filePath,
	}
	return s.repo.Create(&gallery)
}

func (s *galleryService) GetGalleries() ([]model.Gallery, error) {
	return s.repo.FindAll()
}

func (s *galleryService) GetGalleryByID(galleryID uint) (*model.Gallery, error) {
	return s.repo.FindByID(galleryID)
}

func (s *galleryService) UpdateGallery(galleryID uint, imageFile *multipart.FileHeader) error {
	gallery, err := s.repo.FindByID(galleryID)
	if err != nil {
		return err
	}

	fileName := time.Now().Format("20060102150405") + "_" + imageFile.Filename
	filePath := filepath.Join("uploads", fileName)
	if err := util.SaveFile(imageFile, filePath); err != nil {
		return err
	}

	gallery.ImagePath = filePath
	return s.repo.Update(gallery)
}

func (s *galleryService) DeleteGallery(galleryID uint) error {
	return s.repo.Delete(galleryID)
}
