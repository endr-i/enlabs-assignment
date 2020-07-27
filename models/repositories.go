package models

import (
	"assignment/models/transaction"
	"assignment/models/user"
	"github.com/jinzhu/gorm"
)

func InitRepositories(db *gorm.DB) {
	user.InitRepo(db)
	transaction.InitRepo(db)
}
