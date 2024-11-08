package models

import (
	"time"
)

type TokenMetadata struct {
    UserID    uint      `json:"user_id"`
    Email     string    `json:"email"`
    IssuedAt  time.Time `json:"issued_at"`
    ExpiresAt time.Time `json:"expires_at"`
}

type TokenValidationResult struct {
    Valid         		bool           		`json:"valid"`
    Message       		string         		`json:"message"`
    Metadata      		*TokenMetadata 		`json:"metadata,omitempty"`
    RemainingTime 		time.Duration      	`json:"remaining_time,omitempty"`
}

type TokenValidationResponse struct {
    Valid         		bool           		`json:"valid"`
    Message       		string         		`json:"message"`
    Metadata      		*TokenMetadata 		`json:"metadata,omitempty"`
	RemainingTime 		int64      	`json:"remaining_time,omitempty"`
}

type ValidateTokenRequest struct {
    Token string `json:"token" validate:"required"`
}
