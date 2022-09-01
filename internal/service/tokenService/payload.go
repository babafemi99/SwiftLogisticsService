package tokenService

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

var (
	ExpiredTokenErr = errors.New("token Has Expired")
	InvalidToken    = fmt.Errorf("unexpected signing method; invalid token")
)

type TokenPayload struct {
	Id        uuid.UUID `json:"id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (t *TokenPayload) Valid() error {
	if time.Now().After(t.ExpiredAt) {
		return ExpiredTokenErr
	}
	return nil
}

func NewTokenPayload(id uuid.UUID, duration time.Duration) (*TokenPayload, error) {
	return &TokenPayload{
		Id:        id,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}, nil
}
