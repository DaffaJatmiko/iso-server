package service

import (
	"github.com/DaffaJatmiko/project-iso/internal/model"
	"github.com/DaffaJatmiko/project-iso/internal/repository"
)

type AuditService interface {
	CreateAudit(audit *model.Audit) error
	CreateKesesuaian(kesesuaian *model.Kesesuaian) error
	CreateAuditWithKesesuaian(audit *model.Audit) (*model.Audit, error)
	GetAuditByID(id uint) (*model.Audit, error)
	GetAllAudits() ([]model.Audit, error)
	UpdateAudit(audit *model.Audit) error
	DeleteAudit(id uint) error
	CalculatePersentaseKesesuaianDokumen(auditID uint) (*model.PersentaseKesesuaianDokumen, error)
	CalculatePersentaseKesesuaianPerKategori(kategori string) (*model.PersentaseKesesuaianPerKategori, error)
	CalculatePersentaseKesesuaianPerPoinAudit(poinAudit string) (*model.PersentaseKesesuaianPerPoinAudit, error)
	CalculatePersentaseKesesuaianPerPoinAuditPerKategori(pointAudit, kategori string) (*model.PersentaseKesesuaianPerPoinAudit, error)
}

type auditService struct {
	repo repository.AuditRepository
}

func NewAuditService(repo repository.AuditRepository) AuditService {
	return &auditService{repo: repo}
}

func (s *auditService) CreateAudit(audit *model.Audit) error {
	return s.repo.CreateAudit(audit)
}

func (s *auditService) CreateKesesuaian(kesesuaian *model.Kesesuaian) error {
	return s.repo.CreateKesesuaian(kesesuaian)
}

func (s *auditService) CreateAuditWithKesesuaian(audit *model.Audit) (*model.Audit, error) {
	if err := s.repo.CreateAudit(audit); err != nil {
		return nil, err
	}

	kesesuaian := &model.Kesesuaian{
		AuditID: audit.ID,
		Audit:   *audit,
	}

	if err := s.repo.CreateKesesuaian(kesesuaian); err != nil {
		return nil, err
	}

	createdAudit, err := s.repo.GetAuditByID(audit.ID)
	if err != nil {
		return nil, err
	}

	return createdAudit, nil
}

func (s *auditService) GetAuditByID(id uint) (*model.Audit, error) {
	return s.repo.GetAuditByID(id)
}

func (s *auditService) GetAllAudits() ([]model.Audit, error) {
	return s.repo.GetAllAudits()
}

func (s *auditService) UpdateAudit(audit *model.Audit) error {
	return s.repo.UpdateAudit(audit)
}

func (s *auditService) DeleteAudit(id uint) error {
	return s.repo.DeleteAudit(id)
}

func (s *auditService) CalculatePersentaseKesesuaianDokumen(auditID uint) (*model.PersentaseKesesuaianDokumen, error) {
	return s.repo.CalculatePersentaseKesesuaianDokumen(auditID)
}

func (s *auditService) CalculatePersentaseKesesuaianPerKategori(kategori string) (*model.PersentaseKesesuaianPerKategori, error) {
	return s.repo.CalculatePersentaseKesesuaianPerKategori(kategori)
}

func (s *auditService) CalculatePersentaseKesesuaianPerPoinAudit(poinAudit string) (*model.PersentaseKesesuaianPerPoinAudit, error) {
	return s.repo.CalculatePersentaseKesesuaianPerPoinAudit(poinAudit)
}

func (s *auditService) CalculatePersentaseKesesuaianPerPoinAuditPerKategori(poinAudit, kategori string) (*model.PersentaseKesesuaianPerPoinAudit, error) {
	return s.repo.CalculatePersentaseKesesuaianPerPoinAuditPerKategori(poinAudit, kategori)
}
