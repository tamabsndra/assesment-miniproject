package services

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/tamabsndra/miniproject/miniproject-backend/models"
	"github.com/tamabsndra/miniproject/miniproject-backend/repository"
	"github.com/tamabsndra/miniproject/miniproject-backend/utils"
)

type AuthService struct {
	userRepo  *repository.UserRepository
	jwtSecret string
}

func NewAuthService(userRepo *repository.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *AuthService) Login(req models.LoginRequest) (*models.LoginResponse, error) {
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(*user, s.jwtSecret, 24*time.Hour)
	if err != nil {
		return nil, err
	}

	user.Password = ""
	return &models.LoginResponse{
		Token: token,
		User:  *user,
	}, nil
}

func (s *AuthService) Register(req models.User) error {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	req.Password = hashedPassword
	req.CreatedAt = utils.GetCurrentTime()
	req.UpdatedAt = utils.GetCurrentTime()

	return s.userRepo.Create(&req)
}

func (s *AuthService) GetMe(token string) (*models.User, error) {
	id, err := utils.ValidateToken(token, s.jwtSecret)
	if err != nil {
		return nil, err
	}
	return s.userRepo.GetByID(id.UserID)
}
