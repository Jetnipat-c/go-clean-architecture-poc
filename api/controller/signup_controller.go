package controller

import (
	"go-clean-architecture-poc/bootstrap"

	"github.com/gin-gonic/gin"
)

type SignUpController struct {
	Env *bootstrap.Env
}

func (sc *SignUpController) SignUp(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
