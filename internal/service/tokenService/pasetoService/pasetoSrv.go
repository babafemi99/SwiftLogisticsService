package pasetoService

import (
	"aidanwoods.dev/go-paseto"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/chacha20"
	"log"
	"sls/internal/service/tokenService"
	"time"
)

type pasetoMaker struct {
	makerStr []byte
}

func (p *pasetoMaker) CreateToken(id uuid.UUID, duration time.Duration) (string, error) {
	token := paseto.NewToken()
	token.Set("id", id)
	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(duration))

	bytes, err := paseto.V4SymmetricKeyFromBytes(p.makerStr)
	if err != nil {
		return "", err
	}
	encrypt := token.V4Encrypt(bytes, nil)
	return encrypt, nil
}

func (p *pasetoMaker) VerifyToken(token string) (*tokenService.TokenPayload, error) {
	bytes, err := paseto.V4SymmetricKeyFromBytes(p.makerStr)
	if err != nil {
		return nil, fmt.Errorf("error getting symetics")
	}

	parser := paseto.NewParser()
	local, err := parser.ParseV4Local(bytes, token, nil)
	if err != nil {
		fmt.Println("inside here")
		return nil, fmt.Errorf("error parsing -: %v", err)
	}
	id, err := local.GetString("id")
	if err != nil {
		log.Fatalf("err : %v", err)
	}
	iat, err := local.GetTime("iat")
	if err != nil {
		return nil, err
	}
	exp, err := local.GetTime("exp")
	if err != nil {
		return nil, err
	}

	tknpayload := &tokenService.TokenPayload{
		Id:        uuid.MustParse(id),
		IssuedAt:  iat,
		ExpiredAt: exp,
	}
	return tknpayload, nil
}

func NewPasetoMaker(makeStr []byte) tokenService.TokenSrv {
	if len(makeStr) != chacha20.KeySize {
		return nil
	}
	return &pasetoMaker{
		makerStr: makeStr,
	}
}

type tk struct {
	Id uuid.UUID `json:"id"`
}
