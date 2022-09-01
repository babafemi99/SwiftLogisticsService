package userService

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
	"net/http"
	"reflect"
	"sls/internal/Repository/userRepository"
	"sls/internal/entity/errorEntity"
	"sls/internal/entity/userEntity"
	"sls/internal/service/cryptoService"
	"sls/internal/service/timeService"
	"sls/internal/service/validationService"
	"testing"
	"time"
)

var (
	log    = logrus.New()
	tstErr = fmt.Errorf("validation Error")
)

// mock Time Service
type mockRepo struct {
	mock.Mock
}

func (m mockRepo) GetById(id uuid.UUID) (*userEntity.UserAccess, error) {
	args := m.Called(id)
	result := args.Get(0)
	return result.(*userEntity.UserAccess), args.Error(1)
}

func (m mockRepo) ChangePassword(id uuid.UUID, password string) error {
	args := m.Called(password, id)
	result := args.Error(0)
	return result
}

func (m mockRepo) UpdateProfile(req *userEntity.UpdateUserReq) (*userEntity.UpdateUserReq, error) {
	args := m.Called(req)
	result := args.Get(0)
	return result.(*userEntity.UpdateUserReq), args.Error(1)
}

func (m mockRepo) DeactivateAccount(id uuid.UUID) error {
	args := m.Called(id)
	result := args.Error(0)
	return result
}

func (m mockRepo) Persist(user *userEntity.CreateUser) (*userEntity.CreateUser, error) {
	args := m.Called(user)
	result := args.Get(0)
	return result.(*userEntity.CreateUser), args.Error(1)
}

func (m mockRepo) GetByEmail(email string) (*userEntity.UserAccess, error) {
	args := m.Called(email)
	result := args.Get(0)
	return result.(*userEntity.UserAccess), args.Error(1)
}

func (m mockRepo) GetByPhone(phone string) (*userEntity.UserAccess, error) {
	args := m.Called(phone)
	result := args.Get(0)
	return result.(*userEntity.UserAccess), args.Error(1)
}

//  Time Service
type mockTime struct {
	mock.Mock
}

func (m mockTime) Now() time.Time {
	args := m.Called()
	return args.Get(0).(time.Time)
}

// mock Crypto Service
type mockCrypto struct {
	mock.Mock
}

func (m mockCrypto) HashPassword(password string) (string, error) {
	args := m.Called(password)
	result := args.Get(0)
	return result.(string), args.Error(1)
}

func (m mockCrypto) ComparePassword(hashed, plain string) error {
	args := m.Called(hashed, plain)
	result := args.Error(0)
	return result
}

// mock Validation Service
type mockValidation struct {
	mock.Mock
}

func (m mockValidation) Validate(data interface{}) error {
	args := m.Called(data)
	return args.Error(0)
}

func Test_userSrv_DeleteAccount(t *testing.T) {
	id := uuid.New()
	wrongId := uuid.New()
	InvalidId := uuid.New()
	req := &userEntity.DeleteAccountReq{UserId: id}
	req2 := &userEntity.DeleteAccountReq{UserId: wrongId}
	invalidReq := &userEntity.DeleteAccountReq{UserId: InvalidId}

	//mock services
	crypto := new(mockCrypto)
	timeSrv := new(mockTime)

	vldSrv := new(mockValidation)
	vldSrv.On("Validate", req).Return(nil)
	vldSrv.On("Validate", req2).Return(nil)
	vldSrv.On("Validate", invalidReq).Return(tstErr)

	repoSrv := new(mockRepo)
	repoSrv.On("DeactivateAccount", req.UserId).Return(nil)
	repoSrv.On("DeactivateAccount", req2.UserId).Return(tstErr)

	type fields struct {
		log       *logrus.Logger
		cryptoSrv cryptoService.CryptoSrv
		timeSrv   timeService.TimeSrv
		vldSrv    validationService.ValidationService
		repoSrv   userRepository.UserRepo
	}
	type args struct {
		req *userEntity.DeleteAccountReq
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *errorEntity.ErrorRes
	}{
		{
			name: "Correct Credentials",
			fields: fields{
				log:       log,
				cryptoSrv: crypto,
				timeSrv:   timeSrv,
				vldSrv:    vldSrv,
				repoSrv:   repoSrv,
			},
			args: args{
				req: req,
			},
			want: nil,
		},
		{
			name: "Email not Found",
			fields: fields{
				log:       log,
				cryptoSrv: crypto,
				timeSrv:   timeSrv,
				vldSrv:    vldSrv,
				repoSrv:   repoSrv,
			},
			args: args{
				req: req2,
			},
			want: errorEntity.NewErrorRes(http.StatusInternalServerError, "Error completing deleting operation"),
		},
		{
			name: "Invalid Email",
			fields: fields{
				log:       log,
				cryptoSrv: crypto,
				timeSrv:   timeSrv,
				vldSrv:    vldSrv,
				repoSrv:   repoSrv,
			},
			args: args{
				req: invalidReq,
			},
			want: errorEntity.NewErrorRes(http.StatusBadRequest, "Bad Request !"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userSrv{
				log:       tt.fields.log,
				cryptoSrv: tt.fields.cryptoSrv,
				timeSrv:   tt.fields.timeSrv,
				vldSrv:    tt.fields.vldSrv,
				repoSrv:   tt.fields.repoSrv,
			}
			if got := u.DeleteAccount(tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
