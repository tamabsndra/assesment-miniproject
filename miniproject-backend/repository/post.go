package repository

import (
	"database/sql"

	"github.com/tamabsndra/miniproject/miniproject-backend/models"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(post *models.Post) (*models.Post, error) {
	query := `
        INSERT INTO posts (user_id, title, content, is_published, created_at, updated_at)
        VALUES ($1, $2, $3, $4, NOW(), NOW())
        RETURNING id, created_at, updated_at
    `
	err := r.db.QueryRow(
		query,
		post.UserID,
		post.Title,
		post.Content,
		post.IsPublished,
	).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return post, nil
}

func (r *PostRepository) GetAll() ([]models.Post, error) {
	query := `
        SELECT id, user_id, title, content, is_published, created_at, updated_at
        FROM posts
        ORDER BY created_at DESC
    `
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.ID,
			&post.UserID,
			&post.Title,
			&post.Content,
			&post.IsPublished,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (r *PostRepository) GetByID(id uint) (*models.Post, error) {
	post := &models.Post{}
	query := `
		SELECT id, user_id, title, content, is_published, created_at, updated_at
		FROM posts WHERE id = $1
    `
	err := r.db.QueryRow(query, id).Scan(
		&post.ID,
		&post.UserID,
		&post.Title,
		&post.Content,
		&post.IsPublished,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (r *PostRepository) GetByUserID(userID uint) ([]models.Post, error) {
	query := `
		SELECT id, user_id, title, content, is_published, created_at, updated_at
		FROM posts
		WHERE user_id = $1
		ORDER BY created_at DESC
	`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.ID,
			&post.UserID,
			&post.Title,
			&post.Content,
			&post.IsPublished,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (r *PostRepository) Update(id uint, req models.UpdatePostRequest) (*models.Post, error) {
	post := &models.Post{}
	query := `
		UPDATE posts
		SET title = $1, content = $2, is_published = $3, updated_at = NOW()
		WHERE id = $4
		RETURNING id, user_id, title, content, is_published, created_at, updated_at
	`
	err := r.db.QueryRow(
		query,
		req.Title,
		req.Content,
		req.IsPublished,
		id,
	).Scan(
		&post.ID,
		&post.UserID,
		&post.Title,
		&post.Content,
		&post.IsPublished,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (r *PostRepository) Delete(id uint) error {
	_, err := r.db.Exec("DELETE FROM posts WHERE id = $1", id)
	return err
}

func (r *PostRepository) GetPostDetail() ([]models.PostWithUser, error) {
	query := `
		SELECT p.id, p.user_id, p.title, p.content, p.is_published, p.created_at, p.updated_at, u.id, u.name, u.email
		FROM posts p
		JOIN users u ON p.user_id = u.id
		ORDER BY p.created_at DESC
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.PostWithUser
	for rows.Next() {
		var post models.PostWithUser
		err := rows.Scan(
			&post.ID,
			&post.UserID,
			&post.Title,
			&post.Content,
			&post.IsPublished,
			&post.CreatedAt,
			&post.UpdatedAt,
			&post.User.ID,
			&post.User.Name,
			&post.User.Email,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
