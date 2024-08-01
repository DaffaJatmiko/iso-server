package model

import "time"
import "gorm.io/gorm"

type Audit struct {
	gorm.Model
	Auditan             string       `gorm:"type:varchar(255)" json:"auditan"`
	Kategori            string       `gorm:"type:varchar(255)" json:"kategori"`
	NoSuratTugas        string       `gorm:"type:varchar(255)" json:"no_surat_tugas"`
	TanggalMulai        time.Time    `json:"tanggal_mulai"`
	TanggalSelesai      time.Time    `json:"tanggal_selesai"`
	TanggalSurat        time.Time    `json:"tanggal_surat"`
	TanggalBAExit       time.Time    `json:"tanggal_ba_exit"`
	TanggalTerbitIHALHA time.Time    `json:"tanggal_terbit_iha_lha"`
	TanggalSelesaiTL    time.Time    `json:"tanggal_selesai_tl"`
	TanggalSuratSelesai time.Time    `json:"tanggal_surat_selesai"`
	TanggalTindakLanjut time.Time    `json:"tanggal_tindak_lanjut"`
	TanggalVerifikasi   time.Time    `json:"tanggal_verifikasi"`
	InspekturHadir      bool         `json:"inspektur_hadir"`
	JumlahOrang         int          `json:"jumlah_orang"`
	Kesesuaians         []Kesesuaian `gorm:"foreignKey:AuditID;constraint:OnDelete:CASCADE;" json:"kesesuaian"` // Add this line to define the relationship
	HariLiburSurat      int          `json:"hari_libur_surat"`
	HariLiburVerifikasi int          `json:"hari_libur_verifikasi"`
	HariLiburIHA        int          `json:"hari_libur_iha"`
	HariLiburBuktiTL    int          `json:"hari_libur_bukti_tl"`
	HariLiburSelesai    int          `json:"hari_libur_selesai"`
}

type Kesesuaian struct {
	gorm.Model
	AuditID                     uint  `json:"audit_id"`
	Audit                       Audit `gorm:"foreignKey:AuditID" json:"-"`
	JumlahHariSurat             int   `json:"jumlah_hari_surat"`
	KesesuaianSurat             bool  `json:"kesesuaian_surat"`
	SkorSurat                   int   `json:"skor_surat"`
	JumlahHariPelaksanaan       int   `json:"jumlah_hari_pelaksanaan"`
	KesesuaianPelaksanaan       bool  `json:"kesesuaian_pelaksanaan"`
	SkorPelaksanaan             int   `json:"skor_pelaksanaan"`
	KesesuaianSDM               bool  `json:"kesesuaian_sdm"`
	SkorSDM                     int   `json:"skor_sdm"`
	JumlahHariVerifikasi        int   `json:"jumlah_hari_verifikasi"`
	KesesuaianVerifikasi        bool  `json:"kesesuaian_verifikasi"`
	SkorVerifikasi              int   `json:"skor_verifikasi"`
	KesesuaianIHA               bool  `json:"kesesuaian_iha"`
	SkorIHA                     int   `json:"skor_iha"`
	KesesuaianBuktiTL           bool  `json:"kesesuaian_bukti_tl"`
	SkorBuktiTL                 int   `json:"skor_bukti_tl"`
	KesesuaianSelesaiAudit      bool  `json:"kesesuaian_selesai_audit"`
	SkorSelesaiAudit            int   `json:"skor_selesai_audit"`
	PersentaseKesesuaianDokumen int   `json:"persentase_kesesuaian_dokumen"`
}

type PersentaseKesesuaianDokumen struct {
	gorm.Model
	AuditID                     uint  `json:"audit_id"`
	Audit                       Audit `gorm:"foreignKey:AuditID" json:"audit"`
	PersentaseKesesuaianDokumen int   `json:"persentase_kesesuaian_dokumen"`
}

type PersentaseKesesuaianPerKategori struct {
	gorm.Model
	Kategori             string `gorm:"type:varchar(255)" json:"kategori"`
	PersentaseKesesuaian int    `json:"persentase_kesesuaian"`
}

type PersentaseKesesuaianPerPoinAudit struct {
	gorm.Model
	PoinAudit            string `gorm:"type:varchar(255)" json:"poin_audit"`
	PersentaseKesesuaian int    `json:"persentase_kesesuaian"`
}

type PersentaseKesesuaianPerPoinAuditPerKategori struct {
	Kategori   string  `json:"kategori"`
	PoinAudit  string  `json:"poin_audit"`
	Persentase float64 `json:"persentase"`
}
