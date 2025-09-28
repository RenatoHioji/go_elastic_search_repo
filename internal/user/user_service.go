package user

import (
	"github.com/RenatoHioji/go_elastic_search_repo/internal/models"
)

type UserService struct {
	repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}
