package repository

import (
	awesomeProject "app"
	"database/sql"
)

type TodoItem interface {
	Create(item awesomeProject.Item) (int, error)
	GetAll() ([]awesomeProject.Item, error)
	GetByID(id int) (awesomeProject.Item, error)
	Delete(id int) error
	Update(id int, input awesomeProject.UpdateItemInput) error
}

type Repository struct {
	TodoItem
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		TodoItem: NewItemPostgres(db),
	}
}
