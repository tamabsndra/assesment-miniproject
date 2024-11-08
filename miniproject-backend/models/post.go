package models

import "time"

type Post struct {
    ID          uint      `json:"id"`
    UserID      uint      `json:"user_id"`
    Title       string    `json:"title" validate:"required,min=3,max=100"`
    Content     string    `json:"content" validate:"required,min=10"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type PostWithUser struct {
	Post
	User UserInPost `json:"user"`
}

type CreatePostRequest struct {
    Title   string `json:"title" validate:"required,min=3,max=100"`
    Content string `json:"content" validate:"required,min=10"`
}

type UpdatePostRequest struct {
	Title   string `json:"title" validate:"required,min=3,max=100"`
	Content string `json:"content" validate:"required,min=10"`
}
