package repository

import (
	"fmt"
	"github.com/DaffaJatmiko/project-iso/internal/model"
	"github.com/DaffaJatmiko/project-iso/pkg/util"
	"gorm.io/gorm"
	"log"
)

type AuditRepository interface {
	CreateAudit(audit *model.Audit) error
	CreateKesesuaian(kesesuaian *model.Kesesuaian) error
	GetAuditByID(id uint) (*model.Audit, error)
	GetAllAudits() ([]model.Audit, error)
	DeleteAudit(id uint) error
	CalculatePersentaseKesesuaianDokumen(auditID uint) (*model.PersentaseKesesuaianDokumen, error)
	CalculatePersentaseKesesuaianPerKategori(kategori string) (*model.PersentaseKesesuaianPerKategori, error)
	CalculatePersentaseKesesuaianPerPoinAudit(poinAudit string) (*model.PersentaseKesesuaianPerPoinAudit, error)
}

type auditRepository struct {
	DB *gorm.DB
}

func NewAuditRepository(db *gorm.DB) AuditRepository {
	return &auditRepository{DB: db}
}

func (r *auditRepository) CreateAudit(audit *model.Audit) error {
	return r.DB.Create(audit).Error
}

func (r *auditRepository) CreateKesesuaian(kesesuaian *model.Kesesuaian) error {
	// Log the received kesesuaian object
	log.Printf("Creating Kesesuaian: %+v", kesesuaian)

	// Ensure the Audit relationship is correctly preloaded
	if err := r.DB.Preload("Kesesuaians").First(&kesesuaian.Audit, kesesuaian.AuditID).Error; err != nil {
		return err
	}

	log.Printf("Loaded Audit for Kesesuaian: %+v", kesesuaian.Audit)

	// Calculate compliance logic here
	kesesuaian.JumlahHariSurat = util.CountBusinessDays(kesesuaian.Audit.TanggalSurat, kesesuaian.Audit.TanggalMulai)
	kesesuaian.KesesuaianSurat = kesesuaian.JumlahHariSurat > 9
	if kesesuaian.KesesuaianSurat {
		kesesuaian.SkorSurat = 1
	} else {
		kesesuaian.SkorSurat = 0
	}

	kesesuaian.JumlahHariPelaksanaan = int(kesesuaian.Audit.TanggalSelesai.Sub(kesesuaian.Audit.TanggalMulai).Hours()/24) + 1
	kesesuaian.KesesuaianPelaksanaan = kesesuaian.JumlahHariPelaksanaan < 8
	if kesesuaian.KesesuaianPelaksanaan {
		kesesuaian.SkorPelaksanaan = 1
	} else {
		kesesuaian.SkorPelaksanaan = 0
	}

	kesesuaian.KesesuaianSDM = kesesuaian.Audit.JumlahOrang <= 5
	if kesesuaian.KesesuaianSDM {
		kesesuaian.SkorSDM = 1
	} else {
		kesesuaian.SkorSDM = 0
	}

	kesesuaian.JumlahHariVerifikasi = util.CountBusinessDays(kesesuaian.Audit.TanggalTindakLanjut, kesesuaian.Audit.TanggalVerifikasi)
	kesesuaian.KesesuaianVerifikasi = kesesuaian.JumlahHariVerifikasi <= 7
	if kesesuaian.KesesuaianVerifikasi {
		kesesuaian.SkorVerifikasi = 1
	} else {
		kesesuaian.SkorVerifikasi = 0
	}

	kesesuaian.KesesuaianIHA = util.CountBusinessDays(kesesuaian.Audit.TanggalBAExit, kesesuaian.Audit.TanggalTerbitIHA) < 11 &&
		util.CountBusinessDays(kesesuaian.Audit.TanggalBAExit, kesesuaian.Audit.TanggalTerbitLHA) < 11
	if kesesuaian.KesesuaianIHA {
		kesesuaian.SkorIHA = 1
	} else {
		kesesuaian.SkorIHA = 0
	}

	kesesuaian.KesesuaianBuktiTL = util.CountBusinessDays(kesesuaian.Audit.TanggalBAExit, kesesuaian.Audit.TanggalSelesaiTL) <= 40
	if kesesuaian.KesesuaianBuktiTL {
		kesesuaian.SkorBuktiTL = 1
	} else {
		kesesuaian.SkorBuktiTL = 0
	}

	kesesuaian.KesesuaianSelesaiAudit = util.CountBusinessDays(kesesuaian.Audit.TanggalSelesaiTL, kesesuaian.Audit.TanggalSuratSelesai) <= 7
	if kesesuaian.KesesuaianSelesaiAudit {
		kesesuaian.SkorSelesaiAudit = 1
	} else {
		kesesuaian.SkorSelesaiAudit = 0
	}

	totalSkor := kesesuaian.SkorSurat + kesesuaian.SkorPelaksanaan + kesesuaian.SkorSDM + kesesuaian.SkorVerifikasi +
		kesesuaian.SkorIHA + kesesuaian.SkorBuktiTL + kesesuaian.SkorSelesaiAudit

	kesesuaian.PersentaseKesesuaianDokumen = (totalSkor * 100) / 7

	return r.DB.Create(kesesuaian).Error
}

func (r *auditRepository) GetAuditByID(id uint) (*model.Audit, error) {
	var audit model.Audit
	if err := r.DB.Preload("Kesesuaians").First(&audit, id).Error; err != nil {
		return nil, err
	}
	return &audit, nil
}

func (r *auditRepository) GetAllAudits() ([]model.Audit, error) {
	var audits []model.Audit
	if err := r.DB.Preload("Kesesuaians").Find(&audits).Error; err != nil {
		return nil, err
	}
	return audits, nil
}

func (r *auditRepository) DeleteAudit(id uint) error {
	// Hapus entitas terkait
	if err := r.DB.Unscoped().Where("audit_id = ?", id).Delete(&model.Kesesuaian{}).Error; err != nil {
		return err
	}

	// Hapus entitas utama
	if err := r.DB.Unscoped().Delete(&model.Audit{}, id).Error; err != nil {
		return err
	}

	// Mengatur ulang sequence ID untuk Audit
	var maxAuditID uint
	r.DB.Model(&model.Audit{}).Select("id").Order("id desc").First(&maxAuditID)
	resetAuditSequenceQuery := fmt.Sprintf("ALTER SEQUENCE audits_id_seq RESTART WITH %d", maxAuditID+1)
	if err := r.DB.Exec(resetAuditSequenceQuery).Error; err != nil {
		return err
	}

	// Mengatur ulang sequence ID untuk Kesesuaian
	var maxKesesuaianID uint
	r.DB.Model(&model.Kesesuaian{}).Select("id").Order("id desc").First(&maxKesesuaianID)

	// Mengambil nama sequence untuk kesesuaian
	var kesesuaianSequenceName string
	r.DB.Raw("SELECT c.relname FROM pg_class c WHERE c.relkind = 'S' AND c.relname LIKE '%kesesuaian%'").Row().Scan(&kesesuaianSequenceName)

	resetKesesuaianSequenceQuery := fmt.Sprintf("ALTER SEQUENCE %s RESTART WITH %d", kesesuaianSequenceName, maxKesesuaianID+1)
	if err := r.DB.Exec(resetKesesuaianSequenceQuery).Error; err != nil {
		return err
	}
	return nil
}

func (r *auditRepository) CalculatePersentaseKesesuaianDokumen(auditID uint) (*model.PersentaseKesesuaianDokumen, error) {
	var kesesuaian model.Kesesuaian
	if err := r.DB.Preload("Audit").Where("audit_id = ?", auditID).First(&kesesuaian).Error; err != nil {
		return nil, err
	}

	totalSkor := kesesuaian.SkorSurat + kesesuaian.SkorPelaksanaan + kesesuaian.SkorSDM + kesesuaian.SkorVerifikasi +
		kesesuaian.SkorIHA + kesesuaian.SkorBuktiTL + kesesuaian.SkorSelesaiAudit

	persentase := (totalSkor / 7) * 100

	persentaseKesesuaianDokumen := &model.PersentaseKesesuaianDokumen{
		AuditID:                     auditID,
		PersentaseKesesuaianDokumen: persentase,
	}

	return persentaseKesesuaianDokumen, nil
}

func (r *auditRepository) CalculatePersentaseKesesuaianPerKategori(kategori string) (*model.PersentaseKesesuaianPerKategori, error) {
	var totalAudits int64
	if err := r.DB.Model(&model.Audit{}).Where("kategori = ?", kategori).Count(&totalAudits).Error; err != nil {
		return nil, err
	}

	if totalAudits == 0 {
		return nil, nil
	}

	var totalPersentase int
	rows, err := r.DB.Table("kesesuaians").
		Select("persentase_kesesuaian_dokumen").
		Joins("JOIN audits ON audits.id = kesesuaians.audit_id").
		Where("audits.kategori = ?", kategori).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var persentaseKesesuaianDokumen int
		if err := rows.Scan(&persentaseKesesuaianDokumen); err != nil {
			return nil, err
		}

		totalPersentase += persentaseKesesuaianDokumen
	}

	persentase := totalPersentase / int(totalAudits)

	persentaseKesesuaianPerKategori := &model.PersentaseKesesuaianPerKategori{
		Kategori:             kategori,
		PersentaseKesesuaian: persentase,
	}

	return persentaseKesesuaianPerKategori, nil
}

func (r *auditRepository) CalculatePersentaseKesesuaianPerPoinAudit(poinAudit string) (*model.PersentaseKesesuaianPerPoinAudit, error) {
	var totalAudits, totalSkor int64
	if err := r.DB.Model(&model.Audit{}).Count(&totalAudits).Error; err != nil {
		return nil, err
	}

	log.Printf("Total Audits: %d", totalAudits)

	switch poinAudit {
	case "Surat":
		if err := r.DB.Model(&model.Kesesuaian{}).Where("kesesuaian_surat = ?", true).Count(&totalSkor).Error; err != nil {
			return nil, err
		}
	case "Pelaksanaan":
		if err := r.DB.Model(&model.Kesesuaian{}).Where("kesesuaian_pelaksanaan = ?", true).Count(&totalSkor).Error; err != nil {
			return nil, err
		}
	case "SDM":
		if err := r.DB.Model(&model.Kesesuaian{}).Where("kesesuaian_sdm = ?", true).Count(&totalSkor).Error; err != nil {
			return nil, err
		}
	case "Verifikasi":
		if err := r.DB.Model(&model.Kesesuaian{}).Where("kesesuaian_verifikasi = ?", true).Count(&totalSkor).Error; err != nil {
			return nil, err
		}
	case "IHA":
		if err := r.DB.Model(&model.Kesesuaian{}).Where("kesesuaian_iha = ?", true).Count(&totalSkor).Error; err != nil {
			return nil, err
		}
	case "Bukti TL":
		if err := r.DB.Model(&model.Kesesuaian{}).Where("kesesuaian_bukti_tl = ?", true).Count(&totalSkor).Error; err != nil {
			return nil, err
		}
	case "Selesai Audit":
		if err := r.DB.Model(&model.Kesesuaian{}).Where("kesesuaian_selesai_audit = ?", true).Count(&totalSkor).Error; err != nil {
			return nil, err
		}
	}

	log.Printf("Total Skor untuk %s: %d", poinAudit, totalSkor)

	if totalAudits == 0 {
		return &model.PersentaseKesesuaianPerPoinAudit{
			PoinAudit:            poinAudit,
			PersentaseKesesuaian: 0,
		}, nil
	}

	persentase := (totalSkor) * 100 / (totalAudits)
	log.Printf("Persentase Kesesuaian untuk %s: %f", poinAudit, persentase)

	persentaseKesesuaianPerPoinAudit := &model.PersentaseKesesuaianPerPoinAudit{
		PoinAudit:            poinAudit,
		PersentaseKesesuaian: int(persentase),
	}

	return persentaseKesesuaianPerPoinAudit, nil
}

func boolToInt(value bool) int {
	if value {
		return 1
	}
	return 0
}
