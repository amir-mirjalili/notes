package service

import (
	"github.com/amir-mirjalili/notes.git/users/repository"
)

type Repository interface {
	List() ([]repository.User, error)
}
type Service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return Service{repository: repo}
}

func (s *Service) List() ([]repository.User, error) {
	return s.repository.List()
}
