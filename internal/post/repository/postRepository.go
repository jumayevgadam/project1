package repository

import (
	"Project1/internal/post/model"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PostRepository struct {
	DB *sqlx.DB
}

const posts = "posts"

func NewPostRepository(DB *sqlx.DB) *PostRepository {
	return &PostRepository{DB: DB}
}

func (r *PostRepository) GetAll() ([]*model.Post, error) {
	query := `SELECT p.id, p.title, p.description,  p.category_id CategoryId, p.user_id UserId, p.image_path ImagePath, c.name CategoryName, u.name AS UserName
				FROM posts p
				INNER JOIN Categories c ON p.category_id = c.id
				INNER JOIN users u ON p.user_id = u.id
				ORDER BY p.id DESC LIMIT 20`
	var posts []*model.Post
	err := r.DB.Select(&posts, query)
	if err != nil {
		return nil, err
	}
	return posts, nil

}

func (r *PostRepository) GetOne(id int) (*model.Post, error) {
	query := `
        SELECT p.id, p.title, p.description, p.category_id CategoryId, p.user_id UserId
        FROM posts p WHERE p.id = $1`

	var post model.Post
	err := r.DB.Get(&post, query, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) Create(body *model.Post) (*model.Post, error) {
	if body.Title == "" {
		return nil, errors.New("Title is required")
	}

	query := `INSERT INTO posts (title, description, category_id, user_id, image_path)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	row := r.DB.QueryRow(query, body.Title, body.Description, body.CategoryId, body.UserId, body.ImagePath)
	var InsertedId int
	if err := row.Scan(&InsertedId); err != nil {
		return nil, err
	}
	body.Id = InsertedId
	return body, nil
}

func (r *PostRepository) Update(id int, body *model.Post) error {
	query := fmt.Sprintf(
		`UPDATE %s SET title = $1, description = $2, category_id = $3, user_id = $4 WHERE id = $5`, posts)

	_, err := r.DB.Exec(query, body.Title, body.Description, body.CategoryId, body.UserId, id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (r *PostRepository) Delete(id int) error {
	query := fmt.Sprintf(
		`DELETE FROM %s WHERE id = $1`, posts)

	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
