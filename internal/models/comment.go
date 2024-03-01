package models

type Comment struct {
	ID        int    `db:"id" json:"id"`
	Content   string `db:"content" json:"content"`
	UserID    int    `db:"user_id" json:"userID"`
	PostID    int    `db:"post_id" json:"postID"`
	CreatedAt string `db:"created_at" json:"createdAt"`
}
type CommentCreate struct {
	Content string `json:"content" db:"content" `
	UserID  int    `json:"user_id" db:"user_id" `
	PostID  int    `json:"post_id" db:"post_id"`
}
