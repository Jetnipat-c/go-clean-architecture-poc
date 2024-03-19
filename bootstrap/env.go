package bootstrap

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	AppEnv                 string
	ServerAddress          string
	ContextTimeout         int
	DBHost                 string
	DBPort                 string
	DBUser                 string
	DBPass                 string
	DBName                 string
	AccessTokenExpiryHour  int
	RefreshTokenExpiryHour int
	AccessTokenSecret      string
	RefreshTokenSecret     string
}

func NewEnv() *Env {
	env := Env{}

	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			panic("Failed to load .env file")
		}
	}

	env.AppEnv = os.Getenv("APP_ENV")
	env.ServerAddress = os.Getenv("SERVER_ADDRESS")
	contextTimeoutStr := os.Getenv("CONTEXT_TIMEOUT")
	env.ContextTimeout, _ = strconv.Atoi(contextTimeoutStr)
	env.DBHost = os.Getenv("DB_HOST")
	env.DBPort = os.Getenv("DB_PORT")
	env.DBUser = os.Getenv("DB_USER")
	env.DBPass = os.Getenv("DB_PASS")
	env.DBName = os.Getenv("DB_NAME")
	accessTokenExpiryHourStr := os.Getenv("ACCESS_TOKEN_EXPIRY_HOUR")
	env.AccessTokenExpiryHour, _ = strconv.Atoi(accessTokenExpiryHourStr)
	refreshTokenExpiryHourStr := os.Getenv("REFRESH_TOKEN_EXPIRY_HOUR")
	env.RefreshTokenExpiryHour, _ = strconv.Atoi(refreshTokenExpiryHourStr)
	env.AccessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")
	env.RefreshTokenSecret = os.Getenv("REFRESH_TOKEN_SECRET")

	return &env
}
