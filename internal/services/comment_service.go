package services

import (
	"antibomberman/post/internal/models"
	"antibomberman/post/internal/repositories/postgres"
	"strconv"
)

type CommentService interface {
	Create(models.CommentCreate) (models.Comment, error)
	Delete(string) error
	GetByPostId(string) ([]models.Comment, error)
}

type commentService struct {
	commentRepository postgres.CommentRepository
}

func NewCommentService(commentRepo *postgres.CommentRepository) CommentService {
	return &commentService{
		commentRepository: *commentRepo,
	}
}

func (s *commentService) Create(data models.CommentCreate) (models.Comment, error) {
	id, err := s.commentRepository.Create(data)
	if err != nil {
		return models.Comment{}, err
	}
	comment, err := s.commentRepository.GetById(strconv.Itoa(id))
	if err != nil {
		return models.Comment{}, err
	}
	return comment, nil
}
func (s *commentService) Delete(id string) error {
	err := s.commentRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
func (s *commentService) GetByPostId(postId string) ([]models.Comment, error) {
	comments, err := s.commentRepository.GetByPostId(postId)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
