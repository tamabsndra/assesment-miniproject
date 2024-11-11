package services

import (
	"github.com/tamabsndra/miniproject/miniproject-backend/models"
	"github.com/tamabsndra/miniproject/miniproject-backend/repository"
	"github.com/tamabsndra/miniproject/miniproject-backend/utils"
)

type PostService struct {
	postRepo  *repository.PostRepository
	jwtSecret string
}

func NewPostService(postRepo *repository.PostRepository, jwtSecret string) *PostService {
	return &PostService{
		postRepo: postRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *PostService) Create(userID uint, req models.CreatePostRequest) (*models.Post, error) {
	post := &models.Post{
		UserID:  userID,
		Title:   req.Title,
		Content: req.Content,
		IsPublished: req.IsPublished,
	}

	return s.postRepo.Create(post)
}

func (s *PostService) GetAll() ([]models.Post, error) {
	return s.postRepo.GetAll()
}

func (s *PostService) GetByID(id uint) (*models.Post, error) {
	return s.postRepo.GetByID(id)
}

func (s *PostService) GetByUserID(token string) ([]models.Post, error) {
	id, err := utils.ValidateToken(token, s.jwtSecret)
	if err != nil {
		return nil, err
	}
	return s.postRepo.GetByUserID(id.UserID)
}

func (s *PostService) Update(id uint, req models.UpdatePostRequest) (*models.Post, error) {
	return s.postRepo.Update(id, req)
}

func (s *PostService) Delete(id uint) error {
	return s.postRepo.Delete(id)
}

func (s *PostService) GetPostDetail() ([]models.PostWithUser, error) {
	return s.postRepo.GetPostDetail()
}
