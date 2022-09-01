package jwtService

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sls/internal/service/tokenService"
	"testing"
	"time"
)

var (
	id = uuid.New()
)

func TestJWTMaker_CreateToken(t *testing.T) {
	id := uuid.New()
	type fields struct {
		SecretKey string
	}
	type args struct {
		id       uuid.UUID
		duration time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Correct Credentials",
			fields: fields{
				SecretKey: "SWIFTSLOGISTICSSERVICE-ELECTRICTY!VIBESONAFREQUENCY",
			},
			args: args{
				id:       id,
				duration: time.Minute,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			J := &JWTMaker{
				SecretKey: tt.fields.SecretKey,
			}
			got, err := J.CreateToken(tt.args.id, tt.args.duration)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotNil(t, got)
			assert.Nil(t, err)

		})
	}
}

func TestJWTMaker_VerifyToken_InvalidToken(t *testing.T) {
	payload, err := tokenService.NewTokenPayload(id, time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	signedString, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker := NewJWTMaker("BABAYIGA")

	payload, err = maker.VerifyToken(signedString)
	require.Error(t, err)
	require.EqualError(t, err, tokenService.InvalidToken.Error())
	require.Nil(t, payload)

}

func TestJWTMaker_VerifyToken_ExpiredToken(t *testing.T) {
	maker := NewJWTMaker("SECRETKEYTHATNOONEISMEANTTOKNOWBUTNAEINBETHIS")
	token, err := maker.CreateToken(id, time.Minute*-1)
	require.NoError(t, err)
	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, tokenService.ExpiredTokenErr.Error())
	require.Nil(t, payload)
}

func TestJWTMaker_VerifyToken_ValidToken(t *testing.T) {
	maker := NewJWTMaker("SECRETKEYTHATNOONEISMEANTTOKNOWBUTNAEINBETHIS")
	token, err := maker.CreateToken(id, time.Minute*1)
	require.Nil(t, err)
	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotNil(t, payload)
}
