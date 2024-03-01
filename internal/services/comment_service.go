package services

import "antibomberman/post/internal/repositories/postgres"

type CommentService interface {
}

type commentService struct {
	commentRepository postgres.CommentRepository
}

func NewCommentService(commentRepo *postgres.CommentRepository) CommentService {
	return commentService{
		commentRepository: commentRepo,
	}
}
