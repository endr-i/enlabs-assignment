package common

import (
	"assignment/entities/models"
	transactionRepo "assignment/entities/repos/transaction"
	userRepo "assignment/entities/repos/user"
	"github.com/jinzhu/gorm"
	"sync"
)

var (
	repo *repository
	once sync.Once
)

type IRepository interface {
	CreateUserTransaction(*models.Transaction) (*models.User, error)
	CancelOddTransactions(int) error
}

type repository struct {
	db                    *gorm.DB
	userRepository        userRepo.IRepository
	transactionRepository transactionRepo.IRepository
}

func (r *repository) CreateUserTransaction(transaction *models.Transaction) (*models.User, error) {
	user := models.User{Id: transaction.UserId}
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&user).Error; err != nil {
			return err
		}
		//if !tx.First(&transactionModel.Transaction{TransactionId: transaction.TransactionId}).RecordNotFound() {
		//	return errors.New("transaction already handled")
		//}
		balance := user.Balance
		if transaction.State == models.TransactionStateWin {
			balance += transaction.Amount
		}
		if transaction.State == models.TransactionStateLose {
			balance -= transaction.Amount
		}
		if balance > 0 {
			transaction.Status = models.TransactionStatusSuccess
			user.Balance = balance
		} else {
			transaction.Status = models.TransactionStatusFailure
		}

		if err := tx.Create(transaction).Error; err != nil {
			return err
		}

		if transaction.Status == models.TransactionStatusSuccess {
			if err := r.db.Save(&user).Error; err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) CancelOddTransactions(n int) error {
	r.db.Transaction(func(tx *gorm.DB) error {

	})
	return nil
}

func InitRepo(db *gorm.DB, userRepository userRepo.IRepository, transactionRepository transactionRepo.IRepository) IRepository {
	once.Do(func() {
		repo = &repository{
			db:                    db,
			userRepository:        userRepository,
			transactionRepository: transactionRepository,
		}
	})
	return repo
}

func GetRepo() IRepository {
	return repo
}
