package pg

import (
	"github.com/jinzhu/gorm"
	"sync"
)

var (
	migrationVersion = MigrationVersion{Version: -1, Id: 1}
	onceMigration    sync.Once
)

type Migration interface {
	MigrateUp(db *gorm.DB) error
	MigrateDown(db *gorm.DB) error
}

type MigrationVersion struct {
	Id      int `gorm:"PRIMARY_KEY"`
	Version int `gorm:"default:-1"`
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
