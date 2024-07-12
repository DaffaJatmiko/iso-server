package repository

import (
	"github.com/DaffaJatmiko/project-iso/internal/model"
	"gorm.io/gorm"
)

type DocumentRepository interface {
	CreateDocument(document *model.Document) error
	GetDocuments() ([]*model.Document, error)
	GetDocumentById(documentId uint) (*model.Document, error)
	UpdateDocument(document *model.Document) error
	DeleteDocument(documentId uint) error
}

type DocumentRepositoryImpl struct {
	db *gorm.DB
}

func NewDocumentRepository(db *gorm.DB) DocumentRepository {
	return &DocumentRepositoryImpl{db: db}
}

func (d *DocumentRepositoryImpl) CreateDocument(document *model.Document) error {
	return d.db.Create(document).Error
}

func (d *DocumentRepositoryImpl) GetDocuments() (documents []*model.Document, err error) {
	err = d.db.Find(&documents).Error
	if err != nil {
		return nil, err
	}
	return documents, nil
}

func (d *DocumentRepositoryImpl) GetDocumentById(documentId uint) (*model.Document, error) {
	var document model.Document
	err := d.db.First(&document, "id = ?", documentId).Error
	if err != nil {
		return nil, err
	}
	return &document, nil
}

func (d *DocumentRepositoryImpl) UpdateDocument(document *model.Document) error {
	return d.db.Save(document).Error
}

func (d *DocumentRepositoryImpl) DeleteDocument(documentId uint) error {
	return d.db.Delete(&model.Document{}, "id = ?", documentId).Error
}
