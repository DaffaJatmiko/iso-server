package db

import (
	"log"

	"github.com/DaffaJatmiko/project-iso/internal/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	if !SeedExists(db, "documents") {
		seedDocuments(db)
	}

	if !SeedExists(db, "admins") {
		seedAdmins(db)
	}
}

func SeedExists(db *gorm.DB, tableName string) bool {
	var count int64
	db.Table(tableName).Count(&count)
	return count > 0
}

func seedDocuments(db *gorm.DB) {
	documents := []model.Document{
		{FileName: "ISO 9001:2015 Standard", Link: "https://drive.google.com/drive/folders/1zpdjug_m974pjXhwiH6fNLJ1pHi_J_ko?usp=drive_link", Description: "Standar manajemen mutu internasional.", ImagePath: "uploads/image-seed-1.png"},
		{FileName: "Dokumen Terkendali", Link: "https://drive.google.com/drive/folders/15027dXEyGJcZtZRBc1U0nr1Vb-x38fP8?usp=drive_link", Description: "Dokumen yang terkontrol sesuai standar.", ImagePath: "uploads/image-seed-2.png"},
		{FileName: "Rekaman", Link: "https://drive.google.com/drive/folders/1BL0Vrm61HJFBrHRU5Fx1bGLly6ooru-Z?usp=drive_link", Description: "Rekaman dokumen penting.", ImagePath: "uploads/image-seed-3.png"},
		{FileName: "Audit ISO 9001:2015", Link: "https://drive.google.com/drive/folders/19t8yGBvk-ORQqDMKIF8Ggbmhjkwg-Kgc?usp=drive_link", Description: "Proses audit sesuai standar ISO 9001:2015.", ImagePath: "uploads/image-seed-4.png"},
		{FileName: "Tinjauan Manajemen", Link: "https://drive.google.com/drive/folders/1j4OsrdHG2qNLyLSlR8Ch6b9QB4uPy4SI?usp=drive_link", Description: "Dokumen tinjauan manajemen.", ImagePath: "uploads/image-seed-5.png"},
		{FileName: "Rencana Aksi", Link: "https://drive.google.com/drive/folders/16SIrTjSLH8Fn1wi79WU1uiOiTA574dUf?usp=drive_link", Description: "Dokumen rencana aksi.", ImagePath: "uploads/image-seed-6.png"},
		{FileName: "Pemantauan Sasaran Mutu", Link: "https://drive.google.com/drive/folders/1X1lsDgP-XeMKnKHtuEczEcn6R4RQzRkl?usp=drive_link", Description: "Pemantauan sasaran mutu perusahaan.", ImagePath: "uploads/image-seed-7.png"},
		{FileName: "Mitigasi Risiko", Link: "https://drive.google.com/drive/folders/1fzP1YiPJcAJodDVyKPYGLLQgsyRO46Ij?usp=drive_link", Description: "Strategi mitigasi risiko.", ImagePath: "uploads/image-seed-8.jpg"},
		{FileName: "Dokumen Pendukung Eksternal", Link: "https://drive.google.com/drive/folders/1kN2uJnJHu9SKaHB0R3P4K4COeGOVAAxZ?usp=drive_link", Description: "Dokumen pendukung dari eksternal.", ImagePath: "uploads/image-seed-9.jpg"},
		{FileName: "Pengaduan", Link: "https://drive.google.com/drive/folders/1wtclkS6NX1Z5Kv0fM2PgvHZ8KFkJgWxw?usp=drive_link", Description: "Formulir pengaduan.", ImagePath: "uploads/image-seed-10.jpg"},
		{FileName: "Kompetensi", Link: "https://drive.google.com/drive/folders/1WVrDT9_mzSh9nX7nOO6nkDYf2M2AEG7L?usp=drive_link", Description: "Dokumen kompetensi karyawan.", ImagePath: "uploads/image-seed-11.png"},
	}

	for _, doc := range documents {
		if err := db.Create(&doc).Error; err != nil {
			log.Fatalf("failed to seed documents: %v", err)
		}
	}
}

func seedAdmins(db *gorm.DB) {
	admins := []model.Admin{
		{Username: "adminsetitjenkemendag", Email: "set.itjen2.kemendag@gmail.com", Password: hashPassword("@adminsetitjen123")},
		{Username: "adminadhityakusuma", Email: "m.adhityakusuma8@gmail.com", Password: hashPassword("@adminadhityakusuma123")},
	}

	for _, admin := range admins {
		if err := db.Create(&admin).Error; err != nil {
			log.Fatalf("failed to seed admins: %v", err)
		}
	}
}

func hashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}
