package jwtService

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"sls/internal/service/tokenService"
	"time"
)

type JWTMaker struct {
	SecretKey string `json:"secret_key"`
}

func (J *JWTMaker) VerifyToken(token string) (*tokenService.TokenPayload, error) {
	Keyfunc := func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, tokenService.InvalidToken
		}
		return []byte(J.SecretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &tokenService.TokenPayload{}, Keyfunc)
	if err != nil {
		validationError, ok := err.(*jwt.ValidationError)
		if ok && validationError.Is(tokenService.ExpiredTokenErr) {
			return nil, tokenService.ExpiredTokenErr
		}
		return nil, tokenService.InvalidToken
	}

	payload, ok := jwtToken.Claims.(*tokenService.TokenPayload)
	if !ok {
		return nil, tokenService.InvalidToken
	}

	return payload, nil
}

func (J *JWTMaker) CreateToken(id uuid.UUID, duration time.Duration) (string, error) {
	payload, err := tokenService.NewTokenPayload(id, duration)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedString, err := token.SignedString([]byte(J.SecretKey))
	if err != nil {
		return "", err
	}
	return signedString, err
}

func NewJWTMaker(secretKey string) tokenService.TokenSrv {
	return &JWTMaker{SecretKey: secretKey}
}
