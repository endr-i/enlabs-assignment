package migrations

import (
	"assignment/entities/models"
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
	return db.AutoMigrate(models.User{}).Error
}

func (userMigration) MigrateDown(db *gorm.DB) error {
	return db.DropTable(&models.User{}).Error
}

type transactionMigration struct {
}

func (transactionMigration) MigrateUp(db *gorm.DB) error {
	if err := db.AutoMigrate(models.Transaction{}).Error; err != nil {
		return err
	}
	return db.Model(&models.Transaction{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT").Error
}

func (transactionMigration) MigrateDown(db *gorm.DB) error {
	if err := db.DropTable(&models.Transaction{}).Error; err != nil {
		return err
	}
	return db.Model(&models.Transaction{}).RemoveForeignKey("user_id", "users(id)").Error
}

type createUsersMigration struct {
}

func (createUsersMigration) MigrateUp(db *gorm.DB) error {
	users := []models.User{
		{Name: "test1"},
		{Name: "test2"},
		{Name: "test3"},
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
