package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func ConnectDB(){
	err := godotenv.Load()
	if (err != nil){
		log.Fatal("Error loading .env file")
	}

	// Data Source Name
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASS"),
        os.Getenv("DB_NAME"),
	)

	DB, err = sqlx.Connect("postgres", dsn)
	if (err != nil){
		log.Fatal("DB Connection Failed", err)
	}

	fmt.Println("Database connected!")
}
