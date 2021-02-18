package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strings"
)

type Database struct {
	db *gorm.DB
	username string
	password string
	host string
	port string
	name string
	url string
	entitiesFolder []string
}

func Init() Database {
	err := godotenv.Load()
	var database = Database{
		nil,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_URL"),
		strings.Split(os.Getenv("DB_USER"), ":"),
	}

	if err != nil {
		fmt.Printf("Error while loading env file: %v\n", err)
	}
	database.db, err = gorm.Open(postgres.Open(database.url), &gorm.Config{})
	if err != nil {
		fmt.Printf("Database %s is ready\n", database.name)
	} else {
		fmt.Printf("Failed to initialise database: %v\n", err)
	}
	return database
}