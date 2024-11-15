package configs

import (
	"log"

	"github.com/Amangoyal1A/studentapi/models"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"gorm.io/gorm"
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

func AutoMigrate(DB *gorm.DB) {
	if DB == nil {
		log.Fatal("Database connection not initialized. Call LoadDBConfig first.")
	}

	err := DB.AutoMigrate(&models.Student{})
	if err != nil {
		log.Fatalf("Error during migration: %v", err)
	}
	log.Println("Database migration completed")
}