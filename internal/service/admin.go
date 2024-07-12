package service

import (
	"errors"
	"github.com/DaffaJatmiko/project-iso/internal/model"
	"github.com/DaffaJatmiko/project-iso/internal/repository"
	"github.com/DaffaJatmiko/project-iso/pkg/util"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserService interface {
	Register(admin *model.Admin) error
	Login(admin *model.Admin) (string, error)
	Logout(token string) error
	GetUserByUsername(username string) (*model.Admin, error)
	UpdateUser(admin *model.Admin) error
	DeleteUser(admin *model.Admin) error
	RequestPasswordReset(email string) error
	ResetPassword(token, newPassword string) error
}

type UserServiceImpl struct {
	repo      repository.UserRepository
	jwtSecret string
	//redisClient *redis.Client
	smtpConfig util.SMTPConfig
}

func NewUserService(repo repository.UserRepository, jwtSecret string, smtpConfig util.SMTPConfig) *UserServiceImpl {
	return &UserServiceImpl{
		repo:      repo,
		jwtSecret: jwtSecret,
		//redisClient: redisClient,
		smtpConfig: smtpConfig,
	}
}

func (u *UserServiceImpl) Register(admin *model.Admin) error {
	_, err := u.repo.GetUserByUsername(admin.Username)
	if err == nil {
		return errors.New("User already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := &model.Admin{
		Username: admin.Username,
		Email:    admin.Email,
		Password: string(hashedPassword),
	}

	err = u.repo.CreateUser(newUser)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserServiceImpl) Login(admin *model.Admin) (string, error) {
	existingUser, err := u.repo.GetUserByUsername(admin.Username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(admin.Password))
	if err != nil {
		return "", err
	}

	token, err := util.GenerateJWT(admin.Username, u.jwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UserServiceImpl) Logout(token string) error {
	return nil
}

func (u *UserServiceImpl) GetUserByUsername(username string) (*model.Admin, error) {
	user, err := u.repo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (u *UserServiceImpl) CreateUser(admin *model.Admin) error {
	err := u.repo.CreateUser(admin)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserServiceImpl) UpdateUser(admin *model.Admin) error {
	err := u.repo.UpdateUser(admin)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserServiceImpl) DeleteUser(admin *model.Admin) error {
	err := u.repo.DeleteUser(admin)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserServiceImpl) RequestPasswordReset(email string) error {
	admin, err := u.repo.GetUserByEmail(email)
	if err != nil {
		return err
	}

	token, err := util.GenerateResetToken()
	if err != nil {
		return err
	}

	admin.ResetPasswordToken = token
	admin.ResetTokenExpiry = time.Now().Add(1 * time.Hour).Unix()

	err = u.repo.UpdateResetPasswordToken(admin)
	if err != nil {
		return err
	}

	return util.SendResetEmail(u.smtpConfig, email, token)
}

func (u *UserServiceImpl) ResetPassword(token, newPassword string) error {
	admin, err := u.repo.GetUserByResetToken(token)
	if err != nil {
		return err
	}

	if time.Now().Unix() > admin.ResetTokenExpiry {
		return errors.New("Reset Token Expired")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	admin.Password = string(hashedPassword)
	admin.ResetPasswordToken = ""
	admin.ResetTokenExpiry = 0

	return u.repo.UpdateUser(admin)
}
