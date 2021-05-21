package service

import (
	awesomeProject "app"
	"app/pkg/repository"
)

type ItemService struct {
	itemRepo repository.TodoItem
}

func NewItemService(itemRepo repository.TodoItem) *ItemService {
	return &ItemService{itemRepo}
}

func (s *ItemService) Create(item awesomeProject.Item) (int, error) {
	return s.itemRepo.Create(item)
}

func (s *ItemService) GetAll() ([]awesomeProject.Item, error) {
	return s.itemRepo.GetAll()
}

func (s *ItemService) GetByID(id int) (awesomeProject.Item, error) {
	return s.itemRepo.GetByID(id)
}

func (s *ItemService) Delete(id int) error {
	return s.itemRepo.Delete(id)
}

func (s *ItemService) Update(id int, input awesomeProject.UpdateItemInput) error {
	return s.itemRepo.Update(id, input)
}

func (s *ItemService) UpdateTitle(id int, title string) error {
	input := awesomeProject.UpdateItemInput{
		Title: &title,
	}

	err := input.Validate()
	if err != nil {
		return err
	}

	return s.itemRepo.Update(id, input)
}

func (s *ItemService) Done(id int) error {
	done := true
	input := awesomeProject.UpdateItemInput{
		Done: &done,
	}

	err := input.Validate()
	if err != nil {
		return err
	}

	return s.itemRepo.Update(id, input)
}

func (s *ItemService) Undo(id int) error {
	done := false
	input := awesomeProject.UpdateItemInput{
		Done: &done,
	}

	err := input.Validate()
	if err != nil {
		return err
	}

	return s.itemRepo.Update(id, input)
}