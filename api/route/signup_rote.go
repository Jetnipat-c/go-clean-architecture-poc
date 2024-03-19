package route

import (
	"go-clean-architecture-poc/api/controller"
	"go-clean-architecture-poc/bootstrap"

	"github.com/gin-gonic/gin"
)

func NewSignUpRoute(env *bootstrap.Env, group *gin.RouterGroup) {
	sc := controller.SignUpController{
		Env: env,
	}
	group.POST("/signup", sc.SignUp)
}
