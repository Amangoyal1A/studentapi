package configs

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type DBConfig struct {
	Host     string `split_words:"true"  json:"DB_HOST"`
	Port     string `split_words:"true"  json:"DB_PORT"`
	Name     string `split_words:"true"  json:"DB_NAME"`
	Username string `split_words:"true"  json:"DB_USERNAME"`
	Password string `split_words:"true"  json:"DB_PASSWORD"`
}

var DB *DBConfig

func loadDBConfig() {
	// Load environment variables from the .env file
	errs := godotenv.Load()
	if errs != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	DB = &DBConfig{}
	err := envconfig.Process("db", DB)
	if err != nil {
		log.Fatal(err.Error())
	}
}
