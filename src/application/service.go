package application

import (
	"github.com/sample_project/src/application/hello"
	"github.com/sample_project/src/application/user"
	"github.com/sample_project/src/common"
	"github.com/sample_project/src/infrastructure/repository"
)

type Service struct{
	HelloService hello.Service
	UserService user.Service
}

func NewService(c *common.Config, r *repository.Repository) Service {
	return Service{
		HelloService: hello.NewService(),
		UserService: user.NewService(r.UserRepository),
	}
}