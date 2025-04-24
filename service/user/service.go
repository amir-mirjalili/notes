package user

import repository "github.com/amir-mirjalili/notes.git/repository/user"

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
