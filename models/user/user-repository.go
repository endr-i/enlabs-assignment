package user

import (
	"assignment/models/transaction"
	"github.com/jinzhu/gorm"
	"sync"
)

var (
	repo *repository
	once sync.Once
)

type IRepository interface {
	User(int) (*User, error)
}

type repository struct {
	db *gorm.DB
}

func (r repository) User(id int) (*User, error) {
	user := User{Id: id}
	if err := r.db.First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r repository) AddTransaction(id int, tr transaction.Transaction) (*User, error) {
	user := User{Id: id}
	if err := r.db.First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
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
