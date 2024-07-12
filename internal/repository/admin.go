package repository

import (
	"github.com/DaffaJatmiko/project-iso/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(admin *model.Admin) error
	GetUserByUsername(username string) (*model.Admin, error)
	GetUserByEmail(email string) (*model.Admin, error)
	GetUserByResetToken(resetToken string) (*model.Admin, error)
	UpdateUser(admin *model.Admin) error
	DeleteUser(admin *model.Admin) error
	UpdateResetPasswordToken(admin *model.Admin) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (repo *UserRepositoryImpl) CreateUser(admin *model.Admin) error {
	return repo.db.Create(admin).Error
}

func (repo *UserRepositoryImpl) GetUserByUsername(username string) (*model.Admin, error) {
	var admin model.Admin
	err := repo.db.Where("username = ?", username).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (repo *UserRepositoryImpl) GetUserByEmail(email string) (*model.Admin, error) {
	var admin model.Admin
	err := repo.db.Where("email = ?", email).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (repo *UserRepositoryImpl) GetUserByResetToken(resetToken string) (*model.Admin, error) {
	var admin model.Admin
	err := repo.db.Where("reset_password_token = ?", resetToken).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (repo *UserRepositoryImpl) UpdateUser(admin *model.Admin) error {
	return repo.db.Save(admin).Error
}

func (repo *UserRepositoryImpl) DeleteUser(admin *model.Admin) error {
	return repo.db.Delete(admin).Error
}

func (repo *UserRepositoryImpl) UpdateResetPasswordToken(admin *model.Admin) error {
	return repo.db.Model(&admin).Updates(map[string]interface{}{
		"reset_password_token": admin.ResetPasswordToken,
		"reset_token_expiry":   admin.ResetTokenExpiry,
	}).Error
}
