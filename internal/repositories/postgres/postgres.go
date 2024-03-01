package postgres

import (
	"antibomberman/post/internal/config"
	"github.com/jmoiron/sqlx"
	"log"
)

type Postgres struct {
	*sqlx.DB
}

func New(cfg *config.Config) *Postgres {
	db, err := sqlx.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	return &Postgres{db}
}
