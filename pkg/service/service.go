package service

import (
	awesomeProject "app"
	"app/pkg/repository"
)

type TodoItem interface {
	Create(item awesomeProject.Item) (int, error)
	GetAll() ([]awesomeProject.Item, error)
	GetByID(id int) (awesomeProject.Item, error)
	Delete(id int) error
	Update(id int, input awesomeProject.UpdateItemInput) error
	UpdateTitle(id int, title string) error
	Done(id int) error
	Undo(id int) error
}

type Service struct {
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{NewItemService(repo.TodoItem)}
}
