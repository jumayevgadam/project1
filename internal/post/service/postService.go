package service

import (
	"Project1/internal/post/model"
	"Project1/internal/post/repository"
)

type PostService struct {
	repo *repository.PostRepository
}

func NewPostService(repo *repository.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(body *model.Post) (*model.Post, error) {
	return s.repo.Create(body)
}

func (s *PostService) GetPostByID(id int) (*model.Post, error) {
	return s.repo.GetOne(id)
}

func (s *PostService) GetAllPosts() ([]*model.Post, error) {
	return s.repo.GetAll()
}

func (s *PostService) DeletePostById(id int) error {
	return s.repo.Delete(id)
}

func (s *PostService) UpdatePost(id int, body *model.Post) error {
	err := s.repo.Update(id, body)
	if err != nil {
		return err
	}

	return err
}
