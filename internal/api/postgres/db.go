package postgres

import (
	"github.com/go-pg/pg/v10"
	"log"
	"os"
)

func ConnectDB() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     os.Getenv("DB_ADDR"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	})

	_, err := db.Exec("SELECT 1")
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	return db
}
