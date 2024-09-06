package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	UrlDatabaseConnection string
	Port                  int
	SecretKey             []byte
)

// LoadEnviroment inicialize enviroment
func LoadEnviroment() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("WSRS_API_PORT"))
	if err != nil {
		Port = 9000
	}

	UrlDatabaseConnection = fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("WSRS_DATABASE_USER"),
		os.Getenv("WSRS_DATABASE_PASSWORD"),
		os.Getenv("WSRS_DATABASE_NAME"),
	)

	SecretKey = []byte(os.Getenv("WSRS_SECRET_AUTH"))
}
