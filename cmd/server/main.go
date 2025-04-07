package main

import (
	"log"

	"github.com/joho/godotenv"
	mysqldb "github.com/michaelrodriguess/user-service/pkg/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  .env not found, using default environment variables")
	}

	mysqldb.Init()
}
