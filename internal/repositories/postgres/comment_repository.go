package postgres

import "antibomberman/post/internal/models"

type CommentRepository interface {
	Create(models.CommentCreate) (int, error)
	Delete(string) error
	GetByPostId(string) ([]models.Comment, error)
	GetById(string) (models.Comment, error)
}

type commentRepository struct {
	db *Postgres
}

func NewCommentRepository(db *Postgres) CommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (r *commentRepository) Create(data models.CommentCreate) (int, error) {
	var commentId int
	err := r.db.QueryRow(`
        INSERT INTO comments (content, user_id, post_id)
        VALUES ($1, $2, $3)
        RETURNING id
    `, data.Content, data.UserID, data.PostID).Scan(&commentId)

	if err != nil {
		return 0, err
	}

	return commentId, nil
}
func (r *commentRepository) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM comments WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil

}
func (r *commentRepository) GetByPostId(postId string) ([]models.Comment, error) {
	var comments []models.Comment
	err := r.db.Select(&comments, "SELECT * FROM comments WHERE post_id = $1 order by id desc", postId)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
func (r *commentRepository) GetById(id string) (models.Comment, error) {
	var comment models.Comment
	err := r.db.Get(&comment, "SELECT * FROM comments WHERE id = $1", id)
	if err != nil {
		return models.Comment{}, err
	}
	return comment, nil

}
