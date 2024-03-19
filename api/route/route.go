package route

import (
	"go-clean-architecture-poc/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")
	NewSignUpRoute(env, publicRouter)

	// protectedRouter := gin.Group("")
}
