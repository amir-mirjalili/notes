package service

import (
	"github.com/amir-mirjalili/notes.git/notes/params"
	"github.com/amir-mirjalili/notes.git/notes/repository"
)

type Repository interface {
	Create(note params.AddNotRequest) error
	Update(id int, note repository.Note) (repository.Note, error)
	GetByID(id int) (repository.Note, error)
	GetList() ([]repository.Note, error)
	Delete(ID int) error
}

type Service struct {
	repository Repository
}

func NewService(repo Repository) *Service {

	return &Service{repository: repo}
}

func (s *Service) Create(note params.AddNotRequest) error {
	err := s.repository.Create(note)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(ID int, note repository.Note) (repository.Note, error) {
	response, err := s.repository.Update(ID, note)
	return response, err
}

func (s *Service) GetById(id int) (repository.Note, error) {
	note, err := s.repository.GetByID(id)
	return note, err
}

func (s *Service) GetList() ([]repository.Note, error) {
	notes, err := s.repository.GetList()
	return notes, err
}
func (s *Service) Delete(ID int) error {
	return s.repository.Delete(ID)
}
