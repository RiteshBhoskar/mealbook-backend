package db

import (
	"fmt"

	"github.com/joho/godotenv"
)

func Connect() {
	db := godotenv.Load("DATABASE_URL")

	if db != nil {
		fmt.Printf("Error loading .env file")
	}
	fmt.Println(db)
}
