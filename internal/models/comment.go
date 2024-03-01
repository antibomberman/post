package models

type Comment struct {
	ID        int
	Content   string
	UserID    int
	PostID    int
	CreatedAt string
}
