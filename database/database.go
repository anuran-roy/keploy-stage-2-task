package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"echo-gorm/database/models"
)

func Connect() (*gorm.DB, error) {
	dsn, found := os.LookupEnv("DB_DSN")
	if !found {
		panic("DB_DSN environment variable not found!")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Joke{})
	// seed(db)

	return db, err
}

// func setup(db *gorm.DB) {
// 	db.AutoMigrate(models.Joke{})
// }
