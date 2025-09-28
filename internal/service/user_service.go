package service

import (
	"github.com/RenatoHioji/go_elastic_search_repo/internal/models"
	"github.com/RenatoHioji/go_elastic_search_repo/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}
