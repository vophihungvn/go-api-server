package models

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

var createTableUser = gormigrate.Migration{
	ID: "20200815211700",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(&User{}).Error
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.DropTable("people").Error
	},
}

var migrations = []*gormigrate.Migration{
	&createTableUser,
}
