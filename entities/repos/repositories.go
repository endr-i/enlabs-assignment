package repos

import (
	"assignment/entities/repos/common"
	transactionRepo "assignment/entities/repos/transaction"
	userRepo "assignment/entities/repos/user"
	"github.com/jinzhu/gorm"
)

func InitRepositories(db *gorm.DB) {
	userRepository := userRepo.InitRepo(db)
	transactionRepository := transactionRepo.InitRepo(db)
	common.InitRepo(db, userRepository, transactionRepository)
}
