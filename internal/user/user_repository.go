package user

import (
	"github.com/RenatoHioji/go_elastic_search_repo/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo UserRepository) GetAllUsers() (users []models.User, err error) {
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	
	return users, nil
}
