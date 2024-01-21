package user

import (
	"github.com/sample_project/src/domain/entity"
	"github.com/sample_project/src/infrastructure/repository"
)

type Service interface {
	Create(dto CreateUserRequestDto) (*entity.User, error)
}

type service struct {
	repository repository.UserRepository
}

func NewService(repository repository.UserRepository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(dto CreateUserRequestDto) (*entity.User, error) {
	user := entity.User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email: 	 dto.Email,
	}

	createdUser, err := s.repository.CreateUser(&user)

	return createdUser, err
}
