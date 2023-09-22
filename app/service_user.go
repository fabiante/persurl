package app

import (
	"github.com/fabiante/persurl/app/models"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(email string) error {
	user := &models.User{
		Email: email,
	}

	err := s.db.Create(user).Error
	switch {
	case err != nil:
		return mapDBError(err)
	default:
		return nil
	}
}
