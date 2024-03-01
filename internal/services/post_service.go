package services

import (
	"antibomberman/post/internal/models"
	"antibomberman/post/internal/repositories/postgres"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

type PostService interface {
	Create(models.PostCreate) (models.Post, error)
	GetById(string) (models.Post, error)
	All() ([]models.Post, error)
	Update(string, models.PostUpdate) (models.Post, error)
	Delete(string) error
}

type postService struct {
	postRepository postgres.PostRepository
}

func NewPostService(postRepo *postgres.PostRepository) PostService {
	return &postService{
		postRepository: *postRepo,
	}
}
func (s *postService) Create(data models.PostCreate) (models.Post, error) {
	id, err := s.postRepository.Create(data)
	if err != nil {
		return models.Post{}, err
	}
	post, err := s.postRepository.GetById(strconv.Itoa(id))
	if err != nil {
		return models.Post{}, err
	}
	return post, nil
}
func (s *postService) GetById(id string) (models.Post, error) {
	post, err := s.postRepository.GetById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Post{}, fmt.Errorf("post with ID %s not found", id)
		}
		return models.Post{}, err
	}
	return post, nil

}
func (s *postService) All() ([]models.Post, error) {
	posts, err := s.postRepository.All()
	if err != nil {
		return nil, err
	}
	return posts, nil
}
func (s *postService) Update(id string, data models.PostUpdate) (models.Post, error) {
	err := s.postRepository.Update(id, data)
	if err != nil {
		return models.Post{}, err
	}
	post, err := s.postRepository.GetById(id)
	if err != nil {
		return models.Post{}, err
	}
	return post, nil
}
func (s *postService) Delete(id string) error {
	err := s.postRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
