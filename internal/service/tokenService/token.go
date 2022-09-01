package tokenService

import (
	"github.com/google/uuid"
	"time"
)

type TokenSrv interface {
	CreateToken(id uuid.UUID, duration time.Duration) (string, error)
	VerifyToken(token string) (*TokenPayload, error)
}
