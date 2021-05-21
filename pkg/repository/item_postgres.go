package repository

import (
	awesomeProject "app"
	"database/sql"
	"fmt"
	"strings"
)

type ItemPostgres struct {
	db *sql.DB
}

func NewItemPostgres(db *sql.DB) *ItemPostgres {
	return &ItemPostgres{db}
}

func (r *ItemPostgres) Create(item awesomeProject.Item) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (title, done) VALUES($1, $2) RETURNING id", itemsTable)

	row := r.db.QueryRow(query, item.Title, item.Done)

	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ItemPostgres) GetAll() ([]awesomeProject.Item, error) {
	var items []awesomeProject.Item

	query := fmt.Sprintf("SELECT id, title, done FROM %s ORDER BY id", itemsTable)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item awesomeProject.Item

		err = rows.Scan(&item.ID, &item.Title, &item.Done)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		items = append(items, item)
	}

	return items, nil
}

func (r *ItemPostgres) GetByID(id int) (awesomeProject.Item, error) {
	var item awesomeProject.Item

	query := fmt.Sprintf("SELECT id, title, done FROM %s WHERE id = $1", itemsTable)

	row := r.db.QueryRow(query, id)

	err := row.Scan(&item.ID, &item.Title, &item.Done)
	if err != nil {
		return awesomeProject.Item{}, err
	}

	return item, nil
}

func (r *ItemPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", itemsTable)

	_, err := r.db.Exec(query, id)
	return err
}

func (r *ItemPostgres) Update(id int, input awesomeProject.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", itemsTable, setQuery, argId)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}
