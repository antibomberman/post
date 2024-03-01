package postgres

type CommentRepository interface {
}

type commentRepository struct {
	db *Postgres
}

func NewCommentRepository(db *Postgres) CommentRepository {
	return &commentRepository{
		db: db,
	}
}
