package hello

import (
	"github.com/creasty/defaults"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service Service
}

func NewController(service Service) Controller {
	return Controller{service}
}

func (ctrl Controller) Hello(c *gin.Context) {
	dto := HelloDto{}
	if err := c.ShouldBindQuery(&dto); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := defaults.Set(&dto); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	msg := ctrl.service.Hello(dto.Name)

	c.JSON(200, gin.H{
		"message": msg,
	})
}
