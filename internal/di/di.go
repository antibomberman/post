package di

import (
	"antibomberman/post/internal/config"
	"antibomberman/post/internal/repositories/postgres"
	"antibomberman/post/internal/services"
)

type DI struct {
	Cfg            *config.Config
	UserService    services.UserService
	PostService    services.PostService
	CommentService services.CommentService
}

func NewDI(db *postgres.Postgres, cfg *config.Config) *DI {
	userRepo := postgres.NewUserRepository(db)
	postRepo := postgres.NewPostRepository(db)
	commentRepo := postgres.NewCommentRepository(db)

	return &DI{
		Cfg:            cfg,
		UserService:    services.NewUserService(&userRepo, cfg),
		PostService:    services.NewPostService(&postRepo),
		CommentService: services.NewCommentService(&commentRepo),
	}
}
