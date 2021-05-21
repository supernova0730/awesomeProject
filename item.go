package awesomeProject

import "errors"

type Item struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type UpdateItemInput struct {
	Title *string `json:"title"`
	Done  *bool   `json:"done"`
}

func (i *UpdateItemInput) Validate() error {
	if i.Title == nil && i.Done == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
