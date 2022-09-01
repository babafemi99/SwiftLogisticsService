package tokenEntity

import (
	"github.com/google/uuid"
	"time"
)

type tokenPayload struct {
	Id        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Admin     bool      `json:"admin"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}
