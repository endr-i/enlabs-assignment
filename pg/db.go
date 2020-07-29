package pg

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"sync"
)

type Config struct {
	DSN string `default:"host=localhost port=5432 user=postgres dbname=enlabs password=postgresPass"`
}

var (
	db   *gorm.DB
	once sync.Once
)

func Init(config Config, migrations ...Migration) (*gorm.DB, error) {
	var err error
	once.Do(func() {
		if db, err = gorm.Open("postgres", config.DSN); err != nil {
			return
		}
		if _, err = GetMigrationVersion(); err != nil {
			return
		}
		version := migrationVersion.Version
		for i := version; i < len(migrations); i++ {
			m := migrations[i]
			if e := m.MigrateUp(db); e != nil {
				break
			}
			version = i + 1
		}
		if version != migrationVersion.Version {
			migrationVersion.Version = version
			db.Save(&migrationVersion)
		}
	})
	return db, err
}
