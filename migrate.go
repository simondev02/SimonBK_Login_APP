package main

import (
	"SimonBK_Login_APP/domain/models"
	"SimonBK_Login_APP/infra/db"
	"fmt"
	"log"
)

func Runmigrate() {
	err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	// Migrar las tablas
	err = db.DBConn.AutoMigrate(&models.Login{}) // Primero, la tabla "compania"
	if err != nil {
		log.Fatalf("Failed to migrate table Compania: %v", err)
	}

	fmt.Println("Migration successful")
}
