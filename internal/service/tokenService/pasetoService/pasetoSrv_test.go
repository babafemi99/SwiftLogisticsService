package pasetoService

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var (
	id = uuid.New()
)

func Test_pasetoMaker_CreateToken(t *testing.T) {
	maker := NewPasetoMaker([]byte("BABAGODANSWERPRAYERFEBRUARYTODEC"))
	require.NotNil(t, maker)
	token, _ := maker.CreateToken(id, time.Minute*10)
	fmt.Println("final token is: ", token)
}

func Test_pasetoMaker_VerifyToken(t *testing.T) {
	maker := NewPasetoMaker([]byte("BABAGODANSWERPRAYERFEBRUARYTODEC"))
	require.NotNil(t, maker)

	token, _ := maker.CreateToken(id, time.Minute*10)
	tokenPayload, err := maker.VerifyToken(token)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v, \n%v", tokenPayload.Id, id)
	assert.Equal(t, id, tokenPayload.Id)
	require.NoError(t, err)
}
