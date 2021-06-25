package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const AppName string = "GOLANGTEST"

// global
type Env struct {
	AppPort    string
	AppType    string
	DbUser     string
	DbPassword string
	DbName     string
	DbHost     string
	DbDialect  string
	JwtSecret  string
	JwtExp     string
}

func GetENV() Env {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// get app type
	appType := os.Getenv(AppName + "_APP_TYPE")

	return Env{
		AppType:    appType,
		AppPort:    os.Getenv(AppName + "_" + appType + "_APP_PORT"),
		DbUser:     os.Getenv(AppName + "_" + appType + "_DB_USER"),
		DbPassword: os.Getenv(AppName + "_" + appType + "_DB_PASSWORD"),
		DbName:     os.Getenv(AppName + "_" + appType + "_DB_NAME"),
		DbHost:     os.Getenv(AppName + "_" + appType + "_DB_HOST"),
		DbDialect:  os.Getenv(AppName + "_" + appType + "_DB_DIALECT"),
		JwtSecret:  os.Getenv(AppName + "_" + appType + "_JWT_SECRET"),
		JwtExp:     os.Getenv(AppName + "_" + appType + "_JWT_EXP"),
	}
}
