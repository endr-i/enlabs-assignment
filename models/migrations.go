package models

import (
	"assignment/models/transaction"
	"assignment/models/user"
	"assignment/pg"
	"github.com/jinzhu/gorm"
)

var Migrations = []pg.Migration{
	userMigration{},
	transactionMigration{},
	createUsersMigration{}, // creates 3 test users TODO: delete in prod
}

type userMigration struct {
}

func (userMigration) MigrateUp(db *gorm.DB) error {
	return db.AutoMigrate(user.User{}).Error
}

func (userMigration) MigrateDown(db *gorm.DB) error {
	return db.DropTable(&user.User{}).Error
}

type transactionMigration struct {
}

func (transactionMigration) MigrateUp(db *gorm.DB) error {
	return db.AutoMigrate(transaction.Transaction{}).Error
}

func (transactionMigration) MigrateDown(db *gorm.DB) error {
	return db.DropTable(&transaction.Transaction{}).Error
}

type createUsersMigration struct {
}

func (createUsersMigration) MigrateUp(db *gorm.DB) error {
	users := []user.User{
		{}, {}, {},
	}
	var err error
	for _, u := range users {
		if err = db.Create(&u).Error; err != nil {
			break
		}
	}
	return err
}

func (createUsersMigration) MigrateDown(db *gorm.DB) error {
	return nil
}
