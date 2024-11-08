package services

import (
	"github.com/tamabsndra/miniproject/miniproject-backend/models"
	"github.com/tamabsndra/miniproject/miniproject-backend/repository"
)

type PostService struct {
    postRepo *repository.PostRepository
}

func NewPostService(postRepo *repository.PostRepository) *PostService {
    return &PostService{
        postRepo: postRepo,
    }
}

func (s *PostService) Create(userID uint, req models.CreatePostRequest) (*models.Post, error) {
    post := &models.Post{
        UserID:  userID,
        Title:   req.Title,
        Content: req.Content,
    }

    return s.postRepo.Create(post)
}

func (s *PostService) GetAll() ([]models.Post, error) {
    return s.postRepo.GetAll()
}

func (s *PostService) GetByID(id uint) (*models.Post, error) {
    return s.postRepo.GetByID(id)
}

func (s *PostService) GetByUserID(userID uint) ([]models.Post, error) {
	return s.postRepo.GetByUserID(userID)
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
