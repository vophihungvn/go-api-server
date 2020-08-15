package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

var (
	dbInstance *gorm.DB

	// UserRepository for user table
	UserRepository userRepository
)

// InitDb create a DB connection and run migration
func InitDb() {

	// init db instance
	db, err := gorm.Open("postgres", "host=localhost port=5433 user=postgres dbname=gormigrate_test password=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Init DB ....")
	// Init repository
	UserRepository = userRepository{db}

	db.LogMode(true)

	// run migrations
	m := gormigrate.New(db, gormigrate.DefaultOptions, migrations)

	if err = m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")
}

// CloseDB function
func CloseDB() error {
	fmt.Println("DB closing ...")
	return dbInstance.Close()
}
