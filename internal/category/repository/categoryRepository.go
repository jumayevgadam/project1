package repository

import (
	"Project1/internal/category/model"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type CategoryRepository struct {
	DB *sqlx.DB
}

const Categories = "Categories"

func NewCategoryRepository(DB *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{DB: DB}
}

func (r *CategoryRepository) GetAll() ([]*model.Category, error) {
	query := fmt.Sprintf("SELECT id, name, created_at, updated_at FROM %s", Categories)
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*model.Category
	for rows.Next() {
		var category model.Category
		err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepository) GetOne(id int) (*model.Category, error) {
	query := fmt.Sprintf("SELECT id, name, created_at, updated_at FROM %s WHERE id = $1", Categories)
	row := r.DB.QueryRow(query, id)
	var category model.Category
	err := row.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepository) Create(body *model.Category) (*model.Category, error) {
	if body.Name == "" {
		return nil, errors.New("name is required")
	}

	category := &model.Category{
		Name:      body.Name,
		CreatedAt: body.CreatedAt,
		UpdatedAt: body.UpdatedAt,
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (name) VALUES ($1) RETURNING id", Categories)

	row := r.DB.QueryRow(query, category.Name)
	var insertedID int
	if err := row.Scan(&insertedID); err != nil {
		return nil, err
	}

	return &model.Category{
		Id:   insertedID,
		Name: body.Name,
		//CreatedAt: returnedCreatedAt,
		//UpdatedAt: returnedUpdatedAt,
	}, nil
}

func (r *CategoryRepository) Update(id int, body *model.Category) error {
	if body.Name == "" {
		return errors.New("name is required")
	}

	// Use database's built-in timestamp function to get the current timestamp
	query := fmt.Sprintf(
		`UPDATE %s SET name = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2`, Categories)

	_, err := r.DB.Exec(query, body.Name, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *CategoryRepository) Delete(id int) error {
	query := fmt.Sprintf(
		`DELETE FROM %s WHERE id = $1`, Categories)

	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
