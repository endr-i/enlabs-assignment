package pg

import (
	"github.com/jinzhu/gorm"
	"sync"
)

var (
	migrationVersion = MigrationVersion{}
	onceMigration    sync.Once
)

type Migration interface {
	MigrateUp(db *gorm.DB) error
	MigrateDown(db *gorm.DB) error
}

// TODO: handle each migration
type MigrationVersion struct {
	Id      int `gorm:"PRIMARY_KEY"`
	Version int `gorm:"default:0"`
}

func GetMigrationVersion() (*MigrationVersion, error) {
	var err error
	onceMigration.Do(func() {
		db.AutoMigrate(&MigrationVersion{})
		result := db.First(&migrationVersion)
		if result.RecordNotFound() {
			err = db.Create(&migrationVersion).Error
		} else {
			err = result.Error
		}
	})
	return &migrationVersion, err
}
