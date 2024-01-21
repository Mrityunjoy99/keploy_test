package appserver

import (
	"github.com/gin-gonic/gin"
	"github.com/sample_project/src/application"
	"github.com/sample_project/src/application/hello"
	"github.com/sample_project/src/application/user"
	"github.com/sample_project/src/common"
)

func registerApplicationRoutes(g *gin.Engine, s *application.Service, c *common.Config) {
	registerHelloRoutes(g, s.HelloService, c)
	registerUserRoutes(g, s.UserService, c)
}

func registerHelloRoutes(g *gin.Engine, s hello.Service, c *common.Config) {
	helloController := hello.NewController(s)
	g.GET("/hello", helloController.Hello)
}

func registerUserRoutes(g *gin.Engine, s user.Service, c *common.Config) {
	userController := user.NewController(s)
	g.POST("/user", userController.Create)
}
