package service

import (
	"github.com/DaffaJatmiko/project-iso/internal/model"
	"github.com/DaffaJatmiko/project-iso/internal/repository"
	"github.com/DaffaJatmiko/project-iso/pkg/util"
	"mime/multipart"
	"path/filepath"
	"time"
)

type DocumentService interface {
	CreateDocument(document *model.Document, imageFile *multipart.FileHeader) error
	GetDocuments() ([]*model.Document, error)
	GetDocumentByID(documentID uint) (*model.Document, error)
	UpdateDocument(document *model.Document, imageFile *multipart.FileHeader) error
	DeleteDocument(documentID uint) error
}

type DocumentServiceImpl struct {
	repo repository.DocumentRepository
}

func NewDocumentService(repo repository.DocumentRepository) *DocumentServiceImpl {
	return &DocumentServiceImpl{repo}
}

func (service *DocumentServiceImpl) CreateDocument(document *model.Document, imageFile *multipart.FileHeader) error {
	if imageFile != nil {
		fileName := time.Now().Format("20060102150405") + "_" + imageFile.Filename
		filePath := filepath.Join("uploads", fileName)
		if err := util.SaveFile(imageFile, filePath); err != nil {
			return err
		}
		document.ImagePath = filePath
	}
	return service.repo.CreateDocument(document)
}

func (service *DocumentServiceImpl) GetDocuments() ([]*model.Document, error) {
	return service.repo.GetDocuments()
}

func (service *DocumentServiceImpl) GetDocumentByID(documentID uint) (*model.Document, error) {
	return service.repo.GetDocumentById(documentID)
}

func (service *DocumentServiceImpl) UpdateDocument(document *model.Document, imageFile *multipart.FileHeader) error {
	// Get existing document from the database
	existingDocument, err := service.repo.GetDocumentById(document.ID)
	if err != nil {
		return err
	}

	// Update fields
	existingDocument.FileName = document.FileName
	existingDocument.Link = document.Link
	existingDocument.Description = document.Description

	if imageFile != nil {
		fileName := time.Now().Format("20060102150405") + "_" + imageFile.Filename
		filePath := filepath.Join("uploads", fileName)
		if err := util.SaveFile(imageFile, filePath); err != nil {
			return err
		}
		existingDocument.ImagePath = filePath
	}

	return service.repo.UpdateDocument(existingDocument)
}

func (service *DocumentServiceImpl) DeleteDocument(documentID uint) error {
	return service.repo.DeleteDocument(documentID)
}

//func saveFile(file *multipart.FileHeader, path string) error {
//	src, err := file.Open()
//	if err != nil {
//		return err
//	}
//	defer src.Close()
//
//	dst, err := os.Create(path)
//	if err != nil {
//		return err
//	}
//	defer dst.Close()
//
//	if _, err := io.Copy(dst, src); err != nil {
//		return err
//	}
//	return nil
//}
