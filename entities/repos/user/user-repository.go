package user

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
	User(int) (*models.User, error)
	Create(string) (*models.User, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) User(id int) (*models.User, error) {
	user := models.User{Id: id}
	if err := r.db.First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) Create(name string) (*models.User, error) {
	user := models.User{Name: name}
	if err := r.db.Create(&user).Error; err != nil {
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
