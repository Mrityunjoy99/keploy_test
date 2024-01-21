package appserver

import (
	"github.com/gin-gonic/gin"
	"github.com/sample_project/cmd/cmdopts"
	"github.com/sample_project/src/application"
	"github.com/sample_project/src/common"
	"github.com/sample_project/src/infrastructure/database"
	"github.com/sample_project/src/infrastructure/repository"
)

func Start(a *cmdopts.Arguments) {
	c, err := common.NewConfig()
	if err != nil {
		panic(err)
	}

	_, err = database.Connect(c)
	if err != nil {
		panic(err)
	}

	r := repository.NewRepository(c)
	s := application.NewService(c, r)
	router := gin.New()

	registerApplicationRoutes(router, &s, c)
	router.Run(":" + c.App.Port)
}
