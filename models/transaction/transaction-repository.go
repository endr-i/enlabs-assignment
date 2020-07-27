package transaction

import (
	"github.com/jinzhu/gorm"
	"sync"
)

var (
	repo *repository
	once sync.Once
)

type IRepository interface {
	Transaction(int) (*Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func (r repository) Transaction(id int) (*Transaction, error) {
	transaction := Transaction{Id: id}
	err := r.db.First(&transaction).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
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
