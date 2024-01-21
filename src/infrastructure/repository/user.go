package repository

import (
	"github.com/sample_project/src/domain/entity"
	"github.com/sample_project/src/infrastructure/database"
)

type UserRepository interface {
	CreateUser(user *entity.User) (*entity.User, error)
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) CreateUser(user *entity.User) (*entity.User, error) {
	createdUser := user
	err := database.GetInstance().Create(createdUser).Error

	return createdUser, err
}
