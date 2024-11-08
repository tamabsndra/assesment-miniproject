package services

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/tamabsndra/miniproject/miniproject-backend/models"
	"github.com/tamabsndra/miniproject/miniproject-backend/utils"
)

type TokenService struct {
	redis       *redis.Client
	tokenExpiry time.Duration
	jwtSecret   string
}

func NewTokenService(redis *redis.Client, tokenExpiry time.Duration, jwtSecret string) *TokenService {
	return &TokenService{
		redis:       redis,
		tokenExpiry: tokenExpiry,
		jwtSecret:   jwtSecret,
	}
}

func (s *TokenService) ValidateToken(token string) (*models.TokenValidationResult, error) {
	if s.IsTokenBlacklisted(token) {
		return &models.TokenValidationResult{
			Valid:   false,
			Message: "Token has been revoked",
		}, nil
	}

	claims, err := utils.ValidateToken(token, s.jwtSecret)
	if err != nil {
		var message string
		switch err {
		case utils.ErrTokenExpired:
			message = "Token has expired"
		case utils.ErrTokenNotValidYet:
			message = "Token is not active yet"
		case utils.ErrTokenMalformed:
			message = "Token is malformed"
		case utils.ErrTokenTypeInvalid:
			message = "Invalid token type"
		default:
			message = "Invalid token"
		}
		return &models.TokenValidationResult{
			Valid:   false,
			Message: message,
		}, nil
	}

	metadata, err := utils.ExtractTokenMetadata(claims)
	if err != nil {
		return &models.TokenValidationResult{
			Valid:   false,
			Message: "Failed to extract token metadata",
		}, err
	}

	remainingTime := utils.GetTokenRemainingTime(claims)

	return &models.TokenValidationResult{
		Valid:         true,
		Message:       "Token is valid",
		Metadata:      metadata,
		RemainingTime: remainingTime,
	}, nil
}

func (s *TokenService) BlacklistToken(token string) error {
	ctx := context.Background()
	key := fmt.Sprintf("blacklist:%s", token)
	return s.redis.Set(ctx, key, "true", s.tokenExpiry).Err()
}

func (s *TokenService) IsTokenBlacklisted(token string) bool {
	ctx := context.Background()
	key := fmt.Sprintf("blacklist:%s", token)
	exists, err := s.redis.Exists(ctx, key).Result()
	return err == nil && exists > 0
}
