package transaction

import (
	"assignment/entities/models"
	"github.com/jinzhu/gorm"
	"sync"
)

var (
	repo *repository
	once sync.Once
)

type IRepository interface {
	Transaction(int) (*models.Transaction, error)
	Create(transaction *models.Transaction) error
}

type repository struct {
	db *gorm.DB
}

func (r repository) Transaction(id int) (*models.Transaction, error) {
	transaction := models.Transaction{Id: id}
	err := r.db.First(&transaction).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r repository) Create(transaction *models.Transaction) error {
	return r.db.Create(transaction).Error
}

func InitRepo(db *gorm.DB) IRepository {
	once.Do(func() {
		repo = &repository{db: db}
	})
	return repo
}

func GetRepo() IRepository {
	return repo
}
