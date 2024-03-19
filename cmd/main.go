package main

import (
	"go-clean-architecture-poc/api/route"
	"go-clean-architecture-poc/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	env := app.Env

	db := app.DB
	defer bootstrap.ClosDatabaseConnection(db)

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, db, gin)

	gin.Run(env.ServerAddress)

}
