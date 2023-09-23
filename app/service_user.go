package app

import (
	"errors"

	"github.com/fabiante/persurl/app/models"
	"github.com/google/uuid"
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

func (s *UserService) GetUser(email string) (*models.User, error) {
	user := &models.User{}

	err := s.db.Take(user, "email = ?", email).Error
	switch {
	case err != nil:
		return nil, mapDBError(err)
	default:
		return user, nil
	}
}

func (s *UserService) GetUserByKey(key string) (*models.User, error) {
	user := &models.User{}

	err := s.db.Model(user).
		Joins("join user_keys on user_keys.owner_id = users.id").
		Take(user).Error

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, ErrNotFound
	case err != nil:
		return nil, mapDBError(err)
	default:
		return user, nil
	}
}

func (s *UserService) CreateUserKey(user *models.User) (*models.UserKey, error) {
	key := &models.UserKey{
		OwnerID: user.ID,
		Value:   uuid.New().String(),
	}

	err := s.db.Create(key).Error
	switch {
	case err != nil:
		return nil, mapDBError(err)
	default:
		return key, nil
	}
}

func (s *UserService) GetUserKey(value string) (*models.UserKey, error) {
	key := &models.UserKey{}

	err := s.db.Take(key, "value = ?", value).Error
	switch {
	case err != nil:
		return nil, mapDBError(err)
	default:
		return key, nil
	}
}
