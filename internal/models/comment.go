package models

type Comment struct {
	ID        int
	Content   string
	UserID    int
	PostID    int
	CreatedAt string
}
type CommentCreate struct {
	Content string `json:"content" db:"content" `
	UserID  int    `json:"user_id" db:"user_id" `
	PostID  int    `json:"post_id" db:"post_id"`
}
