package models

import (
	"log"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

// InitDb create a DB connection and run migration
func InitDb() *gorm.DB {

	// init db instance
	db, err := gorm.Open("postgres", "host=localhost port=5433 user=postgres dbname=gormigrate_test password=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(true)

	// run migrations
	m := gormigrate.New(db, gormigrate.DefaultOptions, migrations)

	if err = m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")

	return db
}
