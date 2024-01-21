package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Controller struct {
	service Service
}

func NewController(service Service) Controller {
	return Controller{service}
}

func (ctrl Controller) Create(c *gin.Context) {
	dto := CreateUserRequestDto{}
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	validate := validator.New()

	if err := validate.Struct(&dto); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, err := ctrl.service.Create(dto)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})

		return
	}

	responseDto := CreateUserResponseDto{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

	c.JSON(200, responseDto)
}
