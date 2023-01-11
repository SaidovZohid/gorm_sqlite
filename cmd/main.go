package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/SaidovZohid/gorm_sqlite/api"
	"github.com/SaidovZohid/gorm_sqlite/api/models"
)

func main() {
	db, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.User{})

	apiServer := api.New(&api.RouterOptions{
		GormDB: db,
	})

	if apiServer.Run(":8080"); err != nil {
		log.Fatalf("failed to run server")
	}
}
