package main

import (
	"Preparing_Job/internal/api/postgres"
	"fmt"
	"os"
)

func main() {
	db := postgres.ConnectDB()
	defer db.Close()

	_, err := db.Exec("SELECT 1")
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		os.Exit(1)
	}

	fmt.Println("Successfully connected to the database")
}
