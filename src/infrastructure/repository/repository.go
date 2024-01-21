package repository

import "github.com/sample_project/src/common"

type Repository struct {
	UserRepository UserRepository
}

func NewRepository(c *common.Config) *Repository {
	return &Repository{
		UserRepository: NewUserRepository(),
	}

}
